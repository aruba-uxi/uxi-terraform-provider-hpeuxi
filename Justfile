CLIENT_DIR := "pkg/config-api-client"
TOOLS_PROVIDER_DIR := "tools"
OPENAPI_SPEC := "pkg/config-api-client/api"
SOURCE_OPEN_API_SPEC_FILE := ".openapi.source.yaml"

# Show this message and exit.
help:
	@just --list

_remove-client-files:
  cd {{ CLIENT_DIR }} && cat .openapi-generator/FILES | xargs -n 1 rm -f

generate-config-api-client: _remove-client-files
  docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli generate \
  --input-spec /local/{{ OPENAPI_SPEC }}/{{ SOURCE_OPEN_API_SPEC_FILE }} \
  --generator-name go \
  --output /local/{{ CLIENT_DIR }} \
  --git-user-id aruba-uxi \
  --git-repo-id terraform-provider-hpeuxi/{{ CLIENT_DIR }} \
  --openapi-normalizer SET_TAGS_FOR_ALL_OPERATIONS=configuration \
  --additional-properties packageName=config_api_client,enumClassPrefix=true
  rm ./pkg/config-api-client/go.mod
  rm ./pkg/config-api-client/go.sum
  just tidy-client
  just fmt-client

setup-run:
  grep -q "registry.terraform.io/aruba-uxi/hpeuxi" ~/.terraformrc && echo "Dev override found - installing provider locally" || { echo "Dev override not found - please follow README setup guide"; exit 1; }
  go install .

remove-setup-run-dev-override:
  sed -i '' '/registry\.terraform\.io\/aruba-uxi\/hpeuxi/d' ~/.terraformrc

build-local:
  go run github.com/goreleaser/goreleaser/v2@latest release --clean --skip=publish,validate

sign:
  mkdir -p logs
  signhpe --logdir ./logs --in dist/$(ls dist | grep SHA256SUMS) --env --project "HPE Aruba Networking UXI Terraform Provider" --out ./dist

test-client +ARGS='':
  cd {{ CLIENT_DIR }} && go test -v ./... -race -covermode=atomic -coverprofile=.coverage {{ ARGS }}

coverage-client:
  cd {{ CLIENT_DIR }} && go tool cover -html=.coverage -o=.coverage.html

fmt-client:
  python -m tools.lint-attribution format
  go run github.com/segmentio/golines@v0.12.2 -w {{ CLIENT_DIR }}
  go run golang.org/x/tools/cmd/goimports@latest -local github.com/aruba-uxi -w {{ CLIENT_DIR }}
  go run mvdan.cc/gofumpt@latest -w {{ CLIENT_DIR }}

tidy-client:
  cd {{ CLIENT_DIR }} && go mod tidy

lint:
  #!/usr/bin/env bash

  output=$(gofmt -d .)
  if [ -n "$output" ]; then
    echo "$output"
    echo "Error: (gofmt) formatting required" >&2
    exit 1
  fi

  output=$(go run github.com/segmentio/golines@v0.12.2 . --dry-run)
  if [ -n "$output" ]; then
    echo "$output"
    echo "Error: (golines) formatting required" >&2
    exit 1
  fi

  go run github.com/golangci/golangci-lint/cmd/golangci-lint@v1.60.1 run

  python -m tools.lint-attribution lint

  terraform fmt -recursive -check

fmt:
  python -m tools.lint-attribution format
  go run github.com/segmentio/golines@v0.12.2 -w .
  go run golang.org/x/tools/cmd/goimports@latest -local github.com/aruba-uxi -w .
  go run mvdan.cc/gofumpt@latest -w .
  terraform fmt --recursive .

tidy-provider:
  go mod tidy

test-provider +ARGS='':
  TF_ACC=1 go test -v ./test/mocked/... -race -covermode=atomic -coverprofile=.coverage {{ ARGS }}

generate-provider-docs:
  cd {{ TOOLS_PROVIDER_DIR }} && go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs generate --provider-dir ../. --provider-name hpeuxi
  sed -i.backup '/subcategory: ""/d' docs/index.md && rm docs/index.md.backup

validate-provider-docs:
  cd {{ TOOLS_PROVIDER_DIR }} && go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs validate --provider-dir ../. --provider-name hpeuxi

coverage-provider:
  go tool cover -html=.coverage -o=.coverage.html

tidy-tools:
  cd {{ TOOLS_PROVIDER_DIR }} && go mod tidy

acceptance-tests ENV_PATH='' +ARGS='':
  #!/usr/bin/env bash

  if [ -z "{{ ENV_PATH }}" ]; then
    echo "Error: ENV_PATH argument is required" >&2
    exit 1
  fi

  if [ ! -f "{{ ENV_PATH }}" ]; then
    echo "Error: File specified ("{{ ENV_PATH }}") does not exist" >&2
    exit 1
  fi

  # we run these separately so that they do not interfere with each other since GoLang executes
  # tests in different directories at the same time
  for dir in "datasources" "resources"
  do
      echo "Running tests in $dir..."
      env $(cat {{ ENV_PATH }} | xargs) TF_ACC=1 go test -v ./test/live/$dir/... -race {{ ARGS }}
  done

test +ARGS='':
  just test-client {{ ARGS }}
  just test-provider {{ ARGS }}

coverage:
  just coverage-client
  just coverage-provider

tidy:
  just tidy-client
  just tidy-provider
  just tidy-tools

clean:
  find . -name ".coverage*" -type f -delete

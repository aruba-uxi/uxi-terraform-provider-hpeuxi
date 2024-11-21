TMP_DIR := "tmp"
CONFIG_API_CLIENT_DIR := "pkg/config-api-client"
CONFIG_API_PROVIDER_DIR := "."
TOOLS_PROVIDER_DIR := "tools"
OPENAPI_SPEC := "pkg/config-api-client/api"
SOURCE_OPEN_API_SPEC_FILE := ".openapi.source.yaml"

retrieve-config-api-openapi-spec:
  rm -rf {{ TMP_DIR }}
  git clone git@github.com:aruba-uxi/configuration-api.git --depth=1 {{ TMP_DIR }}
  mkdir -p {{ OPENAPI_SPEC }}
  cp {{ TMP_DIR }}/oas/openapi.yaml {{ OPENAPI_SPEC }}/{{ SOURCE_OPEN_API_SPEC_FILE }}
  rm -rf {{ TMP_DIR }}

cleanup-old-client-files:
  cd {{ CONFIG_API_CLIENT_DIR }} && cat .openapi-generator/FILES | xargs -n 1 rm -f

generate-config-api-client: retrieve-config-api-openapi-spec
  just cleanup-old-client-files
  docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli generate \
  --input-spec /local/{{ OPENAPI_SPEC }}/{{ SOURCE_OPEN_API_SPEC_FILE }} \
  --generator-name go \
  --output /local/{{ CONFIG_API_CLIENT_DIR }} \
  --package-name config_api_client \
  --git-user-id aruba-uxi \
  --git-repo-id terraform-provider-hpeuxi/{{ CONFIG_API_CLIENT_DIR }} \
  --openapi-normalizer SET_TAGS_FOR_ALL_OPERATIONS=configuration
  rm ./pkg/config-api-client/go.mod
  rm ./pkg/config-api-client/go.sum
  just tidy-client
  just fmt-client

#nothing to see here
setup-dev:

build-local:
  go run github.com/goreleaser/goreleaser/v2@latest release --clean --skip=publish,validate

sign:
  mkdir -p logs
  signhpe --logdir ./logs --in dist/$(ls dist | grep SHA256SUMS) --env --project "HPE Aruba Networking UXI Terraform Provider" --out ./dist

test-client +ARGS='':
  cd {{ CONFIG_API_CLIENT_DIR }} && go test -v ./... -race -covermode=atomic -coverprofile=.coverage {{ ARGS }}

coverage-client:
  cd {{ CONFIG_API_CLIENT_DIR }} && go tool cover -html=.coverage -o=.coverage.html

fmt-client:
  python -m tools.lint-attribution format
  gofmt -w {{ CONFIG_API_CLIENT_DIR }}
  go run github.com/segmentio/golines@v0.12.2 -w {{ CONFIG_API_CLIENT_DIR }}

tidy-client:
  cd {{ CONFIG_API_CLIENT_DIR }} && go mod tidy

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

fmt:
  python -m tools.lint-attribution format
  gofmt -w .
  go run github.com/segmentio/golines@v0.12.2 -w .

tidy-provider:
  cd {{ CONFIG_API_PROVIDER_DIR }} go mod tidy

test-provider +ARGS='':
  cd {{ CONFIG_API_PROVIDER_DIR }} && TF_ACC=1 go test -v ./test/mocked/... -race -covermode=atomic -coverprofile=.coverage {{ ARGS }}

generate-provider-docs:
  cd {{ TOOLS_PROVIDER_DIR }} && go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs generate --provider-dir ../{{ CONFIG_API_PROVIDER_DIR }} --provider-name uxi
  sed -i.backup '/subcategory: ""/d' ./{{ CONFIG_API_PROVIDER_DIR }}/docs/index.md && rm ./{{ CONFIG_API_PROVIDER_DIR }}/docs/index.md.backup

validate-provider-docs:
  cd {{ TOOLS_PROVIDER_DIR }} && go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs validate --provider-dir ../{{ CONFIG_API_PROVIDER_DIR }} --provider-name uxi

coverage-provider:
  cd {{ CONFIG_API_PROVIDER_DIR }} && go tool cover -html=.coverage -o=.coverage.html

tidy-tools:
  cd {{ TOOLS_PROVIDER_DIR }} && go mod tidy

acceptance-tests +ARGS='':
  #!/usr/bin/env bash

  read -p "This is going to run requests against UXI backend. Continue (y/Y)? " -n 1 -r
  echo
  if [[ $REPLY =~ ^[Yy]$ ]]
  then
    # we run these seperately so that they do not interfere with each other since GoLang executes
    # tests in different directories at the same time
    for dir in "datasources" "resources"
    do
        echo "Running tests in $dir..."
        TF_ACC=1 go test -v ./test/live/$dir/... -race {{ ARGS }}
    done
  fi


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

DEFAULT_EXAMPLE := "full-demo"

plan example=DEFAULT_EXAMPLE +ARGS='':
  cd {{ CONFIG_API_PROVIDER_DIR }} && go install .
  cd {{ CONFIG_API_PROVIDER_DIR }}/examples/{{example}} && terraform plan {{ ARGS }}

apply example=DEFAULT_EXAMPLE +ARGS='':
  cd {{ CONFIG_API_PROVIDER_DIR }} && go install .
  cd {{ CONFIG_API_PROVIDER_DIR }}/examples/{{example}} && terraform apply {{ ARGS }}

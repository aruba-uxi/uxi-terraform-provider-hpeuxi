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
  --git-repo-id terraform-provider-configuration-api/{{ CONFIG_API_CLIENT_DIR }} \
  --openapi-normalizer SET_TAGS_FOR_ALL_OPERATIONS=configuration
  just tidy-client
  just fmt-client

setup-dev:
  curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.60.1
  go install github.com/segmentio/golines@latest

build:
  go install github.com/goreleaser/goreleaser/v2@latest
  PATH="$GOPATH/bin:$PATH" goreleaser release --clean

sign:
  hpesign --sign
test-client +ARGS='':
  cd {{ CONFIG_API_CLIENT_DIR }} && go test -v ./... -race -covermode=atomic -coverprofile=.coverage {{ ARGS }}

coverage-client:
  cd {{ CONFIG_API_CLIENT_DIR }} && go tool cover -html=.coverage -o=.coverage.html

fmt-client:
  gofmt -w {{ CONFIG_API_CLIENT_DIR }}
  golines -w {{ CONFIG_API_CLIENT_DIR }}

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

  output=$(golines . --dry-run)
  if [ -n "$output" ]; then
    echo "$output"
    echo "Error: (golines) formatting required" >&2
    exit 1
  fi

  golangci-lint run

fmt:
  gofmt -w .
  golines -w .

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
    TF_ACC=1 go test -v ./test/live/... -race -covermode=atomic -coverprofile=.coverage {{ ARGS }}
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

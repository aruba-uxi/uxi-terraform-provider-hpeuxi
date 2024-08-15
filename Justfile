TMP_DIR := "tmp"
CONFIG_API_CLIENT_DIR := "pkg/config-api-client"
CONFIG_API_PROVIDER_DIR := "pkg/config-api-provider"
OPENAPI_SPEC := "pkg/config-api-client/api"

retrieve-config-api-openapi-spec:
  rm -rf {{ TMP_DIR }}
  git clone git@github.com:aruba-uxi/configuration-api.git --depth=1 {{ TMP_DIR }}
  mkdir -p {{ OPENAPI_SPEC }}
  cp {{ TMP_DIR }}/oas/openapi.yaml {{ OPENAPI_SPEC }}/openapi.yaml
  rm -rf {{ TMP_DIR }}

generate-config-api-client: retrieve-config-api-openapi-spec
  docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli generate \
  --input-spec /local/{{ OPENAPI_SPEC }}/openapi.yaml \
  --generator-name go \
  --output /local/{{ CONFIG_API_CLIENT_DIR }} \
  --package-name config_api_client \
  --git-user-id aruba-uxi \
  --git-repo-id configuration-api-terraform-provider/{{ CONFIG_API_CLIENT_DIR }} \
  cd {{ CONFIG_API_CLIENT_DIR }} && go mod tidy
  just fmt-client

setup-dev:
  curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.60.1

test-client:
  cd {{ CONFIG_API_CLIENT_DIR }} && go test -v ./... -race -covermode=atomic -coverprofile=.coverage

coverage-client:
  cd {{ CONFIG_API_DIR }} && go tool cover -html=.coverage -o=.coverage.html

fmt-client:
  gofmt -w pkg/config-api-client

lint-client:
  #!/usr/bin/env bash

  cd pkg/config-api-client

  if [ -n "$(gofmt -d .)" ]; then
    echo "Error: (gofmt) formatting required" >&2
    exit 1
  fi

  golangci-lint run

lint-provider:
  #!/usr/bin/env bash

  cd pkg/config-api-provider

  if [ -n "$(gofmt -d .)" ]; then
    echo "Error: (gofmt) formatting required" >&2
    exit 1
  fi

  golangci-lint run

test:
  just test-client

coverage:
  just coverage-client

lint:
  just lint-client
  just lint-provider

fmt:
  just fmt-client

clean:
  find . -name ".coverage*" -type f -delete

plan:
  cd {{ CONFIG_API_PROVIDER_DIR }} && go install .
  cd {{ CONFIG_API_PROVIDER_DIR }}/examples/full-demo && terraform plan

apply:
  cd {{ CONFIG_API_PROVIDER_DIR }} && go install .
  cd {{ CONFIG_API_PROVIDER_DIR }}/examples/full-demo && terraform apply

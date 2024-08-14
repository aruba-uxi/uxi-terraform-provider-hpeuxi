TMP_DIR := "tmp"
CONFIG_API_DIR := "pkg/config-api-client"
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
  --output /local/{{ CONFIG_API_DIR }} \
  --package-name config_api_client \
  --git-user-id aruba-uxi \
  --git-repo-id configuration-api-terraform-provider/{{ CONFIG_API_DIR }} \
  cd {{ CONFIG_API_DIR }} && go mod tidy
  just fmt-client

setup-dev:
  curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.60.1

test-client:
  cd {{ CONFIG_API_DIR }} && go test -v ./... -race -covermode=atomic -coverprofile=.coverage

fmt-client:
  gofmt -w pkg/config-api-client

lint-client:
  #!/usr/bin/env bash

  if [ -n "$(gofmt -d pkg/config-api-client)" ]; then
    echo "Error: (gofmt) formatting required" >&2
    exit 1
  fi

test:
  just test-client

lint:
  just lint-client
  golangci-lint run

fmt:
  just fmt-client

clean:
  find . -name ".coverage" -type f -delete

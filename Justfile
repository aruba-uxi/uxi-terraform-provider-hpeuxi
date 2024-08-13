TMP_DIR := "tmp"
OPENAPI_SPEC := "pkg/config-api-client/api"

retrieve-config-api-openapi-spec:
    git clone git@github.com:aruba-uxi/configuration-api.git --depth=1 {{ TMP_DIR }}
    mkdir -p {{ OPENAPI_SPEC }}
    cp {{ TMP_DIR }}/oas/openapi.yaml {{ OPENAPI_SPEC }}/openapi.yaml
    rm -rf {{ TMP_DIR }}

generate-config-api-client: retrieve-config-api-openapi-spec
    docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli generate \
    -i /local/{{ OPENAPI_SPEC }}/openapi.yaml \
    -g go \
    -o /local/pkg/config-api-client
    cd pkg/config-api-client && go mod tidy

test-client:
    cd pkg/config-api-client && go test -v ./...

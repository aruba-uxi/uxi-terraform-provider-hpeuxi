TMP_DIR := "config-api-client/tmp/configuration-api"
OPENAPI_SPEC := "config-api-client/tmp/openapi.yaml"

retrieve-config-api-openapi-spec:
    git clone git@github.com:aruba-uxi/configuration-api.git --depth=1 {{ TMP_DIR }}
    cp {{ TMP_DIR }}/oas/openapi.yaml {{ OPENAPI_SPEC }}
    rm -rf {{ TMP_DIR }}

generate-config-api-client: retrieve-config-api-openapi-spec
    docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli generate \
    -i /local/{{ OPENAPI_SPEC }} \
    -g go \
    -o /local/config-api-client

test-client:
    cd config-api-client && go test -v ./...

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
  --additional-properties=withGoMod=false
  cd {{ CONFIG_API_DIR }} && go mod tidy

test-client:
  cd {{ CONFIG_API_DIR }} && go test -v ./...

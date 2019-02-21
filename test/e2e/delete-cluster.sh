#!/bin/bash

CLUSTER_API_HTTP=${CLUSTER_API_HTTP:-http}
CLUSTER_API=${CLUSTER_API:-cluster-manager-api-cluster-manager-api}
CLUSTER_API_PORT=${CLUSTER_API_PORT:-80}
CMA_CALLBACK_URL=${CMA_CALLBACK_URL:-https://example.cnct.io}
CMA_CALLBACK_REQUESTID=${CMA_CALLBACK_REQUESTID:-1}

[[ -n $DEBUG ]] && set -o xtrace
set -o errexit
set -o nounset
set -o pipefail

required_envs(){
  if [[ -z "${CLUSTER_NAME+x}" ]];then
            echo >&2 "Cannot continue. \$CLUSTER_NAME is not set."
            exit 1
  fi
  if [[ -z "${AZURE_CLIENT_ID+x}" ]];then
            echo >&2 "Cannot continue. \$AZURE_CLIENT_ID is not set."
            exit 1
  fi
  if [[ -z "${AZURE_CLIENT_SECRET+x}" ]];then
            echo >&2 "Cannot continue. \$AZURE_CLIENT_SECRET is not set."
            exit 1
  fi
  if [[ -z "${AZURE_TENANT_ID+x}" ]];then
            echo >&2 "Cannot continue. \$AZURE_TENANT_ID is not set."
            exit 1
  fi
  if [[ -z "${AZURE_SUBSCRIPTION_ID+x}" ]];then
            echo >&2 "Cannot continue. \$AZURE_SUBSCRIPTION_ID is not set."
            exit 1
  fi
}

main() {
  required_envs

  curl -X DELETE \
    "${CLUSTER_API_HTTP}://${CLUSTER_API}:${CLUSTER_API_PORT}/api/v1/cluster?name=${CLUSTER_NAME}&provider=azure&azure.app_id=${AZURE_CLIENT_ID}&azure.tenant=${AZURE_TENANT_ID}&azure.password=${AZURE_CLIENT_SECRET}&azure.subscription_id=${AZURE_SUBSCRIPTION_ID}&callback.url=${CMA_CALLBACK_URL}&callback.request_id=${CMA_CALLBACK_REQUESTID}" \
    -H 'Cache-Control: no-cache' \
    -H 'Content-Type: application/json' \
    -iks
}

main

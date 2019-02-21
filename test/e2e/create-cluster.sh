#!/bin/bash

CLUSTER_API_HTTP=${CLUSTER_API_HTTP:-http}
CLUSTER_API=${CLUSTER_API:-cluster-manager-api-cluster-manager-api}
CLUSTER_API_PORT=${CLUSTER_API_PORT:-80}
K8S_VERSION=${K8S_VERSION:-1.12.5}
CMA_CALLBACK_URL=${CMA_CALLBACK_URL:-https://example.cnct.io}
CMA_CALLBACK_REQUESTID=${CMA_CALLBACK_REQUESTID:-1}

# azure specific inputs
AZURE_LOCATION=${AZURE_LOCATION:-westus}
AZURE_NODE_SIZE=${AZURE_NODE_SIZE:-Standard_A1}

[[ -n $DEBUG ]] && set -o xtrace
set -o errexit
set -o nounset
set -o pipefail

DATA=$(
  cat <<JSON
{
  "name": "${CLUSTER_NAME}",
  "provider": {
    "name": "azure",
    "k8s_version": "${K8S_VERSION}",
    "azure": {
      "location": "${AZURE_LOCATION}",
      "credentials": {
        "app_id": "${AZURE_CLIENT_ID}",
        "tenant": "${AZURE_TENANT_ID}",
        "password": "${AZURE_CLIENT_SECRET}",
        "subscription_id": "${AZURE_SUBSCRIPTION_ID}"
      },
      "clusterAccount": {
        "client_id": "${AZURE_CLIENT_ID}",
        "client_secret": "${AZURE_CLIENT_SECRET}"
      },
      "instance_groups": [
        {
          "name": "agentpool1",
          "type": "${AZURE_NODE_SIZE}",
          "min_quantity": 1
        }
      ]
    },
    "high_availability": true,
    "network_fabric": ""
  },
  "callback": {
    "url": "${CMA_CALLBACK_URL}",
    "request_id": "${CMA_CALLBACK_REQUESTID}" 
  }

}
JSON
)

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

  curl -X POST \
    "${CLUSTER_API_HTTP}://${CLUSTER_API}:${CLUSTER_API_PORT}/api/v1/cluster" \
    -H 'Cache-Control: no-cache' \
    -H 'Content-Type: application/json' \
    -d "${DATA}" \
    -iks
}

main

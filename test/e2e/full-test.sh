#!/bin/bash

__dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

export CLUSTER_API_HTTP=${CLUSTER_API_HTTP:-http}
export CLUSTER_API=${CLUSTER_API:-cluster-manager-api-cluster-manager-api}
export CLUSTER_API_PORT=${CLUSTER_API_PORT:-80}
export CLUSTER_NAME=${CLUSTER_NAME:-jenkins-test-$(date +%s)}
export CLUSTER_API_NAMESPACE=${CLUSTER_API_NAMESPACE:-cma}
export K8S_VERSION=${K8S_VERSION:-1.12.5}
export CMA_CALLBACK_URL=${CMA_CALLBACK_URL:-https://example.cnct.io}
export CMA_CALLBACK_REQUESTID=${CMA_CALLBACK_REQUESTID:-1}

# azure specific inputs
export AZURE_LOCATION=${AZURE_LOCATION:-westus}
export AZURE_NODE_SIZE=${AZURE_NODE_SIZE:-Standard_A1}

[[ -n $DEBUG ]] && set -o xtrace
set -o errexit
set -o nounset
set -o pipefail

readonly CLIENT_KUBECONFIG="$CLUSTER_NAME-kubeconfig.yaml"

required_envs(){
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

get_kubeconfig(){
  "${__dir}/get-kubeconfig.sh" > "${CLIENT_KUBECONFIG}"
}

test_provisioning(){
  provisioning=$("${__dir}/create-cluster.sh")
  echo "create output:"
  echo $provisioning
  if echo "$provisioning" | grep -o PROVISIONING; then
    echo "Cluster is PROVISIONING"
  else
    echo "Cluster is NOT PROVISIONING"
    return 1
  fi
  return 0
}

test_running(){
  # wait up to 20 minutes for cluster RUNNING
  for tries in $(seq 1 120); do
    running=$("${__dir}/get-cluster.sh")

    if echo $running | grep -o RUNNING; then
      echo "Cluster is RUNNING"
      runningstatus="PASS"
      echo "elapsed seconds=$(( 10 * $tries ))"
      break
    else
      echo "Cluster is NOT RUNNING"
    fi
    sleep 10
  done

  if [ -z ${runningstatus+x} ]; then
    echo "Timed out waiting for RUNNING status"
    return 1
  fi
  return 0
}

test_ready(){
  get_kubeconfig

  nodes=$(kubectl get nodes -o wide --kubeconfig "$CLIENT_KUBECONFIG")
  echo $nodes

  rm "$CLIENT_KUBECONFIG"

  # check for not ready
  if echo $nodes | grep -o NotReady; then
    echo "Node(s) NotReady"
    return 1
  fi

  if echo $nodes | grep -o SchedulingDisabled; then
    echo "Node(s) SchedulingDisabled"
    return 1
  fi

  echo "Nodes READY"
  return 0
}

test_delete(){
  delete=$("${__dir}/delete-cluster.sh")
  echo "delete output:"
  echo $delete

  # wait up to 20 minutes for cluster delete complete
  for tries in $(seq 1 120); do
    deleted=$("${__dir}/get-cluster.sh")

    if echo $deleted | grep -o "ResourceNotFound"; then
      echo "Cluster DELETE is COMPLETE"
      deletedstatus="PASS"
      echo "elapsed seconds=$(( 10 * $tries ))"
      break
    else
      echo "Cluster DELETE is NOT COMPLETE"
    fi
    sleep 10
  done

  if [ -z ${deletedstatus+x} ]; then
    echo "Timed out waiting for DELETE to finish"
    return 1
  fi
  return 0
}


main() {
  required_envs
  fullstatus="PASSED"

  # test create is provisioning
  if ! test_provisioning; then
     echo "test_provisioning FAILED"
  else
     echo "test_provisioning PASSED"
  fi

  if ! test_running; then
    echo "test_running FAILED"
    fullstatus="FAILED"
  else
    echo "test_running PASSED"
  fi

  if ! test_ready; then
    echo "test_ready FAILED"
    fullstatus="FAILED"
  else
    echo "test_ready PASSED"
  fi

  if ! test_delete; then
    echo "test_delete FAILED"
    fullstatus="FAILED"
  else
    echo "test_delete PASSED"
  fi

  echo "full-test $fullstatus"
  if [ "$fullstatus" == "FAILED" ]; then
    exit 1
  fi

  exit 0
}

main

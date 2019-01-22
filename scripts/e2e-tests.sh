#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail
set -o xtrace

function finish {
  kind delete cluster
}
trap finish EXIT

echo "The current environment contains these variables: $(env)"
echo "The current directory is $(pwd)"

service docker start

kind create cluster --wait=10m --loglevel=debug

export KUBECONFIG=$(kind get kubeconfig-path)

kubectl create clusterrolebinding superpowers --clusterrole=cluster-admin --user=system:serviceaccount:kube-system:default

echo "The current pods are:"
kubectl get pods --all-namespaces
kubectl describe pods --all-namespaces

helm init --wait --debug

helm install --name my-release stable/cert-manager --wait --debug
helm install --name nginx-ingress stable/nginx-ingress --wait --debug

# TODO: This is incorrect. We should be installing the repo containing the PR we 
# wish to test.
go get github.com/samsung-cnct/cma-aks
cd ${GOPATH}/src/github.com/samsung-cnct/cma-aks
helm install --name cma-aks deployments/helm/cma-aks/ --wait --debug

helm repo add cnct https://charts.migrations.cnct.io
helm install --name cluster-manager-api cnct/cluster-manager-api --wait --debug
helm install --name cma-operator cnct/cma-operator --wait --debug

# TODO: Add tests here

#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail
set -o xtrace

service docker start

kind create cluster --retain --wait=1m --loglevel=debug

export KUBECONFIG=$(kind get kubeconfig-path)

kubectl create clusterrolebinding superpowers --clusterrole=cluster-admin --user=system:serviceaccount:kube-system:default

helm init --wait

helm install --name my-release stable/cert-manager --wait
helm install --name nginx-ingress stable/nginx-ingress --wait

# TODO: This is incorrect. We should be installing the repo containing the PR we 
# wish to test.
go get github.com/samsung-cnct/cma-aks
cd ${GOPATH}/src/github.com/samsung-cnct/cma-aks
helm install --name cma-aks deployments/helm/cma-aks/ --wait

helm repo add cnct https://charts.migrations.cnct.io
helm install --name cluster-manager-api cnct/cluster-manager-api --wait
helm install --name cma-operator cnct/cma-operator --wait

# TODO: Add tests here

kind delete cluster

#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail
set -o xtrace

function finish {
  kind delete cluster
}
trap finish EXIT

kind create cluster --wait=10m --loglevel=debug
export KUBECONFIG=$(kind get kubeconfig-path)

kubectl create clusterrolebinding superpowers --clusterrole=cluster-admin --user=system:serviceaccount:kube-system:default

helm init -c
export HELM_HOST=localhost:44134
helm plugin install https://github.com/rimusz/helm-tiller || true
helm tiller start-ci

# install cert-manager (required by cma-aks chart)
# Note: removed --wait, it times out downloading the .tgz file
kubectl apply \
    -f https://raw.githubusercontent.com/jetstack/cert-manager/release-0.6/deploy/manifests/00-crds.yaml
kubectl create namespace cert-manager
kubectl label namespace cert-manager certmanager.k8s.io/disable-validation=true
helm repo update
helm install --name cert-manager --namespace cert-manager stable/cert-manager
sleep 20
helm install --name nginx-ingress stable/nginx-ingress

helm repo add cnct https://charts.cnct.io
helm install --name cma-aks --set image.repo=quay.io/samsung_cnct/cma-aks:test cnct/cma-aks
helm install -f test/e2e/cma-values.yaml --name cluster-manager-api cnct/cluster-manager-api
helm install -f test/e2e/cma-operator-values.yaml --name cma-operator cnct/cma-operator

helm tiller stop

# copy test scripts to kind container
docker cp test/e2e/ kind-control-plane:/root/
# create kubernetes job to run tests
apk add gettext
envsubst < test/e2e/run-tests-job.yaml | kubectl apply -f -
# wait for tests to complete TODO: adjust timeout as necessary
kubectl wait --for=condition=complete job/cma-aks-e2e-tests --timeout=30m
# output logs after job completes
kubectl logs job/cma-aks-e2e-tests -n pipeline-tools

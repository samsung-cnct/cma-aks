# CMA Azure AKS Helper API
[![Build Status](https://jenkins.cnct.io/buildStatus/icon?job=cma-aks/master)](https://jenkins.cnct.io/job/cma-aks/job/master/)

## Overview

The cma-aks repo provides a helper API for [cluster-manager-api](https://github.com/samsung-cnct/cluster-manager-api) by talking to Microsoft Azures API to interact with Azure Kubernetes Services (AKS) clusters.

## Getting started

See [Protocol Documentation](https://github.com/samsung-cnct/cma-aks/blob/master/docs/api-generated/api.md)
- [open api in swagger editor](https://editor.swagger.io/?url=https://raw.githubusercontent.com/samsung-cnct/cma-aks/master/assets/generated/swagger/api.swagger.json)
- [open api in swagger ui](http://petstore.swagger.io/?url=https://raw.githubusercontent.com/samsung-cnct/cma-aks/master/assets/generated/swagger/api.swagger.json)


### Requirements
- Kubernetes 1.7+
- [nginx-ingress](https://github.com/helm/charts/tree/master/stable/nginx-ingress)
- [cert-manager](https://github.com/helm/charts/tree/master/stable/cert-manager)

### Deploy
```bash
$ helm install deployments/helm/cma-aks --name cma-aks
```

## Contributing
Utilizes:
- [azure-sdk-for-go](https://github.com/Azure/azure-sdk-for-go)
- [Protocol Buffers](https://developers.google.com/protocol-buffers/)

To generate code:
```bash
$ make -f build/Makefile generators
```

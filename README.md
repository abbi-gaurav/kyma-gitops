# Overview

Explore Gitops for developers building stuff on [project "kyma"](https://kyma-project.io) using a pull based approach. The approach is based on the [open source project flux](https://www.weave.works/oss/flux/).

## Setup

* Create `flux` namespace

    ```shell
    kubectl apply -f setup/ns.yaml
    ```

* Install Flux as a Helm chart

    ```shell
    helm repo add fluxcd https://charts.fluxcd.io
    kubectl apply -f https://raw.githubusercontent.com/fluxcd/flux/helm-0.10.1/deploy-helm/flux-helm-release-crd.yaml
    helm upgrade -i flux \
    --set helmOperator.create=true \
    --set helmOperator.createCRD=false \
    --set git.url=git@github.com:abbi-gaurav/kyma-gitops \
    --set git.path=resources \
    --set git.pollInterval=2m \
    --namespace flux \
    fluxcd/flux --tls
    ```

* Give write access for repository.

  * Run command
  
    ```shell
    fluxctl identity --k8s-fwd-ns flux
    ```
  
  * In the repository, go to Setting > Deploy keys, click on Add deploy key, give it a Title, check Allow write access, paste the Flux public key and click Add key.

## Various flows

* [Lambda deployment](code/lambdas/README.md)
# Overview

Explore Gitops for developers building stuff on [project "kyma"](https://kyma-project.io).

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
    --set git.path=resources
    --namespace flux \
    fluxcd/flux --tls
    ```

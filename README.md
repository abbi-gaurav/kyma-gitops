# Overview

Explore Gitops for developers building stuff on [project "kyma"](https://kyma-project.io) using a pull based approach. The approach is based on the [open source project flux](https://www.weave.works/oss/flux/).

## Prerequisites

* [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/) CLI referncing to the Kubernetes cluster on which Kyma has been installed.
* Install [Helm](https://helm.sh/)
* [Add certificates to Helm Home](https://kyma-project.io/docs/components/security/#details-tls-in-tiller-add-certificates-to-helm-home)

## Setup

* Create `flux` namespace

    ```shell
    kubectl apply -f setup/ns.yaml
    ```

* Install Flux as a Helm chart

    ```shell
    helm repo add fluxcd https://charts.fluxcd.io
    helm upgrade -i flux \
    --set helmOperator.create=false \
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
[![CircleCI](https://circleci.com/gh/abbi-gaurav/kyma-gitops/tree/dev.svg?style=svg)](https://circleci.com/gh/abbi-gaurav/kyma-gitops/tree/dev)

# Overview

Explore Gitops for developers building stuff on [project "kyma"](https://kyma-project.io) using a pull based approach. The approach is based on the [open source project flux](https://www.weave.works/oss/flux/).

## Prerequisites

* [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/) CLI referncing to the Kubernetes cluster on which Kyma has been installed.
* Install [Helm](https://helm.sh/)
* [Add certificates to Helm Home](https://kyma-project.io/docs/components/security/#details-tls-in-tiller-add-certificates-to-helm-home)

## Cluster level setup

This *example setup* shows using Kyma with Flux to deploy cluster level as well as a dev namespace. The idea can be extended to multiple namespaces.

**Kyma Cluster Admin Flow**

* Admins will work on the [master branch](https://github.com/abbi-gaurav/kyma-gitops).
* Push any cluster level changes to master.
* Flux will pick up the changes and deploy them to cluster.

**Kyma Developer flow**

* Developers will work on the [dev branch](https://github.com/abbi-gaurav/kyma-gitops/tree/dev).
* Write code and configuratations.
* Push them to dev branch.
* Flux will pick up the changes and deploy them to dev namespace.

> Note: This can be extended to multiple namespaces.

![Kyma Flux Flow](assets/Kyma&#32;Flux&#32;flow.png)

* Create `flux` namespace

    ```shell
    kubectl apply -f setup/ns.yaml
    ```

* Install Flux as a Helm chart for cluster level resources

    ```shell
    helm repo add fluxcd https://charts.fluxcd.io
    helm upgrade -i flux \
    --set helmOperator.create=false \
    --set helmOperator.createCRD=false \
    --set git.url=git@github.com:abbi-gaurav/kyma-gitops \
    --set git.path=cluster-resources \
    --set git.pollInterval=2m \
    --set registry.excludeImage="*" \
    --set syncGarbageCollection.enabled=true \
    --set syncGarbageCollection.dry=false \
    --namespace flux \
    fluxcd/flux --tls
    ```

* Give write access for repository.

  * Run command
  
    ```shell
    fluxctl identity --k8s-fwd-ns=flux
    ```
  
  * In the repository, go to Setting > Deploy keys, click on Add deploy key, give it a Title, check Allow write access, paste the Flux public key and click Add key.

## Dev branch setup

* Create a `dev` branch from the Kyma console.

* Install Flux as a Helm chart for development resources.

    ```shell
    helm upgrade -i dev-flux \
    --set helmOperator.create=false \
    --set helmOperator.createCRD=false \
    --set git.url=git@github.com:abbi-gaurav/kyma-gitops \
    --set git.branch=dev \
    --set git.path=deployment \
    --set git.pollInterval=2m \
    --set registry.excludeImage="*" \
    --set clusterRole.create=false \
    --set syncGarbageCollection.enabled=true \
    --set syncGarbageCollection.dry=false \
    --namespace dev \
    fluxcd/flux --tls
    ```

* Give write access for repository.

  * Run command
  
    ```shell
    fluxctl identity --k8s-fwd-ns=dev
    ```
  
  * In the repository, go to Setting > Deploy keys, click on Add deploy key, give it a Title, check Allow write access, paste the Flux public key and click Add key.

## Various flows

* [Lambda deployment](code/lambdas/README.md)
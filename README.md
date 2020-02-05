# Overview

Explore Gitops for developers building stuff on [project "kyma"](https://kyma-project.io) using a pull based approach. The approach is based on the [open source project flux](https://www.weave.works/oss/flux/).

## Prerequisites

* [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/) CLI referncing to the Kubernetes cluster on which Kyma has been installed.
* Install [Helm](https://helm.sh/)
* [Add certificates to Helm Home](https://kyma-project.io/docs/components/security/#details-tls-in-tiller-add-certificates-to-helm-home)

## Installing flux

Flux will be the component that will run inside Kyma and continuously watch for any changes in the github repository and apply them to the Kyma cluster.

This step will be done by the admin of the Kyma cluster.

This command will install the flux operator in the `dev namespace` for the github branch `master`. The deployment path specified is directory called `deployment`. You can read more about various configuration options in the [flux documentation](https://github.com/fluxcd/flux/tree/master/chart/flux#configuration).

Any commits to the master branch in the deployment directory will be applied to the Kyma cluster.

* Install Flux as a helm chart.

```shell
  helm upgrade -i dev-flux \
  --set helmOperator.create=false \
  --set helmOperator.createCRD=false \
  --set git.url=git@github.com:abbi-gaurav/kyma-gitops \
  --set git.branch=master \
  --set git.path=extensions \
  --set git.pollInterval=2m \
  --set git.readonly=true \
  --set registry.excludeImage="*" \
  --set clusterRole.create=false \
  --set syncGarbageCollection.enabled=true \
  --set syncGarbageCollection.dry=false \
  --namespace dev \
  fluxcd/flux --tls
```

* Give write access for github repository.

  * Run command
  
    ```shell
    fluxctl identity --k8s-fwd-ns=flux
    ```
  
  * In the repository, go to `Setting > Deploy keys`
  * Click on `Add deploy key`, give it a Title, check Allow write access, paste the Flux public key and click Add key.

## Kyma Developer flow

Let's say our devlopment team is working in the dev namespace. This is where all their microservices, lambdas and other resources will be deployed.

* Developers will work on the [git hub repository](https://github.com/abbi-gaurav/kyma-gitops.git).
* Write code and configuratations.
* Push them to dev branch.
* Flux will pick up the changes and deploy them to dev namespace.

> Note: This can be extended to multiple namespaces.

![Kyma Flux Flow](assets/Kyma&#32;Flux&#32;flow&#32;Dev.png)

## GitOps Examples

* [Java Microservice](extensions/java/my-java-extension/README.md)
* [Scala Microservice](extensions/scala/my-scala-extension/README.md)
* [Go Microservice](extensions/go/my-go-extension/README.md)
* [Kyma Lambda](extensions/lambdas/sample-lambda/README.md)

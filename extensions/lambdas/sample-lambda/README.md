# Overview

Sample Kyma lambda deployed to a Kyma cluster using GitOps. It is exposed as a Kyma API over internet.

## Developer Flow

* Generate Kubernetes resources to deploy on Kyma and expose API

```shell script
make create-function
make expose-function
```

* Push the changes to the repository

```shell script
git add . && git commit -m "my Kyma lambda" && git push origin master
```

## GitOps

* Flux running inside Kyma cluster will pick up the git commits from the repo and apply it to Kyma cluster.

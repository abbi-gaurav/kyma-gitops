#!/usr/bin/env bash

API_FILE=$1
NAME=$2
BRANCH=$3

cat <<EOF > "${API_FILE}"
apiVersion: gateway.kyma-project.io/v1alpha2
kind: Api
metadata:
  name: ${NAME}-${BRANCH}
spec:
  authentication: []
  hostname: ${NAME}-${BRANCH}
  service:
    name: ${NAME}
    port: 8080
EOF
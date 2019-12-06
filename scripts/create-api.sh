#!/usr/bin/env bash

API_FILE=$1
NAME=$2

cat <<EOF > "${API_FILE}"
apiVersion: gateway.kyma-project.io/v1alpha2
kind: Api
metadata:
  name: ${NAME}
spec:
  authentication: []
  hostname: ${NAME}
  service:
    name: ${NAME}
    port: 8080
EOF
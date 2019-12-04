#!/usr/bin/env bash

API_FILE=$1
NAME=$2
NS=$3
DOMAIN=$4

cat <<EOF > "${API_FILE}"
apiVersion: gateway.kyma-project.io/v1alpha2
kind: Api
metadata:
  name: ${NAME}
  namespace: ${NS}
spec:
  authentication: []
  hostname: ${NAME}.${DOMAIN}
  service:
    name: ${NAME}
    port: 8080
EOF
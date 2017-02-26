#!/bin/bash

set -e

# This script is used during docker-compose/tests to setup authorization user on fresh etcd

ETCD_HOST_PORT="${ETCD_HOST:?required}"
if [[ "${ETCD_PORT:-X}" != "X" ]]; then
  ETCD_HOST_PORT="${ETCD_HOST_PORT}:${ETCD_PORT}"
fi
if [[ "${ETCD_PASSWORD}X" != "X" ]]; then
  if [[ "$(curl -s ${ETCD_HOST_PORT}/v2/auth/users |  jq -r .message)" != "Insufficient credentials" ]]; then
    if [[ "$(curl -s ${ETCD_HOST_PORT}/v2/auth/users |  jq -r .users)" == "null" ]]; then
      echo "Creating missing root user..."
      curl -s ${ETCD_HOST_PORT}/v2/auth/users/root -X PUT -d "{\"user\":\"${ETCD_USERNAME:-root}\",\"password\":\"${ETCD_PASSWORD:?required}\"}"
    fi

    auth_enabled=$(curl -s ${ETCD_HOST_PORT}/v2/auth/enable | jq -r .enabled)
    if [[ "${auth_enabled}" != "true" ]]; then
      echo "Enabling etcd authorization..."
      curl -s ${ETCD_HOST_PORT}/v2/auth/enable -X PUT
    fi
  fi

  echo
  echo "Verifying crdentials..."
  curl -s -u ${ETCD_USERNAME:-root}:${ETCD_PASSWORD:?required} ${ETCD_HOST_PORT}/v2/auth/users
else
  echo "Verifying that no credentials are required to ${ETCD_HOST_PORT:?required}..."
  curl -v ${ETCD_HOST_PORT:?}/v2/keys
  if [[ "$(curl -s ${ETCD_HOST_PORT}/v2/auth/users |  jq -r .message)" == "Insufficient credentials" ]]; then
    echo "Etcd credentials are required"
    exit 1
  fi
fi

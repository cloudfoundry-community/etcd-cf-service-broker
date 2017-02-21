#!/bin/sh

set -e

echo Running broker
env | sort

/scripts/initialize_etcd_auth.sh
etcd-cf-service-broker

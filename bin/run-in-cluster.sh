#!/usr/bin/env bash

set -euxo pipefail

source "$(dirname "$0")/common.sh"

start_in_cluster
kubectl logs -f search-autocomplete-local-run

kill $KUBECTL_PROXY_PID

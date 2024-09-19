#!/usr/bin/env bash

set -euxo pipefail

source "$(dirname "$0")/common.sh"

start_in_cluster

pushd "$PROJECT_ROOT/bruno"
bru run
popd

kill $KUBECTL_PROXY_PID

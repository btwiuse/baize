#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

currentDir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
rootDir="$currentDir"/..
set_workspace_env=$("$rootDir"/print-workspace-status.sh | sed -E "s/^([^ ]+) (.*)\$$/export \\1=\\2/g")
readonly set_workspace_env
$set_workspace_env

cd "$rootDir"

docker build -t btwiuse/ubuntu:executor-base . -f base.Dockerfile
docker build -t btwiuse/baize-executor:"${STABLE_DOCKER_TAG}" . -f executor.Dockerfile
docker build -t btwiuse/baize-server:"${STABLE_DOCKER_TAG}" . -f server.Dockerfile
docker build -t btwiuse/remote-cache:"${STABLE_DOCKER_TAG}" . -f remote-cache.Dockerfile

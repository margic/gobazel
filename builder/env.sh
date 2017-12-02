#!/bin/sh

echo "Checking for DOCKER_REGISTRY_PASSWORD env var"

# can't have the password open on env vars. set it in credential helper and unset.
if [ "$DOCKER_REGISTRY_PASSWORD" ] ; then
    echo "DOCKER_REGISTRY_PASSWORD set, adding password to credential helper"
    echo "{\"ServerURL\":\"${DOCKER_REGISTRY_HOST}\",\"Username\":\"${DOCKER_REGISTRY_USER}\",\"Secret\":\"${DOCKER_REGISTRY_PASSWORD}\"}" | docker-credential-pass store
    unset DOCKER_REGISTRY_PASSWORD
else
    echo "DOCKER_REGISTRY_PASSWORD not set."
fi

echo "Checking for DOCKER_REGISTRY_PASSWORD done"


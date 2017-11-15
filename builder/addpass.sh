#!/bin/bash

echo "{\"ServerURL\":\"${DOCKER_REGISTRY_HOST}\",\"Username\":\"${DOCKER_REGISTRY_USER}\",\"Secret\":\"${DOCKER_REGISTRY_PASSWORD}\"}" | docker-credential-pass store
unset DOCKER_REGISTRY_PASSWORD

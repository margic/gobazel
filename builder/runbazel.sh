#!/bin/sh

# TODO THIS DOESN'T DO WHAT I INTENDED
# see https://github.com/margic/gobazel/issues/1
. /opt/gobazel/addpass.sh


# to make it easier to reuse the container and therefore the bazel cache
# this script will read the bazel command from stdin and execute the bazel
# command with that
read line
echo "command: $line"
bazel $line


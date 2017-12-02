#!/bin/sh

# TODO THIS DOESN'T DO WHAT I INTENDED AND CAUSES COMPLICATED RUN COMMAND
# see https://github.com/margic/gobazel/issues/1
. /opt/gobazel/env.sh


# to make it easier to reuse the container and therefore the bazel cache
# this script will read the bazel command from stdin and execute the bazel
# command with that
read line
echo "bazel command: $line"
bazel $line


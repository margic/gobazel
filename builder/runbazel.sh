#!/bin/sh

. /opt/gobazel/addpass.sh


# was going to use this to copy .kube to ~ but seems it works from source folder
# if [ ! -d "/root/.kube" ]; then
#     if [ -d "/source/.kube"]; then
#         cp -r /source/.kube /root/.kube
#     fi
# fi

# on start bazel check if there is a bazelcmd file in the /source volume
# doing this so it's possible to override the command running in the container
# allowin multiple reuse of a container to run the bazel command
# This allows us to take advantage of bazel caching rather than running a new
# container each time. If bazelcmd file is present bazel will run with that
# command else it will use the command from the command line args
if [ -e /source/bazelcmd ] ; then
    echo "/source/buildcmd found running bazel $(cat /source/bazelcmd)"
    bazel `cat /source/bazelcmd` 
else
    echo "/source/bazelcmd not found running bazel $($@)"
    bazel "$@"
fi


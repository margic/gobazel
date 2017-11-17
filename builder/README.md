I've set the project up for building with bazel.
Unfortunately building with bazel on mac cross compiling for linux based docker
currently does not work with some of the main packages used (such as viper).
I expect this will be fixed but for the time being the best way to build
the project is in a linux container.

The builder docker images is provided to fill the gap and provide a native
linux build environment for the project.

It's intended to be used with my [gobazel demo project](https://github.com/margic/gobazel").

Docker Credential Helper
The docker credential helper is provided to make enable login to index.docker.io
using the linux pass package installed on the image. Configuration is provided in
the ~/.docker/config.json

Passing docker creds
When starting the builder a container pass DOCKER_REGISTRY_HOST, DOCKER_REGSITRY_USER,
DOCKER_REGSITRY_PASSWORD in. DON'T DO THIS IN CLOUD CONTAINERS. 



I've set the project up for building with bazel.
Unfortunately building with bazel on mac cross compiling for linux based docker
currently does not work with some of the main packages used (such as viper).
I expect this will be fixed but for the time being the best way to build
the project is in a linux container.

The builder docker images is provided to fill the gap and provide a native
linux build environment for the project.
# gobazel
Trying out bazel with go microservices.

Clone repo
Run `dep ensure` to create vendor folder

Use builder image to build targets

Bazel requires a build file in each package. When adding dependencies to the vendor
folder using `dep ensure -add pkg` newly added packages will not have the build file
by default. Run `bazel run //:gazelle` to automatically add build files to vendor
pachages.
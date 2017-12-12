# Go rules
git_repository(
    name = "io_bazel_rules_go",
    remote = "https://github.com/bazelbuild/rules_go.git",
    commit = "44941765617a5040d4dbb96966073180e2d70f42",
)
load("@io_bazel_rules_go//go:def.bzl", "go_rules_dependencies", "go_register_toolchains")
go_rules_dependencies()
go_register_toolchains()

#Docker rules
git_repository(
    name = "io_bazel_rules_docker",
    remote = "https://github.com/margic/rules_docker.git",
    commit = "d875857ec4c7a564aa6ff1456dd5d00e1c89312f",
)

load(
    "@io_bazel_rules_docker//container:container.bzl",
    "container_push",
    container_repositories = "repositories",
)

load(
    "@io_bazel_rules_docker//go:image.bzl",
    _go_image_repos = "repositories",
)

_go_image_repos()
container_repositories()

# This requires rules_docker to be fully instantiated before
# it is pulled in.
git_repository(
    name = "io_bazel_rules_k8s",
    commit = "8240d175e08b3e4c2a1f3d6038d33800fb1cd692",
    remote = "https://github.com/bazelbuild/rules_k8s.git",
)

load("@io_bazel_rules_k8s//k8s:k8s.bzl", "k8s_repositories", "k8s_defaults")
k8s_repositories()

k8s_defaults(
  # This becomes the name of the @repository and the rule
  # you will import in your BUILD files.
  name = "k8s_deploy",
  kind = "deployment",
  # This is the name of the cluster as it appears in:
  #   kubectl config current-context
  cluster = "minikube",
  namespace = "gobazel",
)
# gazelle:repository_macro deps.bzl%go_dependencies
workspace(name = "com_github_uhthomas_tesladump")

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "bazel_gazelle",
    sha256 = "222e49f034ca7a1d1231422cdb67066b885819885c356673cb1f72f748a3c9d4",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.22.3/bazel-gazelle-v0.22.3.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.22.3/bazel-gazelle-v0.22.3.tar.gz",
    ],
)

http_archive(
    name = "com_github_bazelbuild_buildtools",
    sha256 = "a02ba93b96a8151b5d8d3466580f6c1f7e77212c4eb181cba53eb2cae7752a23",
    strip_prefix = "buildtools-3.5.0",
    url = "https://github.com/bazelbuild/buildtools/archive/3.5.0.tar.gz",
)

http_archive(
    name = "rules_proto",
    sha256 = "d8992e6eeec276d49f1d4e63cfa05bbed6d4a26cfe6ca63c972827a0d141ea3b",
    strip_prefix = "rules_proto-cfdc2fa31879c0aebe31ce7702b1a9c8a4be02d2",
    urls = ["https://github.com/bazelbuild/rules_proto/archive/cfdc2fa31879c0aebe31ce7702b1a9c8a4be02d2.tar.gz"],
)

http_archive(
    name = "rules_python",
    sha256 = "b6d46438523a3ec0f3cead544190ee13223a52f6a6765a29eae7b7cc24cc83a0",
    urls = ["https://github.com/bazelbuild/rules_python/releases/download/0.1.0/rules_python-0.1.0.tar.gz"],
)

# pin to e15c9ebf203b7fa708e69ff5f1cdcf427d7edf6f, as sequential container_push
# is not available in v0.15.0
http_archive(
    name = "io_bazel_rules_docker",
    sha256 = "893726fd83049cece8bfec873091c43877449f28987adf5c13b17801ecf0a788",
    strip_prefix = "rules_docker-e15c9ebf203b7fa708e69ff5f1cdcf427d7edf6f",
    urls = ["https://github.com/bazelbuild/rules_docker/archive/e15c9ebf203b7fa708e69ff5f1cdcf427d7edf6f.tar.gz"],
)

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "7904dbecbaffd068651916dce77ff3437679f9d20e1a7956bff43826e7645fcc",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.25.1/rules_go-v0.25.1.tar.gz",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.25.1/rules_go-v0.25.1.tar.gz",
    ],
)

load("@rules_python//python:pip.bzl", "pip_repositories")

pip_repositories()

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

go_rules_dependencies()

go_register_toolchains(
    go_version = "1.15.8",
    nogo = "@//:nogo",
)

load("//:deps.bzl", "go_dependencies")

go_dependencies()

load("@rules_proto//proto:repositories.bzl", "rules_proto_dependencies", "rules_proto_toolchains")

rules_proto_dependencies()

rules_proto_toolchains()

load("@io_bazel_rules_docker//repositories:repositories.bzl", container_repositories = "repositories")

container_repositories()

load("@io_bazel_rules_docker//repositories:deps.bzl", container_deps = "deps")

container_deps()

load("@io_bazel_rules_docker//go:image.bzl", _go_image_repos = "repositories")

_go_image_repos()

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

gazelle_dependencies()

# http://bazel.build/
# vim: set ft=python sts=2 sw=2 et:

workspace(name = "com_github_chronostachyon_playground")

git_repository(
    name = "io_bazel_rules_go",
    remote = "https://github.com/bazelbuild/rules_go.git",
    tag = "0.4.0",
)

load("@io_bazel_rules_go//go:def.bzl", "go_repositories")
go_repositories()

new_http_archive(
    name = "docker_ubuntu",
    url = "https://codeload.github.com/tianon/docker-brew-ubuntu-core/zip/b6f1fe19228e5b6b7aed98dcba02f18088282f90",
    build_file = "ubuntu.BUILD",
    type = "zip",
    sha256 = "2c63dd81d714b825acd1cb3629c57d6ee733645479d0fcdf645203c2c35924c5",
)

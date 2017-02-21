# http://bazel.build/
# vim: set ft=python sts=2 sw=2 et:

load("@io_bazel_rules_go//go:def.bzl", "go_binary")
load("@bazel_tools//tools/build_defs/docker:docker.bzl", "docker_build")

go_binary(
    name = "hello_server",
    srcs = ["main.go"],
    visibility = ["//visibility:public"],
)

docker_build(
    name = "hello",
    base = "@docker_ubuntu//:xenial",
    files = [":hello_server"],
    ports = ["80"],
    cmd = "/hello_server -listen=:80",
    repository = "us.gcr.io/playground-159317",
)

# http://bazel.build/
# vim: set ft=python sts=2 sw=2 et:

load("@bazel_tools//tools/build_defs/docker:docker.bzl", "docker_build")

genrule(
    name = "xenial_tar",
    srcs = ["docker-brew-ubuntu-core-b6f1fe19228e5b6b7aed98dcba02f18088282f90/xenial/ubuntu-xenial-core-cloudimg-amd64-root.tar.gz"],
    outs = ["xenial_tar.tar"],
    cmd = "zcat <$< >$@",
)

docker_build(
    name = "xenial",
    tars = [":xenial_tar"],
    visibility = ["//visibility:public"],
)

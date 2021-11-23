load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "section-1-lib",
    srcs = ["hello-world.go"],
    importpath = "chapter-1/section-1",
)

go_binary(
    name = "section-1",
    embed = [":section-1-lib"],
)

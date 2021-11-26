load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "fetch-lib",
    srcs = ["fetch.go"],
    importpath = "chapter-1/section-5",
)

# bazel run //chapter-1/section-5:fetch -- https://www.google.ca/
go_binary(
    name = "fetch",
    embed = [":fetch-lib"],
)

# TODO: Exercise 1.7, 1.8, 1.9


load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "echo1-lib",
    srcs = ["echo1.go"],
    importpath = "chapter-1/section-2",
)

# Contains changes for exercise 1
# bazel run //chapter-1/section-2:echo1 -- testing 1 2 3
go_binary(
    name = "echo1",
    embed = [":echo1-lib"],
)


go_library(
    name = "echo2-lib",
    srcs = ["echo2.go"],
    importpath = "chapter-1/section-2",
)

# Contains changes for exercise 2
# bazel run //chapter-1/section-2:echo2 -- testing 1 2 3
go_binary(
    name = "echo2",
    embed = [":echo2-lib"],
)


go_library(
    name = "echo3-lib",
    srcs = ["echo3.go"],
    importpath = "chapter-1/section-2",
)

# bazel run //chapter-1/section-2:echo3 -- testing 1 2 3
go_binary(
    name = "echo3",
    embed = [":echo3-lib"],
)

# TODO - Exercise 3


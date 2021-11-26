load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "dup1-lib",
    srcs = ["dup1.go"],
    importpath = "chapter-1/section-3",
)

# bazel run //chapter-1/section-3:dup1
go_binary(
    name = "dup1",
    embed = [":dup1-lib"],
)

go_library(
    name = "dup2-lib",
    srcs = ["dup2.go"],
    importpath = "chapter-1/section-3",
)

# bazel run //chapter-1/section-3:dup2
go_binary(
    name = "dup2",
    embed = [":dup2-lib"],
)


go_library(
    name = "dup3-lib",
    srcs = ["dup3.go"],
    importpath = "chapter-1/section-3",
)

# bazel run //chapter-1/section-3:dup3
go_binary(
    name = "dup3",
    embed = [":dup3-lib"],
)


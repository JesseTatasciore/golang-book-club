load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "boiling-lib",
    srcs = ["boiling.go"],
    importpath = "chapter-2/section-2",
)

# bazel run //chapter-2/section-2:boiling
go_binary(
    name = "boiling",
    embed = [":boiling-lib"],
)

go_library(
    name = "ftoc-lib",
    srcs = ["ftoc.go"],
    importpath = "chapter-2/section-2",
)

# bazel run //chapter-2/section-2:ftoc
go_binary(
    name = "ftoc",
    embed = [":ftoc-lib"],
)

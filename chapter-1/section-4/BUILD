load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "lissajous-lib",
    srcs = ["lissajous.go"],
    importpath = "chapter-1/section-4",
)

# bazel run //chapter-1/section-4:lissajous > result.gif
go_binary(
    name = "lissajous",
    embed = [":lissajous-lib"],
)

# TODO: Exercise 1.5 and 1.6
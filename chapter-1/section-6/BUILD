load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "fetchall-lib",
    srcs = ["fetchall.go"],
    importpath = "chapter-1/section-6",
)

# bazel run //chapter-1/section-6:fetchall -- https://www.google.ca/ https://www.google.ca/
go_binary(
    name = "fetchall",
    embed = [":fetchall-lib"],
)

# TODO: Exercise 1.7, 1.8, 1.9


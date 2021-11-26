load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "server1-lib",
    srcs = ["server1.go"],
    importpath = "chapter-1/section-7",
)

# bazel run //chapter-1/section-7:server1
go_binary(
    name = "server1",
    embed = [":server1-lib"],
)

go_library(
    name = "server2-lib",
    srcs = ["server2.go"],
    importpath = "chapter-1/section-7",
)

# bazel run //chapter-1/section-7:server2
go_binary(
    name = "server2",
    embed = [":server2-lib"],
)

go_library(
    name = "server3-lib",
    srcs = ["server3.go"],
    importpath = "chapter-1/section-7",
)

# bazel run //chapter-1/section-7:server3
go_binary(
    name = "server3",
    embed = [":server3-lib"],
)

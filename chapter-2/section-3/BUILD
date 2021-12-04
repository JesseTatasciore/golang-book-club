load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "echo4-lib",
    srcs = ["echo4.go"],
    importpath = "chapter-2/section-3",
)

# bazel run //chapter-2/section-3:echo4 -- a bc def
# bazel run //chapter-2/section-3:echo4 -- -s / a bc def
# bazel run //chapter-2/section-3:echo4 -- -n a bc def
# bazel run //chapter-2/section-3:echo4 -- -help
go_binary(
    name = "echo4",
    embed = [":echo4-lib"],
)

go_library(
    name = "tempconv-lib",
    srcs = [
        "tempconv.go",
        "conv.go",
    ],
    importpath = "chapter-2/section-3/tempconv",
)

go_library(
    name = "cf-lib",
    srcs = ["cf.go"],
    deps = [":tempconv-lib"],
    importpath = "chapter-2/section-3",
)

# bazel run //chapter-2/section-3:cf -- 100
go_binary(
    name = "cf",
    embed = [":cf-lib"],
)

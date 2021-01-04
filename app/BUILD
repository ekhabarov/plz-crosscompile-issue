go_binary(
    name = "bin",
    srcs = [ "main.go" ],
    deps = [ "//app/prober:prober" ],
)

subinclude("///pleasings//docker", "///pleasings//k8s")

docker_image(
    name = "image",
    srcs = [":bin"],
    base_image = "//common/docker:golang",
    dockerfile = "//k8s:dockerfile",
    image = "app_image",
    version = "image",
)

k8s_config(
    name = "yaml",
    srcs = [
        "//k8s:deployment",
        "//k8s:service",
    ],
    containers = [":image"],
    params = {
        "API_NAME": "app",
    },
)

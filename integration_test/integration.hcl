param "image-name" {
  default = "centos"
}

param "image-tag" {
  default = "latest"
}

param "local-converge-directory" {
  default = "{{ env `PWD`}}/build/converge_0.1.1_linux_amd64"
}

docker.image "integration-test" {
  name    = "{{ param `image-name` }}"
  tag     = "{{ param `image-tag` }}"
  timeout = "60s"
}

docker.container "integration-test" {
  name  = "integration-test-{{ param `image-name` }}"
  image = "{{param `image-name`}}:{{param `image-tag`}}"

  expose = [
    "2693",
    "2694"
  ]
  publish_all_ports = "false"
  ports = [
    "26930:2693",
    "29640:2694",
  ]

  volumes     = ["{{param `local-converge-directory`}}:/converge"]
  working_dir = "/converge"

  command = ["./converge", "server", "--no-token"]

  depends = ["docker.image.integration-test"]
}

task "basic converge on container" {
  interpreter = "/bin/bash"
  check_flags = ["-n"]

  check = "./converge plan --rpc-addr :26930 samples/basic.hcl"
  apply = "./converge apply --rpc-addr :26930 samples/basic.hcl"

  depends = ["docker.container.integration-test"]
}

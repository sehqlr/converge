param "image-name" {}

param "image-tag" {
  default = "latest"
}

param "local-samples-directory" {
  default = "{{ env `PWD`}}/samples"
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
    "2693:2693",
    "2694:2694",
  ]

  volumes     = [
    "{{param `local-converge-directory`}}:/converge",
    "{{param `local-samples-directory`}}:/samples",
  ]
  working_dir = "/converge"

  command = ["./converge", "server", "--no-token"]

  depends = ["docker.image.integration-test"]
}

task.query "sleep" {
  query = "sleep 10"
}

task "basic converge on container" {
  interpreter = "/bin/bash"
  check_flags = ["-n"]

  check = "./converge plan --rpc-addr=localhost:2693 /samples/basic.hcl"
  apply = "./converge apply --rpc-addr=localhost:2693 /samples/basic.hcl"

  depends = ["task.query.sleep", "docker.container.integration-test"]
}

param "image-name" {
  default = "centos"
}

param "converge-directory" {
  default = "{{ env `PWD`}}"
}

param "converge-bin" {
  default = "{{ param `converge-directory` }}/build/converge_0.1.1_linux_amd64"
}

docker.image "integration-test" {
  name    = "{{ param `image-name` }}"
  tag     = "latest"
  timeout = "60s"
}

docker.container "integration-test" {
  name  = "integration-test-{{ param `image-name` }}"
  image = "{{param `image-name`}}:latest"

  expose = [
    "2693", 
    "2694"
  ]
  publish_all_ports = "false"
  ports = [
    "2693:26930",
    "2964:26940",
  ]

  volumes     = ["{{param `converge-directory`}}:/converge"]
  working_dir = "/converge"

  command = [
    "cp {{param `converge-bin`}} converge-{{ param `image-name` }}",
    "./converge-{{param `image-name`}} server -no-token",
  ]

  depends = ["docker.image.integration-test"]
}

task "basic converge on container" {
  interpreter = "/bin/bash"
  check_flags = ["-n"]

  check = <<END
./converge --rpc-addr :26930 plan samples/basic.hcl
END

  depends = ["docker.container.integration-test"]
}

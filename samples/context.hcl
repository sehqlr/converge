context "con" {
  workDir = "/tmp"
}

task "touch" {
  check = "test -f converge-test"
  apply = "touch converge-test"
  withContext = "con"
}

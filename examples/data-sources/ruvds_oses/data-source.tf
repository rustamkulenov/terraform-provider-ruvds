data "ruvds_oses" "oses" {
  with_type = "linux"
}

output "linux_oses" {
  value = data.ruvds_oses.oses
}
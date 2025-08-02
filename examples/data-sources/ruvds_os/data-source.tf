data "ruvds_os" "ubuntu_2204" {
  with_code = "53-ubuntu-22.04-lts-eng"
}

output "os_ubuntu_2204" {
  value = data.ruvds_os.ubuntu_2204
}

# Gets information about a specific OS.
# Output:
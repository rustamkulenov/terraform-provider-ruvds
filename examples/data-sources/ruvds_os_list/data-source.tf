data "ruvds_os_list" "oses" {
  with_type = "linux"
}

output "linux_oses" {
  value = data.ruvds_os_list.oses
}

# Get list of all Linux OSes.
# Output:
#  + linux_oses      = {
#      + codes     = [
#          + "8-debian-8-eng",
#          + "10-centos-7-eng",
#          + "12-ubuntu-16.04-lts-eng",
#          + "18-ubuntu-18.04-lts-eng",
#          + "19-centos-7.6.1810-eng",
#          + "36-debian-10-eng",
#          + "42-centos-8-eng",
#          + "44-ubuntu-20.04-lts-eng",
#          + "45-debian-11-eng",
#          + "46-centos-stream-9-eng",
#          + "52-debian-12-eng",
#          + "53-ubuntu-22.04-lts-eng",
#        ]
#      + with_type = "linux"
#    }

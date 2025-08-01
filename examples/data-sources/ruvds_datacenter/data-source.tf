data "ruvds_datacenter" "zur1" {
  with_code = "ZUR1"
}

output "datacenter_zur1" {
  value = data.ruvds_datacenter.zur1
}

# Gets information about a specific datacenter in Switzerland.
# Output:
#  + datacenter_zur1 = {
#      + code      = "ZUR1"
#      + country   = "CH"
#      + id        = 2
#      + name      = "ZUR1: Швейцария, Цюрих"
#      + with_code = "ZUR1"
#    }
resource "ruvds_vps" "my_vps" {
  datacenter_id   = 1
  tariff_id       = 14
  payment_period  = 2
  os_id           = 255
  cpu             = 1
  ram             = 1.0
  drive           = 20
  drive_tariff_id = 1
  ip              = 1
}

output "my_vps_output" {
  value = resource.ruvds_vps.my_vps
}

# Creates a new virtual server with the specified hardcoded parameters (mandatory).
# These parameters can also be replaced with variables or data sources as needed. 
#
# > tofu plan
#
# Output:
#  # ruvds_vps.my_vps will be created
#  + resource "ruvds_vps" "my_vps" {
#      + cpu             = 1
#      + create_progress = (known after apply)
#      + datacenter_id   = 1
#      + drive           = 20
#      + drive_tariff_id = 1
#      + id              = (known after apply)
#      + ip              = 1
#      + os_id           = 255
#      + paid_till       = (known after apply)
#      + payment_period  = 2
#      + ram             = 1
#      + status          = (known after apply)
#      + tariff_id       = 14
#    }
#
#Plan: 1 to add, 0 to change, 0 to destroy.
#
#  + my_vps_output      = {
#      + additional_drive           = null
#      + additional_drive_tariff_id = null
#      + computer_name              = null
#      + cpu                        = 1
#      + create_progress            = (known after apply)
#      + datacenter_id              = 1
#      + ddos_protection            = null
#      + drive                      = 20
#      + drive_tariff_id            = 1
#      + id                         = (known after apply)
#      + ip                         = 1
#      + os_id                      = 255
#      + paid_till                  = (known after apply)
#      + payment_period             = 2
#      + ram                        = 1
#      + ssh_key_id                 = null
#      + status                     = (known after apply)
#      + tariff_id                  = 14
#      + template_id                = null
#      + user_comment               = null
#      + vram                       = null
#    }

# > tofu apply
#
#ruvds_vps.my_vps: Creating...
#ruvds_vps.my_vps: Creation complete after 1s
#
#Apply complete! Resources: 1 added, 0 changed, 1 destroyed.
#
#Outputs:
#
#my_vps_output = {
#  "additional_drive" = tonumber(null)
#  "additional_drive_tariff_id" = tonumber(null)
#  "computer_name" = tostring(null)
#  "cpu" = 1
#  "create_progress" = 0
#  "datacenter_id" = 1
#  "ddos_protection" = tonumber(null)
#  "drive" = 20
#  "drive_tariff_id" = 1
#  "id" = 2225553
#  "ip" = 1
#  "os_id" = 255
#  "paid_till" = ""
#  "payment_period" = 2
#  "ram" = 1
#  "ssh_key_id" = tostring(null)
#  "status" = "initializing"
#  "tariff_id" = 14
#  "template_id" = tostring(null)
#  "user_comment" = tostring(null)
#  "vram" = tonumber(null)
#}

# If then remove the server from this configuration and run `tofu plan` again, you will see that the server will be marked for destruction:
#  # ruvds_vps.my_vps will be destroyed
#  # (because ruvds_vps.my_vps is not in configuration)
#  - resource "ruvds_vps" "my_vps" {
#      - cpu             = 1 -> null
#      - create_progress = 0 -> null
#      - datacenter_id   = 1 -> null
#      - drive           = 20 -> null
#      - drive_tariff_id = 1 -> null
#      - id              = 2225631 -> null
#      - ip              = 1 -> null
#      - os_id           = 255 -> null
#      - payment_period  = 2 -> null
#      - ram             = 1 -> null
#      - status          = "initializing" -> null
#      - tariff_id       = 14 -> null
#    }
#
#Plan: 0 to add, 0 to change, 1 to destroy.

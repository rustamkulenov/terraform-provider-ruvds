data "ruvds_vps_list" "my_vps_list" {
}

output "my_vps_list_output" {
  value = data.ruvds_vps_list.my_vps_list
}

# Get list of all VPS instances.
# Output:
# + my_vps_list_output = {
#      + servers = [
#          + {
#              + additional_drive           = null
#              + additional_drive_tariff_id = null
#              + cpu                        = 15
#              + create_progress            = 100
#              + datacenter_id              = 1
#              + ddos_protection            = 0
#              + drive                      = 20
#              + drive_tariff_id            = 1
#              + id                         = 1234567
#              + ip                         = 1
#              + os_id                      = 255
#              + paid_till                  = ""
#              + payment_period             = 2
#              + ram                        = 1
#              + status                     = "active"
#              + tariff_id                  = 14
#              + template_id                = null
#              + user_comment               = "small VDI with Ubuntu"
#              + vram                       = 0
#            },
#          + {
#              + additional_drive           = null
#              + additional_drive_tariff_id = null
#              + cpu                        = 15
#              + create_progress            = 100
#              + datacenter_id              = 1
#              + ddos_protection            = 0
#              + drive                      = 20
#              + drive_tariff_id            = 1
#              + id                         = 1234568
#              + ip                         = 1
#              + os_id                      = 255
#              + paid_till                  = ""
#              + payment_period             = 2
#              + ram                        = 1
#              + status                     = "active"
#              + tariff_id                  = 14
#              + template_id                = null
#              + user_comment               = "another small VDI with Ubuntu"
#              + vram                       = 0
#            },
#        ]
#    }

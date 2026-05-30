data "ruvds_tariffs" "test" {
  only_active = true
}

output "tariffs" {
  value = data.ruvds_tariffs.test
}

# Get list of all active tariffs.
# Output:
#   + tariffs = {
#       + additional_drive         = [
#           + {
#               + id        = 4
#               + is_active = true
#               + name      = "Huge HDD"
#               + price     = 0.634765625
#             },
#         ]
#       + additional_service       = [
#           + {
#               + id        = 1
#               + is_active = true
#               + name      = "Remote Desktop Services Client Access Licenses"
#               + price     = 1081
#             },
#           + {
#               + id        = 2
#               + is_active = true
#               + name      = "Microsoft SQL Server Standard MVL SAL "
#               + price     = 2642
#             },
#           + {
#               + id        = 4
#               + is_active = true
#               + name      = "DDoS Protection (IP)"
#               + price     = 400
#             },
#           + {
#               + id        = 6
#               + is_active = true
#               + name      = "Гб (бэкап образа диска сервера)"
#               + price     = 6.5
#             },
#           + {
#               + id        = 7
#               + is_active = true
#               + name      = "Microsoft Office Standard MVL SAL"
#               + price     = 2099
#             },
#           + {
#               + id        = 8
#               + is_active = true
#               + name      = <<-EOT
#                     Microsoft Office Professional Plus
#                      MVL SAL
#                 EOT
#               + price     = 2865
#             },
#           + {
#               + id        = 10
#               + is_active = true
#               + name      = "Microsoft SQL Server Standard 2Lic Core"
#               + price     = 24462
#             },
#           + {
#               + id        = 11
#               + is_active = true
#               + name      = "Microsoft SQL Server Web 2Lic Core"
#               + price     = 1389
#             },
#           + {
#               + id        = 12
#               + is_active = true
#               + name      = "ISPmanager 6 Lite"
#               + price     = 200
#             },
#           + {
#               + id        = 17
#               + is_active = true
#               + name      = "Антивирус Касперского"
#               + price     = 800
#             },
#           + {
#               + id        = 19
#               + is_active = true
#               + name      = "Plesk for VPS Web Admin Edition"
#               + price     = 0
#             },
#           + {
#               + id        = 23
#               + is_active = true
#               + name      = "Снапшот"
#               + price     = 0.800000011920929
#             },
#           + {
#               + id        = 27
#               + is_active = true
#               + name      = "1C License"
#               + price     = 1595
#             },
#         ]
#       + drive                    = [
#           + {
#               + id        = 1
#               + is_active = true
#               + name      = "HDD"
#               + price     = 9
#             },
#           + {
#               + id        = 3
#               + is_active = true
#               + name      = "SSD"
#               + price     = 18.5
#             },
#           + {
#               + id        = 7
#               + is_active = true
#               + name      = "NVMe"
#               + price     = 21.5
#             },
#           + {
#               + id        = 9
#               + is_active = true
#               + name      = "HDDEurope"
#               + price     = 10.5
#             },
#           + {
#               + id        = 10
#               + is_active = true
#               + name      = "SSDEurope"
#               + price     = 21.5
#             },
#         ]
#       + only_active              = true
#       + payment_period_discounts = [
#           + {
#               + discount       = 0.05000000074505806
#               + payment_period = 3
#             },
#           + {
#               + discount       = 0.10000000149011612
#               + payment_period = 4
#             },
#           + {
#               + discount       = 0.20000000298023224
#               + payment_period = 5
#             },
#         ]
#       + vps                      = [
#           + {
#               + cpu       = 135
#               + id        = 14
#               + ip        = 180
#               + is_active = true
#               + name      = "Regular"
#               + ram       = 324
#               + vram      = 1.953125
#             },
#           + {
#               + cpu       = 311
#               + id        = 15
#               + ip        = 180
#               + is_active = true
#               + name      = "Premium"
#               + ram       = 324
#               + vram      = 1.953125
#             },
#           + {
#               + cpu       = 130
#               + id        = 21
#               + ip        = 180
#               + is_active = true
#               + name      = "HugeServer"
#               + ram       = 307
#               + vram      = 1.953125
#             },
#           + {
#               + cpu       = 106
#               + id        = 22
#               + ip        = 180
#               + is_active = true
#               + name      = "Promo32"
#               + ram       = 252
#               + vram      = 1.953125
#             },
#           + {
#               + cpu       = 375
#               + id        = 26
#               + ip        = 180
#               + is_active = true
#               + name      = "Powerful"
#               + ram       = 375
#               + vram      = 1.953125
#             },
#           + {
#               + cpu       = 154
#               + id        = 40
#               + ip        = 180
#               + is_active = true
#               + name      = "RegularEurope"
#               + ram       = 367
#               + vram      = 1.953125
#             },
#           + {
#               + cpu       = 356
#               + id        = 41
#               + ip        = 180
#               + is_active = true
#               + name      = "PremiumEurope"
#               + ram       = 367
#               + vram      = 1.953125
#             },
#         ]
#     }
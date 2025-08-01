data "ruvds_datacenters" "dcs" {
  in_country = "RU"
}

output "dcs_in_ru" {
  value = data.ruvds_datacenters.dcs
}

# Gets list of datacenters in Russia.
# Output:
#  + dcs_in_ru       = {
#      + codes      = [
#          + "BUNKER",
#          + "M9",
#          + "LINXDATACENTER",
#          + "ITPARK",
#          + "EKB",
#          + "SIBTELCO",
#          + "OSTANKINO",
#          + "PORTTELEKOM",
#          + "TELEMAKS",
#          + "SMARTKOM",
#          + "ARKTICHESKIJ COD",
#        ]
#      + in_country = "RU"
#    }

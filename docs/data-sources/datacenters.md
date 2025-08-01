---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "ruvds_datacenters Data Source - ruvds"
subcategory: ""
description: |-
  Datacenters data source
---

# ruvds_datacenters (Data Source)

Datacenters data source

## Example Usage

```terraform
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
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `in_country` (String) Country code to filter datacenters

### Read-Only

- `codes` (List of String) Data Center codes

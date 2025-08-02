# Terraform Provider for [RUVDS](https://ruvds.com)

_This repository is built on the [Terraform Plugin Framework](https://github.com/hashicorp/terraform-plugin-framework)._

[RuVDS](https://ruvds.com) is an IAAS (cloud and VDS/VPS) provider offering virtual servers, dedicated hosting, domain registration, and DDoS protection at competitive prices. It features data centers in 7 countries, supporting both Linux and Windows environments. 

This repository is a [Terraform](https://www.terraform.io)/[OpenTofu](https://opentofu.org/) provider allowing dynamic resource description, planning and provisioning on RuVDP and containing:

- Resources and data sources (`internal/provider/`),
- Examples (`examples/`) and generated documentation (`docs/`),
- Github actions pipeline (`.github/`) for building, testing, releasing,
- Tools for generating documentation, processing files, formatting *.tf files, etc (`tools/`)
- Miscellaneous meta files.

## Requirements

- [Terraform](https://developer.hashicorp.com/terraform/downloads) >= 1.0
or 
[OpenTofu](http://https://opentofu.org/)
- [Go](https://golang.org/doc/install) >= 1.23

## Building The Provider

1. Clone the repository
1. Enter the repository directory
1. Build the provider using the Go `install` or `buld` command:

```shell
make install
```

## Functionality of the provider

This [OpenTofu provider](https://opentofu.org/docs/language/providers/) contains 
[data sources](https://opentofu.org/docs/language/data-sources/), 
[resources](https://opentofu.org/docs/language/resources/),
[functions](https://opentofu.org/docs/language/functions/#provider-defined-functions) as described in the table below:

| Entity | Type | Name | Description
| --- | ---  | ---  | --- |
| Data center | Data Source | [ruvds_datacenters](docs/data-sources/datacenters.md) | Gets list of data centers filtered by country
| Data center | Data Source | [ruvds_datacenter](docs/data-sources/datacenter.md) | Gets information about specific Data center by code
| OS | - | - | - |
| Templates | - | - | - |
| Servers | - | - | - |
| SSH Keys | - | - | - |
| Tarifs | - | - | - |
| Balance | - | - | - |
| Payments | - | - | - |

## Using the provider

For using RuVDS provider you need to obtain API v2 token first. After that you can use it in your Terraform/OpenTofu configurations for getting data and provision resources.

Sample usage of the provider:
```yaml
terraform {
  required_providers {
    ruvds = {
      source  = "hashicorp/ruvds"
      # Choose required version
      #version = "1.0.0"
    }
  }
}

provider "ruvds" {
    # Provide your own API v2 token
    token = "apiv2.*****YOUR-TOKEN*****"
}

# Get a data center by its code
data "ruvds_datacenter" "zur1" {
  with_code = "ZUR1"
}

output "datacenter_zur1" {
  value = data.ruvds_datacenter.zur1
}

# Get list of datacenters in Russia
data "ruvds_datacenters" "dcs" {
    in_country = "RU"
}

output "dcs_in_ru" {
    value = data.ruvds_datacenters.dcs
}
```

After 
```shell
tofu plan
```

you'll see output like this:

```text
data.ruvds_datacenter.zur1: Reading...
data.ruvds_datacenters.dcs: Reading...
data.ruvds_datacenters.dcs: Read complete after 1s
data.ruvds_datacenter.zur1: Read complete after 1s [name=ZUR1: Швейцария, Цюрих]

Changes to Outputs:
  + datacenter_zur1 = {
      + code      = "ZUR1"
      + country   = "CH"
      + id        = 2
      + name      = "ZUR1: Швейцария, Цюрих"
      + with_code = "ZUR1"
    }
  + dcs_in_ru       = {
      + codes      = [
          + "BUNKER",
          + "M9",
          + "LINXDATACENTER",
          + "ITPARK",
          + "EKB",
          + "SIBTELCO",
          + "OSTANKINO",
          + "PORTTELEKOM",
          + "TELEMAKS",
          + "SMARTKOM",
          + "ARKTICHESKIJ COD",
        ]
      + in_country = "RU"
    }
```
## Developing the Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (see [Requirements](#requirements) above).

The provider is built on the [Terraform Plugin Framework](https://github.com/hashicorp/terraform-plugin-framework) and uses [RuVDS API v2.24](https://ruvds.com/api-docs/).

To compile the provider, run `go install`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

`go build` will build the provider and put into current folder.

To generate or update documentation, run `make generate`.

### Testing

The repository contains unit and integration tests.

_Unit tests_ do not consume real resources and can be run even without internet connection at any time.
```shell
make test
```

In order to run the full suite of _Acceptance tests_, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run.

```shell
make testacc
```

### Environment variables

The following environment variables are used in building pipeline and code:

| Variable | Purpose |
| -- | -- |
| TF_ACC | If set to '1' then acceptance tests will be run by github actions or localy (make testacc) |
| RUVDS_API_TOKEN  | API v2 Token for RuVDS. Shall be set if you want to run aceptance tests (or unit tests for testing API) localy |

How to set environment variables in VS Code (in `.vscode/settings.json`):

```json
{
    "go.testEnvVars": {        
        "TF_ACC": "0",
        "RUVDS_API_TOKEN": "apiv2.*****YOUR-TOKEN*******"
    }
}
```

### Local provider override

If you need to use/debug local provider in opentofu scripts then you'll probably need to configure tofu like this. Create `~/.tofurc` file with similar content (replace with your path to compiled provider folder):

```text
provider_installation {

  # Use provided path as an overridden package directory
  # for the hashicorp/ruvds provider. This disables the version and checksum
  # verifications for this provider and forces OpenTofu to look for the
  # provider plugin in the given directory.
  dev_overrides {
    "hashicorp/ruvds" = "/Users/rustam/projects/devops/terraform-provider-ruvds"
  }

  # For all other providers, install them directly from their origin provider
  # registries as normal. If you omit this, OpenTofu will _only_ use
  # the dev_overrides block, and so no other providers will be available.
  direct {}
}
```

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccOSDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: testAccOSDataSourceConfig,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"data.ruvds_os.ubuntu_2204",
						tfjsonpath.New("id"),
						knownvalue.Int32Exact(53),
					),
					statecheck.ExpectKnownValue(
						"data.ruvds_os.ubuntu_2204",
						tfjsonpath.New("ssh_keys_supported"),
						knownvalue.Bool(true),
					),
				},
			},
		},
	})
}

const testAccOSDataSourceConfig = `
data "ruvds_os" "ubuntu_2204" {
  with_code = "53-ubuntu-22.04-lts-eng"
}

output "os_ubuntu_2204" {
  value = data.ruvds_os.ubuntu_2204
}
`

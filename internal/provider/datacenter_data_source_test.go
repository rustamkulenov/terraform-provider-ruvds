package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccDatacenterDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: testAccDatacenterDataSourceConfig,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"data.ruvds_datacenter.test",
						tfjsonpath.New("id"),
						knownvalue.Int32Exact(3),
					),
					statecheck.ExpectKnownValue(
						"data.ruvds_datacenter.test",
						tfjsonpath.New("country"),
						knownvalue.StringExact("GB"),
					),
				},
			},
		},
	})
}

const testAccDatacenterDataSourceConfig = `
 data "ruvds_datacenter" "test" {
  with_code = "LD8"
}
`

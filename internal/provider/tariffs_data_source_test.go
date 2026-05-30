package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccTariffsDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccTariffsDataSourceConfig,
				Check: resource.ComposeTestCheckFunc(
					// Check number of responses (this number may change if RUVDS adds or removes tariffs, so we just check that there are some)
					resource.TestCheckResourceAttr("data.ruvds_tariffs.test", "vps.#", "7"),
					// Check that only active are requested
					resource.TestCheckResourceAttr("data.ruvds_tariffs.test", "only_active", "true"),
					// Check value of the 1st response
					resource.TestCheckResourceAttr("data.ruvds_tariffs.test", "vps.0.is_active", "true"),
				),
			},
		},
	})
}

const testAccTariffsDataSourceConfig = `
data "ruvds_tariffs" "test" {
  only_active = true
}

output "tariffs" {
  value = data.ruvds_tariffs.test
}
`

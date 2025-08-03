package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
)

func TestAccVpsListDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config:            testAccVpsListDataSourceConfig,
				ConfigStateChecks: []statecheck.StateCheck{},
			},
		},
	})
}

const testAccVpsListDataSourceConfig = `
data "ruvds_vps_list" "my_vps_list" {  
}

output "my_vps_list_output" {
  value = data.ruvds_vps_list.my_vps_list
}
`

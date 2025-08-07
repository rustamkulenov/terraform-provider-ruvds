package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestVpsResourceImport(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// First create a test resource
			{
				Config: testAccVpsResourceConfig(),
			},
			// Then test importing it
			{
				ResourceName:      "ruvds_vps.my_srv",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccVpsResourceConfig() string {
	return `

`
}

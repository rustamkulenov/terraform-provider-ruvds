package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestSshResourceImport(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// First create a test resource
			{
				Config: testAccSshResourceConfig(),
			},
			// Then test importing it
			{
				ResourceName:      "ruvds_ssh.key",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccSshResourceConfig() string {
	return `
resource "ruvds_ssh" "key" {  
	name="MacBook"
	public_key="ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDjW5QH7QZqJtR7X8wXJkS7Q6Xy7Yt8zRtNjYJ7a1Xv7s3w5ZK1f8z2wY5Xk"
}
`
}

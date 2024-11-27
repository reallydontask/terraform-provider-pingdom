package pingdom


import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

//Make this meaningful
func TestAccDataSourcePingdomProbe_basic(t *testing.T) {
	datasourceName := "data.pingdom_probe.test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourcePingdomProbeConfig(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckPingdomResourceID(datasourceName),
				),
			},
		},
	})
}

func testAccDataSourcePingdomProbeConfig() string {
	return `
data "pingdom_probe" "test" {
}

output "parsed_probes" {
  value = join(",",[ for v in data.pingdom_probe.all.probes : v.ip ])
}
`
}
package pingdom

import (
	"context"
	
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourcePingdomProbe() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourcePingdomProbeRead,

		Schema: map[string]*schema.Schema{
			"probes": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"active": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"country": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"countryiso": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"city": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"hostname": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ip": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ipv6": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"region": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourcePingdomProbeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*Clients).Pingdom
	probes, err := client.Probes.List()
	if err != nil {
		return diag.FromErr(err)
	}

	var diags diag.Diagnostics

	// Transform the probes into the expected schema format
	probesData := make([]map[string]interface{}, len(probes))
	for i, probe := range probes {
		probesData[i] = map[string]interface{}{
			"active":     probe.Active,
			"country":    probe.Country,
			"countryiso": probe.CountryISO,
			"city":       probe.City,
			"hostname":   probe.Hostname,
			"ip":         probe.IP,
			"ipv6":       probe.IPv6,
			"name":       probe.Name,
			"region":     probe.Region,
		}
	}

	// Set the "probes" field with the data
	if err := d.Set("probes", probesData); err != nil {
		return diag.FromErr(err)
	}

	// Set a dummy ID since data sources require an ID
	d.SetId("all_probes")

	return diags
}
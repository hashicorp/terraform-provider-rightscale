package rs

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/rightscale/terraform-provider-rs/rs/rsc"
)

// Example:
//
// data "rs_cm_datacenter" "ec2-us-east-1a" {
//     cloud = ${data.rs_cm_cloud.ec2_us_east.id}
//     filter {
//         name = "us-east-1a"
//     }
// }

func dataSourceDatacenters() *schema.Resource {
	return &schema.Resource{
		Read: resourceDatacenterRead,

		Schema: map[string]*schema.Schema{
			"cloud": {
				Type:        schema.TypeString,
				Description: "href to datacenter cloud",
				Required:    true,
				ForceNew:    true,
			},
			"filter": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"resource_uid": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
					},
				},
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"resource_uid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceDatacenterRead(d *schema.ResourceData, m interface{}) error {
	client := m.(rsc.Client)
	cloud := d.Get("cloud").(string)
	loc := &rsc.Locator{Namespace: "rs_cm", Href: cloud}

	res, err := client.List(loc, "datacenters", cmFilters(d))
	if err != nil {
		return err
	}

	if len(res) == 0 {
		return nil
	}
	for k, v := range res[0].Fields {
		d.Set(k, v)
	}
	d.SetId(res[0].Locator.Href)
	return nil
}

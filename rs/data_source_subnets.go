package rs

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/rightscale/terraform-provider-rs/rs/rsc"
)

// Example:
//
// data "rs_cm_subnet" "ssh" {
//     filter {
//         name = "infra"
//     }
//     cloud = ${data.rs_cm_cloud.ec2_us_east_1.id}
// }

func dataSourceSubnets() *schema.Resource {
	return &schema.Resource{
		Read: resourceSubnetRead,

		Schema: map[string]*schema.Schema{
			"cloud": {
				Type:        schema.TypeString,
				Description: "ID of the subnet cloud",
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
							Type:        schema.TypeString,
							Description: "name of subnet, uses partial match",
							Optional:    true,
							ForceNew:    true,
						},
						"resource_uid": {
							Type:        schema.TypeString,
							Description: "cloud ID of subnet",
							Optional:    true,
							ForceNew:    true,
						},
						"datacenter": {
							Type:        schema.TypeString,
							Description: "ID of the subnet datacenter resource",
							Optional:    true,
							ForceNew:    true,
						},
						"instance": {
							Type:        schema.TypeString,
							Description: "ID of instance resource attached to subnet",
							Optional:    true,
							ForceNew:    true,
						},
						"network": {
							Type:        schema.TypeString,
							Description: "ID of network resource that owns subnet",
							Optional:    true,
							ForceNew:    true,
						},
						"visibility": {
							Type:        schema.TypeString,
							Description: "Visibility of the subnet to filter by (private, shared, etc)",
							Optional:    true,
							ForceNew:    true,
						},
					},
				},
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_default": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"cidr_block": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"resource_uid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"visibility": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceSubnetRead(d *schema.ResourceData, m interface{}) error {
	client := m.(rsc.Client)
	cloud := d.Get("cloud").(string)
	loc := &rsc.Locator{Namespace: "rs_cm", Href: cloud}

	res, err := client.List(loc, "subnets", cmFilters(d))
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

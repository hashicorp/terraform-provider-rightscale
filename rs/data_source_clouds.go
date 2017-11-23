package rs

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
	"github.com/rightscale/terraform-provider-rs/rs/rsc"
)

// Example:
//
// data "rs_cm_cloud" "ec2_us_east_1" {
//     filter {
//         name = "EC2 us-east-1"
//         cloud_type = "amazon"
//     }
// }

func dataSourceClouds() *schema.Resource {
	return &schema.Resource{
		Read: resourceCloudRead,

		Schema: map[string]*schema.Schema{
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
						"description": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"cloud_type": {
							Type:         schema.TypeString,
							Optional:     true,
							ForceNew:     true,
							ValidateFunc: validation.StringInSlice(supportedCloudTypes, false),
						},
					},
				},
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"cloud_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"datacenters": {
				Type:     schema.TypeList,
				Elem:     dataSourceDatacenters(),
				Computed: true,
			},
			"instance_types": {
				Type:     schema.TypeList,
				Elem:     dataSourceInstanceTypes(),
				Computed: true,
			},
		},
	}
}

func resourceCloudRead(d *schema.ResourceData, m interface{}) error {
	client := m.(rsc.Client)
	loc := &rsc.Locator{Namespace: "rs_cm", Type: "clouds"}

	res, err := client.List(loc, "", cmFilters(d))
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

var supportedCloudTypes = []string{"aws", "blue_skies", "eucalyptus",
	"rackspace", "cloud_stack", "amazon", "open_stack", "open_stack_grizzly",
	"open_stack_v2", "open_stack_v3", "soft_layer", "google", "azure", "azure_v2",
	"hp", "rackspace_next_gen", "vscale", "uca"}

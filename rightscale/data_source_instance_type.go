package rightscale

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/terraform-providers/terraform-provider-rightscale/rightscale/rsc"
)

// Example:
//
// data "rightscale_instance_type" "n1-standard" {
//   cloud_href = ${data.rightscale_cloud.gce.id}
//   filter {
//     name = "n1-standard"
//   }
// }

func dataSourceInstanceType() *schema.Resource {
	return &schema.Resource{
		Read: resourceInstanceTypeRead,

		Schema: map[string]*schema.Schema{
			"cloud_href": {
				Type:        schema.TypeString,
				Description: "ID of instance cloud resource",
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
							Description: "name of instance type, uses partial match",
							Optional:    true,
							ForceNew:    true,
						},
						"description": {
							Type:        schema.TypeString,
							Description: "description of instance type, uses partial match",
							Optional:    true,
							ForceNew:    true,
						},
						"resource_uid": {
							Type:        schema.TypeString,
							Description: "cloud ID - if this filter is set additional retry logic will fire to allow for cloud resource discovery",
							Optional:    true,
							ForceNew:    true,
						},
						"cpu_architecture": {
							Type:        schema.TypeString,
							Description: "CPU architecture of instance type, e.g. 'x86_64'",
							Optional:    true,
							ForceNew:    true,
						},
					},
				},
			},

			// Read-only fields
			"cpu_architecture": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"cpu_count": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"cpu_speed": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"links": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeMap},
				Computed: true,
			},
			"memory": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"resource_uid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"href": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceInstanceTypeRead(d *schema.ResourceData, m interface{}) error {
	client := m.(rsc.Client)
	cloud := d.Get("cloud_href").(string)
	loc := &rsc.Locator{Namespace: "rs_cm", Href: cloud}

	// if 'resource_uid' filter is set, we expect it to show up.
	// retry for 5 min to allow rightscale to poll cloud to discover.
	if uidset := cmUIDSet(d); uidset {
		timeout := 300
		err := cmIndexRetry(client, loc, "instance_types", d, timeout)
		if err != nil {
			return err
		}
	}

	res, err := client.List(loc, "instance_types", cmFilters(d))
	if err != nil {
		return err
	}

	if len(res) == 0 {
		return nil
	}
	for k, v := range res[0].Fields {
		d.Set(k, v)
	}
	d.Set("href", res[0].Locator.Href)
	d.SetId(res[0].Locator.Namespace + ":" + res[0].Locator.Href)
	return nil
}

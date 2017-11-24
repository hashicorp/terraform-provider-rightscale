package rs

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/rightscale/terraform-provider-rs/rs/rsc"
)

// Example:
//
// data "rs_cm_volume_snapshot" "mysql_master" {
//     filter {
//         name = "mysql_master"
//     }
//     cloud = ${data.rs_cm_cloud.ec2_us_east_1.id}
// }

func dataSourceCMVolumeSnapshot() *schema.Resource {
	return &schema.Resource{
		Read: resourceVolumeSnapshotRead,

		Schema: map[string]*schema.Schema{
			"cloud": {
				Type:        schema.TypeString,
				Description: "ID of the volume snapshot cloud",
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
							Description: "name of volume snapshot, uses partial match",
							Optional:    true,
							ForceNew:    true,
						},
						"description": {
							Type:        schema.TypeString,
							Description: "description of volume snapshot, uses partial match",
							Optional:    true,
							ForceNew:    true,
						},
						"resource_uid": {
							Type:        schema.TypeString,
							Description: "cloud ID of volume snapshot",
							Optional:    true,
							ForceNew:    true,
						},
						"state": {
							Type:        schema.TypeString,
							Description: "state of the volume snapshot",
							Optional:    true,
							ForceNew:    true,
						},
						"deployment": {
							Type:        schema.TypeString,
							Description: "ID of deployment resource that owns volume snapshot",
							Optional:    true,
							ForceNew:    true,
						},
						"parent_volume": {
							Type:        schema.TypeString,
							Description: "ID of volume resource from which volume snapshot was created",
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
			"resource_uid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"size": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"created_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"updated_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceVolumeSnapshotRead(d *schema.ResourceData, m interface{}) error {
	client := m.(rsc.Client)
	cloud := d.Get("cloud").(string)
	loc := &rsc.Locator{Namespace: "rs_cm", Href: cloud}

	res, err := client.List(loc, "volume_snapshots", cmFilters(d))
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

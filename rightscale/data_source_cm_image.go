package rs

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
	"github.com/rightscale/terraform-provider-rightscale/rightscale/rsc"
)

// Example:
//
// data "rightscale_cm_image" "centos_5" {
//     filter {
//         name = "centos 5"
//     }
//     cloud = ${data.rightscale_cm_cloud.gce.id}
// }

func dataSourceCMImage() *schema.Resource {
	return &schema.Resource{
		Read: resourceImageRead,

		Schema: map[string]*schema.Schema{
			"cloud_href": {
				Type:        schema.TypeString,
				Description: "ID of image cloud resource",
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
							Description: "name of image, uses partial match",
							Optional:    true,
							ForceNew:    true,
						},
						"description": {
							Type:        schema.TypeString,
							Description: "description of image, uses partial match",
							Optional:    true,
							ForceNew:    true,
						},
						"resource_uid": {
							Type:        schema.TypeString,
							Description: "cloud id of image",
							Optional:    true,
							ForceNew:    true,
						},
						"os_platform": {
							Type:         schema.TypeString,
							Description:  "The image's operating system to filter on. Examples: Linux or Windows.",
							Optional:     true,
							ForceNew:     true,
							ValidateFunc: validation.StringInSlice([]string{"windows", "linux"}, false),
						},
						"cpu_architecture": {
							Type:        schema.TypeString,
							Description: "CPU architecture of image, e.g. 'x86_64', uses partial match",
							Optional:    true,
							ForceNew:    true,
						},
						"image_type": {
							Type:         schema.TypeString,
							Description:  `The Image Type to filter on. This will be either "machine", "machine_azure", "ramdisk" or "kernel"`,
							Optional:     true,
							ForceNew:     true,
							ValidateFunc: validation.StringInSlice([]string{"machine", "machine_azure", "ramdisk", "kernel"}, false),
						},
						"visibility": {
							Type:         schema.TypeString,
							Description:  "The visibility of the Image in the cloud to filter on. Examples: private, public.",
							Optional:     true,
							ForceNew:     true,
							ValidateFunc: validation.StringInSlice([]string{"machine", "machine_azure", "ramdisk", "kernel"}, false),
						},
					},
				},
			},

			// Read-only fields
			"cpu_architecture": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"image_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"links": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeMap},
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"os_platform": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"resource_uid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"root_device_storage": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"virtualization_type": {
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

func resourceImageRead(d *schema.ResourceData, m interface{}) error {
	client := m.(rsc.Client)
	cloud := d.Get("cloud_href").(string)
	loc := &rsc.Locator{Namespace: "rs_cm", Href: cloud}

	res, err := client.List(loc, "images", cmFilters(d))
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

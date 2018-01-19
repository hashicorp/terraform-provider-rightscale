package rightscale

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/rightscale/terraform-provider-rightscale/rightscale/rsc"
)

// Example:
//
// data "rightscale_ssh_key" "ssh" {
//   filter {
//     name = "infra"
//   }
//   cloud_href = ${data.rightscale_cloud.ec2_us_east_1.id}
//	 view = "sensitive"
// }

func dataSourceSSHKey() *schema.Resource {
	return &schema.Resource{
		Read: datasourceSSHKeyRead,

		Schema: map[string]*schema.Schema{
			"cloud_href": {
				Type:        schema.TypeString,
				Description: "ID of the SSH key cloud",
				Required:    true,
				ForceNew:    true,
			},
			"view": {
				Type:        schema.TypeString,
				Description: "option to return private key material",
				Optional:    true,
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
							Description: "name of SSH key, uses partial match",
							Optional:    true,
							ForceNew:    true,
						},
						"resource_uid": {
							Type:        schema.TypeString,
							Description: "cloud ID of SSH key",
							Optional:    true,
							ForceNew:    true,
						},
					},
				},
			},

			// Read-only fields
			"links": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeMap},
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"material": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"resource_uid": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourceSSHKeyRead(d *schema.ResourceData, m interface{}) error {
	client := m.(rsc.Client)
	cloud := d.Get("cloud_href").(string)
	nsParams := make(map[string]string)
	view := d.Get("view").(string)
	if len(view) == 0 {
		nsParams["view"] = "default"
	} else {
		if view == "sensitive" {
			nsParams["view"] = "sensitive"
		} else {
			return fmt.Errorf("view type is set but invalid: valid options are 'default' or 'sensitive'")
		}
	}

	loc := &rsc.Locator{Namespace: "rs_cm", Href: cloud, NamespaceParams: nsParams}

	res, err := client.List(loc, "ssh_keys", cmFilters(d))
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

package rightscale

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/rightscale/terraform-provider-rightscale/rightscale/rsc"
)

// Example:
//
// resource "rightscale_ssh_key" "ssh" {
//   name = "infra"
//   cloud_href = ${data.rightscale_cloud.ec2_us_east_1.id}
// }

func resourceSSHKey() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSSHKeyRead,
		Exists: resourceExists,
		Delete: resourceDelete,
		Create: resourceCreateFunc("rs_cm", "ssh_keys", sshKeyWriteFields),
		Update: resourceUpdateFunc(sshKeyWriteFields),

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Description: "name of SSH key",
				Type:        schema.TypeString,
				Required:    true,
			},
			"cloud_href": &schema.Schema{
				Description: "The ID of the cloud to operate against",
				Type:        schema.TypeString,
				Required:    true,
			},
			"resource_uid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"material": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"links": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeMap},
				Computed: true,
			},
		},
	}
}

func sshKeyWriteFields(d *schema.ResourceData) rsc.Fields {
	return rsc.Fields{"ssh_key": rsc.Fields{"name": d.Get("name")}, "cloud_href": d.Get("cloud_href")}
}

func resourceSSHKeyRead(d *schema.ResourceData, m interface{}) error {
	client := m.(rsc.Client)
	loc, err := locator(d)
	if err != nil {
		return err
	}

	// set ActionParams Locator to always read this resource with 'view: "sensitive"' for access to private key material
	loc.ActionParams = map[string]string{"view": "sensitive"}

	res, err := client.Get(loc)
	if err != nil {
		return handleRSCError(d, err)
	}
	for k, v := range res.Fields {
		d.Set(k, v)
	}
	return nil
}

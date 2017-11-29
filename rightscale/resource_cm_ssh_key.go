package rs

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/rightscale/terraform-provider-rightscale/rightscale/rsc"
)

func resourceCMSSHKey() *schema.Resource {
	return &schema.Resource{
		Read:   resourceRead,
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
			"resource_uid": {
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
	return rsc.Fields{"name": d.Get("name")}
}

package rightscale

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/terraform-providers/terraform-provider-rightscale/rightscale/rsc"
)

// Example:
//
// resource "rightscale_credential" "database_password" {
//   name = "DATABASE_PASSWORD"
//	 value = "rightscale11"
//	 description = "Top Secret database password"
// }

func resourceCredential() *schema.Resource {
	return &schema.Resource{
		Read:   resourceCredentialRead,
		Exists: resourceExists,
		Delete: resourceDelete,
		Create: resourceCreateFunc("rs_cm", "credentials", credentialWriteFields),
		Update: resourceUpdateFunc(credentialWriteFields),

		Schema: map[string]*schema.Schema{
			"description": {
				Description: "description of the credential object",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"name": {
				Description: "name of the credential object",
				Type:        schema.TypeString,
				Required:    true,
			},
			"value": {
				Description: "value of the credential object",
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
			},
			"updated_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"created_at": {
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

func credentialWriteFields(d *schema.ResourceData) rsc.Fields {
	fields := rsc.Fields{}
	fieldopts := []string{"name", "value", "description"}
	for _, f := range fieldopts {
		if v, ok := d.GetOk(f); ok {
			fields[f] = v
		}
	}
	return rsc.Fields{"credential": fields}
}

func resourceCredentialRead(d *schema.ResourceData, m interface{}) error {
	client := m.(rsc.Client)
	loc, err := locator(d)
	if err != nil {
		return err
	}

	// set ActionParams Locator to always read this resource (currently) with 'view: "sensitive"'
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

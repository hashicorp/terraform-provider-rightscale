package rs

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
	"github.com/rightscale/terraform-provider-rs/rs/rsc"
)

func resourceCMSSHKey() *schema.Resource {
	return &schema.Resource{
		Read:   resourceRead,
		Exists: resourceExists,
		Delete: resourceDelete, // can fail if ssh_key is locked - that's what we want
		Create: resourceCMSSHKeyCreate,
		Update: resourceCMSSHKeyUpdate,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Description: "name of ssh_key",
				Type:        schema.TypeString,
				Required:    true,
			},
			"description": &schema.Schema{
				Description: "description of ssh_key",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"resource_group_href": &schema.Schema{
				Description: "ID of the Windows Azure Resource Group attached to the ssh_key",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"locked": &schema.Schema{
				Description: "whether ssh_key is locked",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"server_tag_scope": &schema.Schema{
				Description:  "routing scope for tags for servers in the ssh_key",
				Type:         schema.TypeString,
				Optional:     true,
				InputDefault: "ssh_key",
				ValidateFunc: validation.StringInSlice([]string{"account", "ssh_key"}, false),
			},
		},
	}
}

func resourceCMSSHKeyCreate(d *schema.ResourceData, m interface{}) error {
	var mustLock bool
	{
		locked, ok := d.GetOk("locked")
		mustLock = ok && locked.(bool)
	}

	client := m.(rsc.Client)
	res, err := client.Create("rs_cm", "ssh_key", ssh_keyFields(d))
	if err != nil {
		return err
	}
	for k, v := range res.Fields {
		d.Set(k, v)
	}

	if mustLock {
		if err := updateLock(d, client); err != nil {
			// Attempt to delete previously created ssh_key, ignore errors
			client.Delete(res.Locator)
			return err
		}
		d.Set("locked", true)
	}

	// set ID last so Terraform does not assume the ssh_key has been
	// created until all operations have completed successfully.
	d.SetId(res.Locator.Namespace + ":" + res.Locator.Href)
	return nil
}

func resourceCMSSHKeyUpdate(d *schema.ResourceData, m interface{}) error {
	d.Partial(true)

	client := m.(rsc.Client)
	loc, err := locator(d)
	if err != nil {
		return err
	}

	// update lock
	if err := updateLock(d, client); err != nil {
		return handleRSCError(d, err)
	}
	d.SetPartial("locked")

	// then the other fields
	if err := client.Update(loc, ssh_keyFields(d)); err != nil {
		return handleRSCError(d, err)
	}

	d.Partial(false)
	return nil
}

// updateLock is a helper function that takes care of locking or unlocking the
// ssh_key according to the value of the "locked" resource data field.
func updateLock(d *schema.ResourceData, client rsc.Client) error {
	loc, err := locator(d)
	if err != nil {
		return err
	}
	lock := d.Get("locked").(bool)
	if lock {
		return client.Run(loc, "@res.lock()")
	}
	return client.Run(loc, "@res.unlock()")
}

func ssh_keyFields(d *schema.ResourceData) rsc.Fields {
	fields := rsc.Fields{"name": d.Get("name")}
	for _, f := range []string{"description", "resource_group_href", "server_tag_scope"} {
		if v, ok := d.GetOk(f); ok {
			fields[f] = v
		}
	}
	return fields
}

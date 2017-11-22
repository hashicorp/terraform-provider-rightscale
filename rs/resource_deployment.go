package rs

import (
	"errors"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/rightscale/terraform-provider-rs/rs/rsc"
)

var deploymentSchema = map[string]*schema.Schema{
	"name": &schema.Schema{
		Description: "name of deployment",
		Type:        schema.TypeString,
		Required:    true,
	},
	"description": &schema.Schema{
		Description: "description of deployment",
		Type:        schema.TypeString,
		Optional:    true,
	},
	"resource_group_href": &schema.Schema{
		Description: "href of the Windows Azure Resource Group attached to the deployment",
		Type:        schema.TypeString,
		Optional:    true,
	},
	"locked": &schema.Schema{
		Description: "whether deployment is locked",
		Type:        schema.TypeBool,
		Optional:    true,
	},
	"server_tag_scope": &schema.Schema{
		Description:  "routing scope for tags for servers in the deployment",
		Type:         schema.TypeString,
		Optional:     true,
		InputDefault: "deployment",
		ForceNew:     true,
		ValidateFunc: func(v interface{}, _ string) (warns []string, errs []error) {
			if v == "" || v == "account" || v == "deployment" {
				return nil, nil
			}
			return nil, []error{errors.New(`server_tag_scope must be "account" or "deployment"`)}
		},
	},
}

func resourceDeployment() *schema.Resource {
	return &schema.Resource{
		Schema: deploymentSchema,
		Read:   resourceRead,
		Exists: resourceExists,
		Delete: resourceDelete,
		Create: resourceDeploymentCreate,
		Update: resourceDeploymentUpdate,
	}
}

func resourceDeploymentCreate(d *schema.ResourceData, m interface{}) error {
	lock, ok := d.GetOk("locked")
	client := m.(rsc.Client)
	res, err := client.Create("rs_cm", "deployment", deploymentFields(d))
	if err != nil {
		return err
	}
	updateSchema(d, res)
	if ok {
		if err := updateLock(d, client, lock.(bool)); err != nil {
			// Attempt to delete previously created deployment
			client.Delete(res.Locator)
			return err
		}
	}
	// set ID last so Terraform does not assume the deployment has been
	// created until all operations have completed successfully.
	d.SetId(res.Locator.Namespace + ":" + res.Locator.Href)
	return nil
}

func resourceDeploymentUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(rsc.Client)
	loc, err := locator(d)
	if err != nil {
		return err
	}
	if err := client.Update(loc, deploymentFields(d)); err != nil {
		return err
	}
	if lock, ok := d.GetOk("locked"); ok {
		return updateLock(d, client, lock.(bool))
	}
	return nil
}

func updateLock(d *schema.ResourceData, client rsc.Client, lock bool) error {
	loc, err := locator(d)
	if err != nil {
		return err
	}
	if lock != d.Get("locked") {
		if lock {
			if err := client.Run(loc, "@res.lock()"); err != nil {
				return err
			}
			d.Set("locked", true)
		} else {
			if err := client.Run(loc, "@res.unlock()"); err != nil {
				return err
			}
			d.Set("locked", false)
		}
	}
	return nil
}

func deploymentFields(d *schema.ResourceData) rsc.Fields {
	fields := rsc.Fields{"name": d.Get("name")}
	if desc, ok := d.GetOk("description"); ok {
		fields["description"] = desc
	}
	if rghref, ok := d.GetOk("resource_group_href"); ok {
		fields["resource_group_href"] = rghref
	}
	if scope, ok := d.GetOk("server_tag_scope"); ok {
		fields["server_tag_scope"] = scope
	}
	return fields
}

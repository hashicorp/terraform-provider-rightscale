package rs

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
	"github.com/rightscale/terraform-provider-rs/rs/rsc"
)

func resourceCMServer() *schema.Resource {
	return &schema.Resource{
		Read:   resourceRead,
		Exists: resourceExists,
		Delete: resourceDelete, // can fail if server is locked - that's what we want
		Create: resourceCMServerCreate,
		Update: resourceCMServerUpdate,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Description: "name of server",
				Type:        schema.TypeString,
				Required:    true,
			},
			"description": &schema.Schema{
				Description: "description of server",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"deployment_href": &schema.Schema{
				Description: "ID of deployment in which to create server",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"instance": &schema.Schema{
				Description: "server instance details",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				ForceNew:    true,
				Elem:        resourceCMInstance(),
			},
			"optimized": &schema.Schema{
				Description:  "A flag indicating whether Instances of this Server should be optimized for high-performance volumes (e.g. Volumes supporting a specified number of IOPS). Not supported in all Clouds.",
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"true", "false"}, false),
			},
		},
	}
}

func resourceCMServerCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(rsc.Client)
	res, err := client.Create("rs_cm", "server", serverFields(d))
	if err != nil {
		return err
	}
	for k, v := range res.Fields {
		d.Set(k, v)
	}

	d.SetId(res.Locator.Namespace + ":" + res.Locator.Href)
	return nil
}

func resourceCMServerUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(rsc.Client)
	loc, err := locator(d)
	if err != nil {
		return err
	}

	if err := client.Update(loc, serverFields(d)); err != nil {
		return handleRSCError(d, err)
	}

	return nil
}

func serverFields(d *schema.ResourceData) rsc.Fields {
	fields := rsc.Fields{}
	if i, ok := d.GetOk("instance"); ok {
		ifs := instanceFields(i.(*schema.ResourceData))
		fields["instance"] = ifs
	}
	for _, f := range []string{"name", "resource_group_href", "server_tag_scope"} {
		if v, ok := d.GetOk(f); ok {
			fields[f] = v
		}
	}
	return fields
}

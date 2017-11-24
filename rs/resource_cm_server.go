package rs

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/rightscale/terraform-provider-rs/rs/rsc"
)

func resourceCMServer() *schema.Resource {
	return &schema.Resource{
		Read:   resourceRead,
		Exists: resourceExists,
		Delete: resourceDelete,
		Create: resourceCreateFunc("rs_cm", "server", serverWriteFields),
		Update: resourceUpdateFunc(serverWriteFields),

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
				Elem:        resourceCMInstance(),
			},
			"optimized": &schema.Schema{
				Description: "A flag indicating whether Instances of this Server should be optimized for high-performance volumes (e.g. Volumes supporting a specified number of IOPS). Not supported in all Clouds.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"links": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeMap},
				Computed: true,
			},
		},
	}
}

func serverWriteFields(d *schema.ResourceData) rsc.Fields {
	fields := rsc.Fields{}
	if i, ok := d.GetOk("instance"); ok {
		fields["instance"] = instanceWriteFields(i.(*schema.ResourceData))
	}
	if o, ok := d.GetOk("optimized"); ok {
		if o.(bool) {
			fields["optimized"] = "true"
		} else {
			fields["optimized"] = "false"
		}
	}
	for _, f := range []string{"name", "resource_group_href", "server_tag_scope"} {
		if v, ok := d.GetOk(f); ok {
			fields[f] = v
		}
	}
	return fields
}

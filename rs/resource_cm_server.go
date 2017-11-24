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
		Create: resourceCreateFunc("rs_cm", "server", serverFields),
		Update: resourceUpdateFunc(serverFields),

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

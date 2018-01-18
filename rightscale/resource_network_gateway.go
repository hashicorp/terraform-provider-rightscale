package rightscale

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
	"github.com/rightscale/terraform-provider-rightscale/rightscale/rsc"
)

func resourceNetworkGateway() *schema.Resource {
	return &schema.Resource{
		Read:   resourceRead,
		Exists: resourceExists,
		Delete: resourceDelete,
		Create: resourceCreateFunc("rs_cm", "network_gateways", networkGatewayWriteFields),
		Update: resourceUpdateFunc(networkGatewayWriteFields),

		Schema: map[string]*schema.Schema{
			"cloud_href": &schema.Schema{
				Description: "ID of cloud in which to create network gateway",
				Type:        schema.TypeString,
				Required:    true,
			},
			"description": &schema.Schema{
				Description: "description of network gateway",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"name": &schema.Schema{
				Description: "name of network gateway",
				Type:        schema.TypeString,
				Required:    true,
			},
			"type": &schema.Schema{
				Description:  "type of network gateway",
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"internet", "vpc"}, false),
			},

			// Read-only fields
			"created_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"links": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeMap},
				Computed: true,
			},
			"resource_uid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"updated_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func networkGatewayWriteFields(d *schema.ResourceData) rsc.Fields {
	fields := rsc.Fields{
		"cloud_href": d.Get("cloud_href"),
		"name":       d.Get("name"),
		"type":       d.Get("type"),
	}
	if desc, ok := d.GetOk("description"); ok {
		fields["description"] = desc
	}
	return rsc.Fields{"network_gateway": fields}
}

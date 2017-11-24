package rs

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/rightscale/terraform-provider-rs/rs/rsc"
)

func resourceCMSubnet() *schema.Resource {
	return &schema.Resource{
		Read:   resourceRead,
		Exists: resourceExists,
		Delete: resourceDelete,
		Create: resourceCMSubnetCreate,
		Update: resourceUpdateFunc(subnetWriteFields),

		Schema: map[string]*schema.Schema{
			"cidr_block": &schema.Schema{
				Description: "range of IP addresses for subnet, this parameter is required for Amazon clouds",
				Type:        schema.TypeString,
				Required:    true,
			},
			"datacenter_href": &schema.Schema{
				Description: "ID of subnet datacenter",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"description": &schema.Schema{
				Description: "description of subnet",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"name": &schema.Schema{
				Description: "name of subnet",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"network_href": &schema.Schema{
				Description: "ID of subnet network",
				Type:        schema.TypeString,
				Required:    true,
			},
			"route_table_href": &schema.Schema{
				Description: "ID of subnet route table",
				Type:        schema.TypeString,
				Required:    true,
			},

			// Read-only fields
			"is_default": &schema.Schema{
				Description: "whether the subnet is the network default subnet",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"links": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeMap},
				Computed: true,
			},
			"resource_uid": &schema.Schema{
				Description: "cloud ID of subnet resource",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"state": {
				Description: "indicates whether subnet is pending, available etc.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"visibility": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceCMSubnetCreate(d *schema.ResourceData, m interface{}) error {
	var rt string
	{
		if r, ok := d.GetOk("route_table_href"); ok {
			rt = r.(string)
		}
	}

	client := m.(rsc.Client)

	// first create subnet with no default route table
	fields := subnetWriteFields(d)
	delete(fields, "route_table_href")
	res, err := client.Create("rs_cm", "subnet", fields)
	if err != nil {
		return err
	}
	for k, v := range res.Fields {
		d.Set(k, v)
	}

	// then update with default route table if any
	if rt != "" {
		d.Set("route_table_href", rt)
		if err := resourceUpdateFunc(subnetWriteFields)(d, client); err != nil {
			// Attempt to delete previously created subnet, ignore errors
			client.Delete(res.Locator)
			return err
		}
	}

	// set ID last so Terraform does not assume the subnet has been
	// created until all operations have completed successfully.
	d.SetId(res.Locator.Namespace + ":" + res.Locator.Href)
	return nil
}

func subnetWriteFields(d *schema.ResourceData) rsc.Fields {
	fields := rsc.Fields{
		"cidr_block":   d.Get("cidr_block"),
		"network_href": d.Get("network_href"),
	}
	for _, f := range []string{
		"datacenter_href", "description", "name",
	} {
		if v, ok := d.GetOk(f); ok {
			fields[f] = v
		}
	}
	return fields
}

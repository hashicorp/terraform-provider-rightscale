package rs

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
	"github.com/rightscale/terraform-provider-rs/rs/rsc"
)

func resourceCMSecurityGroupRule() *schema.Resource {
	return &schema.Resource{
		Read:   resourceRead,
		Exists: resourceExists,
		Delete: resourceDelete,
		Create: resourceCMSecurityGroupRuleCreate,
		Update: resourceUpdateFunc(securityGroupRuleUpdateFields),

		Schema: map[string]*schema.Schema{
			"action": {
				Type:         schema.TypeString,
				Description:  "Allow or deny rule. Defaults to allow. Supported by AzureRM cloud only.",
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"allow", "deny"}, false),
			},
			"cidr_ips": {
				Type:        schema.TypeString,
				Description: "An IP address range in CIDR notation. Required if source_type is 'cidr_ips'.",
				Required:    true,
				ForceNew:    true,
			},
			"description": {
				Type:        schema.TypeString,
				Description: "Description of rule.",
				Optional:    true,
			},
			"direction": {
				Type:         schema.TypeString,
				Description:  "Direction of traffic.",
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"ingress", "egress"}, false),
			},
			"group_name": {
				Type:        schema.TypeString,
				Description: "Name of source Security Group. Required if source_type is 'group'.",
				Optional:    true,
				ForceNew:    true,
			},
			"group_owner": {
				Type:        schema.TypeString,
				Description: "Owner of source Security Group. Required if source_type is 'group'.",
				Optional:    true,
				ForceNew:    true,
			},
			"priority": {
				Type:        schema.TypeInt,
				Description: "Lower takes precedence. Supported by AzureRM cloud only.",
				Optional:    true,
				ForceNew:    true,
			},
			"protocol": {
				Type:        schema.TypeString,
				Description: "Protocol to filter on.",
				Required:    true,
				ForceNew:    true,
			},
			"protocol_details": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Required: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"end_port": {
							Type:        schema.TypeString,
							Description: "End of port range (inclusive). Required if protocol is 'tcp' or 'udp'.",
							Optional:    true,
							ForceNew:    true,
						},
						"icmp_code": {
							Type:        schema.TypeString,
							Description: "ICMP code. Required if protocol is 'icmp'.",
							Optional:    true,
							ForceNew:    true,
						},
						"icmp_type": {
							Type:        schema.TypeString,
							Description: "ICMP type. Required if protocol is 'icmp'.",
							Optional:    true,
							ForceNew:    true,
						},
						"start_port": {
							Type:        schema.TypeString,
							Description: "Start of port range (inclusive). Required if protocol is 'tcp' or 'udp'.",
							Optional:    true,
							ForceNew:    true,
						},
					},
				},
			},
			"security_group_href": {
				Type:        schema.TypeString,
				Description: "ID of parent security group",
				Required:    true,
				ForceNew:    true,
			},

			// Read-only fields
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

func resourceCMSecurityGroupRuleCreate(d *schema.ResourceData, m interface{}) error {
	var desc string
	{
		if de, ok := d.GetOk("description"); ok {
			desc = de.(string)
		}
	}

	client := m.(rsc.Client)

	// first create network with no default route table
	fields := securityGroupRuleCreateFields(d)
	res, err := client.Create("rs_cm", "networks", fields)
	if err != nil {
		return err
	}
	for k, v := range res.Fields {
		d.Set(k, v)
	}

	// then update with description if any
	if desc != "" {
		d.Set("description", desc)
		if err := resourceUpdateFunc(securityGroupRuleUpdateFields)(d, client); err != nil {
			// Attempt to delete previously created network, ignore errors
			client.Delete(res.Locator)
			return err
		}
	}

	// set ID last so Terraform does not assume the network has been
	// created until all operations have completed successfully.
	d.SetId(res.Locator.Namespace + ":" + res.Locator.Href)
	return nil
}

func securityGroupRuleCreateFields(d *schema.ResourceData) rsc.Fields {
	fields := rsc.Fields{
		"protocol":            d.Get("protocol"),
		"security_group_href": d.Get("security_group_href"),
		"source_type":         d.Get("source_type"),
	}
	for _, f := range []string{
		"action", "cidr_ips", "direction", "group_name", "group_owner", "priority",
		"protocol_details",
	} {
		if v, ok := d.GetOk(f); ok {
			fields[f] = v
		}
	}
	return fields
}

func securityGroupRuleUpdateFields(d *schema.ResourceData) rsc.Fields {
	fields := rsc.Fields{}
	if v, ok := d.GetOk("description"); ok {
		fields["description"] = v
	}
	return fields
}

package rs

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
	"github.com/rightscale/terraform-provider-rs/rs/rsc"
)

func resourceCMInstance() *schema.Resource {
	return &schema.Resource{
		Read:   resourceRead,
		Exists: resourceExists,
		Delete: resourceDelete, // can fail if instance is locked - that's what we want
		Create: resourceCMInstanceCreate,
		Update: resourceCMInstanceUpdate,

		Schema: map[string]*schema.Schema{
			"associate_public_ip_address": &schema.Schema{
				Description: "Specify whether or not you want a public IP assigned when this Instance is launched. Only applies to Network-enabled Instances. If this is not specified, it will default to true.",
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
			},
			"cloud_specific_attributes": &schema.Schema{
				Description: "Cloud specific attributes that have no generic rightscale abstraction",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				ForceNew:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"admin_username": &schema.Schema{
							Description: "The user that will be granted administrative privileges. Supported by AzureRM cloud only. For more information, review the documentation.",
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
						},
						"automatic_instance_store_mapping": &schema.Schema{
							Description:  "A flag indicating whether instance store mapping should be enabled. Not supported in all Clouds.",
							Type:         schema.TypeString,
							Optional:     true,
							ForceNew:     true,
							ValidateFunc: validation.StringInSlice([]string{"true", "false"}, false),
						},
						"availability_set": &schema.Schema{
							Description: "Availability set for raw instance. Supported by Azure v2 cloud only.",
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
						},
						"create_boot_volume": &schema.Schema{
							Description:  "If enabled, the instance will launch into volume storage. Otherwise, it will boot to local storage.",
							Type:         schema.TypeString,
							Optional:     true,
							ForceNew:     true,
							ValidateFunc: validation.StringInSlice([]string{"true", "false"}, false),
						},
						"create_default_port_forwarding_rules": &schema.Schema{
							Description:  "Automatically create default port forwarding rules (enabled by default). Supported by Azure cloud only.",
							Type:         schema.TypeString,
							Optional:     true,
							ForceNew:     true,
							ValidateFunc: validation.StringInSlice([]string{"true", "false"}, false),
						},
						"delete_boot_volume": &schema.Schema{
							Description:  "If enabled, the associated volume will be deleted when the instance is terminated.",
							Type:         schema.TypeString,
							Optional:     true,
							ForceNew:     true,
							ValidateFunc: validation.StringInSlice([]string{"true", "false"}, false),
						},
						"disk_gb": &schema.Schema{
							Description: "The size of root disk. Supported by UCA cloud only.",
							Type:        schema.TypeInt,
							Optional:    true,
							ForceNew:    true,
						},
						"ebs_optimized": &schema.Schema{
							Description:  "Whether the instance is able to connect to IOPS-enabled volumes.",
							Type:         schema.TypeString,
							Optional:     true,
							ForceNew:     true,
							ValidateFunc: validation.StringInSlice([]string{"true", "false"}, false),
						},
						"iam_instance_profile": &schema.Schema{
							Description: "The name or ARN of the IAM Instance Profile (IIP) to associate with the instance (Amazon only)",
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
						},
						"keep_alive_id": &schema.Schema{
							Description: "The id of keep alive. Supported by UCA cloud only.",
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
						},
						"keep_alive_url": &schema.Schema{
							Description: "he ulr of keep alive. Supported by UCA cloud only.",
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
						},
						"local_ssd_count": &schema.Schema{
							Description: "Additional local SSDs. Supported by GCE cloud only",
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
						},
						"local_ssd_interface": &schema.Schema{
							Description: "The type of SSD(s) to be created. Supported by GCE cloud only",
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
						},
						"max_spot_price": &schema.Schema{
							Description: "Specify the max spot price you will pay for. Required when 'pricing_type' is 'spot'. Only applies to clouds which support spot-pricing and when 'spot' is chosen as the 'pricing_type'. Should be a Float value >= 0.001, eg: 0.095, 0.123, 1.23, etc...",
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
						},
						"memory_mb": &schema.Schema{
							Description: "The size of instance memory. Supported by UCA cloud only.",
							Type:        schema.TypeInt,
							Optional:    true,
							ForceNew:    true,
						},
						"metadata": &schema.Schema{
							Description: "Extra data used for configuration, in query string format.",
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
						},
						"num_cores": &schema.Schema{
							Description: "The number of instance cores. Supported by UCA cloud only.",
							Type:        schema.TypeInt,
							Optional:    true,
							ForceNew:    true,
						},
						"placement_tenancy": &schema.Schema{
							Description:  "The tenancy of the server you want to launch. A server with a tenancy of dedicated runs on single-tenant hardware and can only be launched into a VPC.",
							Type:         schema.TypeString,
							Optional:     true,
							ForceNew:     true,
							ValidateFunc: validation.StringInSlice([]string{"default", "dedicated"}, false),
						},
						"preemptible": &schema.Schema{
							Description:  "Launch a preemptible instance. A preemptible instance costs much less, but lasts only 24 hours. It can be terminated sooner due to system demands. Supported by GCE cloud only.",
							Type:         schema.TypeString,
							Optional:     true,
							ForceNew:     true,
							ValidateFunc: validation.StringInSlice([]string{"true", "false"}, false),
						},
						"pricing_type": &schema.Schema{
							Description:  "Specify whether or not you want to utilize 'fixed' (on-demand) or 'spot' pricing. Defaults to 'fixed' and only applies to clouds which support spot instances. Can only be set on when creating a new Instance, Server, or ServerArray, or when updating a Server or ServerArray's next_instance.WARNING: By using spot pricing, you acknowledge that your instance/server/array may not be able to be launched (and arrays may be unable to grow) as newly launched instances might be stuck in bidding, and/or existing instances may be terminated at any time, due to the cloud's spot pricing changes and availability.",
							Type:         schema.TypeString,
							Optional:     true,
							ForceNew:     true,
							ValidateFunc: validation.StringInSlice([]string{"fixed", "spot"}, false),
						},
						"root_volume_performance": &schema.Schema{
							Description: "The number of IOPS (I/O Operations Per Second) this root volume should support. Only available on clouds supporting performance provisioning.",
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
						},
						"root_volume_size": &schema.Schema{
							Description: "The size for root disk. Not supported in all Clouds.",
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
						},
						"root_volume_type_uid": &schema.Schema{
							Description: "The type of root volume for instance. Only available on clouds supporting root volume type.",
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
						},
						"service_account": &schema.Schema{
							Description: "Email of service account for instance. Scope will default to cloud-platform. Supported by GCE cloud only.",
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
						},
					},
				},
			},
			"datacenter_href": &schema.Schema{
				Description: "The ID of the instance datacenter",
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
			},
			"deployment_href": &schema.Schema{
				Description: "The ID of the instance deployment",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"image_href": &schema.Schema{
				Description: "The ID of the instance image",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"instance_type_href": &schema.Schema{
				Description: "The ID of the instance type",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"ip_forwarding_enabled": &schema.Schema{
				Description: "Allows this Instance to send and receive network traffic when the source and destination IP addresses do not match the IP address of this Instance.",
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
			},
			"kernel_image_href": &schema.Schema{
				Description: "The ID of the instance kernel image.",
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
			},
			"name": &schema.Schema{
				Description: "The name of the instance",
				Type:        schema.TypeString,
				Required:    true,
			},
			"placement_group_href": &schema.Schema{
				Description: "The placement group to launch the instance in. Not supported by all clouds & instance types.",
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
			},
			"ramdisk_image_href": &schema.Schema{
				Description: "The ID of the ramdisk image",
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
			},
			"security_group_hrefs": &schema.Schema{
				Description: "The IDs of the security groups",
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
				ForceNew:    true,
			},
			"ssh_key_href": &schema.Schema{
				Description: "The ID of the SSH key to use",
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
			},
			"subnet_hrefs": &schema.Schema{
				Description: "The IDs of the instance subnets",
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
			},
			"user_data": &schema.Schema{
				Description: "User data that RightScale automatically passes to your instance at boot time",
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
			},
		},
	}
}

func resourceCMInstanceCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(rsc.Client)
	res, err := client.Create("rs_cm", "instance", instanceFields(d))
	if err != nil {
		return err
	}
	for k, v := range res.Fields {
		d.Set(k, v)
	}
	d.SetId(res.Locator.Namespace + ":" + res.Locator.Href)
	return nil
}

func resourceCMInstanceUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(rsc.Client)
	loc, err := locator(d)
	if err != nil {
		return err
	}

	if err := client.Update(loc, instanceFields(d)); err != nil {
		return handleRSCError(d, err)
	}

	return nil
}

func instanceFields(d *schema.ResourceData) rsc.Fields {
	fields := rsc.Fields{}
	for _, f := range []string{
		"name", "instance_type_href", "image_href", "deployment_href",
		"datacenter_href", "kernel_image_href", "placement_group_href",
		"ramdisk_image_href", "ssh_key_href", "security_group_hrefs",
		"subnet_hrefs", "user_data",
	} {
		if v, ok := d.GetOk(f); ok {
			fields[f] = v
		}
	}
	if a, ok := d.GetOk("associate_public_ip_address"); ok {
		if a.(bool) {
			fields["associate_public_ip_address"] = "true"
		} else {
			fields["associate_public_ip_address"] = "false"
		}
	}
	if e, ok := d.GetOk("ip_forwarding_enabled"); ok {
		if e.(bool) {
			fields["ip_forwarding_enabled"] = "true"
		} else {
			fields["ip_forwarding_enabled"] = "false"
		}
	}
	if a, ok := d.GetOk("cloud_specific_attributes"); ok {
		fields["cloud_specific_attributes"] = a.([]interface{})[0]
	}
	return rsc.Fields{"instance": fields}
}

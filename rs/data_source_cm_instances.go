package rs

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
	"github.com/rightscale/terraform-provider-rs/rs/rsc"
)

// Examples:
//
// data "rs_cm_instance" "mysql" {
//     filter {
//         resource_uid = "vpc-c31ee987"
//     }
//     cloud = ${data.rs_cm_cloud.ec2_us_east_1.id}
// }
//
// data "rs_cm_instance" "worker_2" {
//     filter {
//         name = "Worker #2"
//     }
//     server_array = ${data.rs_cm_server_array.workers}
// }

func dataSourceInstances() *schema.Resource {
	return &schema.Resource{
		Read: resourceInstanceRead,

		Schema: map[string]*schema.Schema{
			"cloud": {
				Type:        schema.TypeString,
				Description: "ID of instance cloud resource, exclusive with 'server_array'",
				Optional:    true,
				ForceNew:    true,
			},
			"server_array": {
				Type:        schema.TypeString,
				Description: "ID of instance server array resource, exclusive with 'cloud'",
				Optional:    true,
				ForceNew:    true,
			},
			"filter": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "name of instance, uses partial match",
							Optional:    true,
							ForceNew:    true,
						},
						"datacenter": {
							Type:        schema.TypeString,
							Description: "ID of the instance datacenter resource",
							Optional:    true,
							ForceNew:    true,
						},
						"os_platform": {
							Type:         schema.TypeString,
							Description:  "OS platform of instance",
							Optional:     true,
							ForceNew:     true,
							ValidateFunc: validation.StringInSlice([]string{"windows", "linux"}, false),
						},
						"parent": {
							Type:        schema.TypeString,
							Description: "ID of instance server or server array parent resource",
							Optional:    true,
							ForceNew:    true,
						},
						"server_template": {
							Type:        schema.TypeString,
							Description: "ID of instance server template resource",
							Optional:    true,
							ForceNew:    true,
						},
						"state": {
							Type:        schema.TypeString,
							Description: "state of instance",
							Optional:    true,
							ForceNew:    true,
						},
						"placement_group": {
							Type:        schema.TypeString,
							Description: "ID of instance placement group resource",
							Optional:    true,
							ForceNew:    true,
						},
						"public_dns_name": {
							Type:        schema.TypeString,
							Description: "Public DNS name of instance",
							Optional:    true,
							ForceNew:    true,
						},
						"private_dns_name": {
							Type:        schema.TypeString,
							Description: "Private DNS name of instance",
							Optional:    true,
							ForceNew:    true,
						},
						"public_ip": {
							Type:        schema.TypeString,
							Description: "Public IP of instance",
							Optional:    true,
							ForceNew:    true,
						},
						"private_ip": {
							Type:        schema.TypeString,
							Description: "Private IP instance",
							Optional:    true,
							ForceNew:    true,
						},
						"deployment": {
							Type:        schema.TypeString,
							Description: "ID of deployment resource that owns instance",
							Optional:    true,
							ForceNew:    true,
						},
						"resource_uid": {
							Type:        schema.TypeString,
							Description: "cloud ID of instance, e.g. 'vpc-2124fe46'",
							Optional:    true,
							ForceNew:    true,
						},
					},
				},
			},
			"name": {
				Type:     schema.TypeString,
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
			"created_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"updated_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"pricing_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"cloud_specific_attributes": {
				Type:     schema.TypeMap,
				Computed: true,
			},
			"public_ip_addresses": {
				Type:     schema.TypeList,
				Elem:     schema.Schema{Type: schema.TypeString},
				Computed: true,
			},
			"private_ip_addresses": {
				Type:     schema.TypeList,
				Elem:     schema.Schema{Type: schema.TypeString},
				Computed: true,
			},
			"locked": {
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}

func resourceInstanceRead(d *schema.ResourceData, m interface{}) error {
	client := m.(rsc.Client)
	var loc *rsc.Locator
	if cloud, ok := d.GetOk("cloud"); ok {
		loc = &rsc.Locator{Namespace: "rs_cm", Href: cloud.(string)}
	} else if array, ok := d.GetOk("server_array"); ok {
		loc = &rsc.Locator{Namespace: "rs_cm", Href: array.(string)}
	} else {
		return fmt.Errorf("instance data source must specify one of 'cloud' or 'server_array'")
	}
	res, err := client.List(loc, "instances", cmFilters(d))
	if err != nil {
		return err
	}

	if len(res) == 0 {
		return nil
	}
	for k, v := range res[0].Fields {
		d.Set(k, v)
	}
	d.SetId(res[0].Locator.Href)
	return nil
}

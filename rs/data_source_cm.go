package rs

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/rightscale/terraform-provider-rs/rs/rsc"
)

// cmFilters is a helper function that returns fields representing valid
// RightScale API 1.5 filter parameters built from the resource data "filter"
// field.
func cmFilters(d *schema.ResourceData) rsc.Fields {
	var filters rsc.Fields
	if f, ok := d.GetOk("filter"); ok {
		actual := f.(map[string]interface{})
		mapped := make(map[string]interface{}, len(actual))
		for k, v := range actual {
			mapped[mapCMFilterKey(k)] = v
		}
		filters = rsc.Fields{"filter[]": mapped}
	}
	return filters
}

// mapCMFilterKey returns the RightScale API filter key for the given terraform filter key.
func mapCMFilterKey(k string) string {
	if m, ok := mappedCMKeys[k]; ok {
		return m
	}
	return k
}

// mappedCMKeys maps Terraform filter keys to RightScale API 1.5.
var mappedCMKeys = map[string]string{
	"cloud":                  "cloud_href",
	"datacenter":             "datacenter_href",
	"deployment":             "deployment_href",
	"instance":               "instance_href",
	"multi_cloud_image":      "multi_cloud_image_href",
	"parent":                 "parent_href",
	"parent_volume":          "parent_volume_href",
	"parent_volume_snapshot": "parent_volume_snapshot_href",
	"placement_group":        "placement_group_href",
	"server_template":        "server_template_href",
}

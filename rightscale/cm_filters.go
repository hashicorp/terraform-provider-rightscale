package rightscale

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/rightscale/terraform-provider-rightscale/rightscale/rsc"
)

// cmFilters is a helper function that returns fields representing valid
// RightScale API 1.5 filter parameters built from the resource data "filter"
// field.
func cmFilters(d *schema.ResourceData) rsc.Fields {
	var filters rsc.Fields
	if f, ok := d.GetOk("filter"); ok {
		filters = rsc.Fields{"filter[]": f}
	}
	return filters
}

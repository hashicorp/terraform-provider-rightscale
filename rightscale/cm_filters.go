package rightscale

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/rightscale/terraform-provider-rightscale/rightscale/rsc"
)

// cmFilters is a helper function that returns fields representing valid
// RightScale rcl filter parameters built from the resource data "filter"
// field.
func cmFilters(d *schema.ResourceData) rsc.Fields {
	var filters rsc.Fields
	var arrify []string
	if f, ok := d.GetOk("filter"); ok {
		filtersList := f.([]interface{})
		for _, filterIF := range filtersList {
			v := filterIF.(map[string]interface{})
			for filter, value := range v {
				if value != "" {
					z := fmt.Sprintf("%v==%v", filter, value)
					arrify = append(arrify, z)
				}
			}
		}
		filters = rsc.Fields{"filter": arrify}
	}
	return filters
}

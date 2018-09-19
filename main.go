package main

import (
	"github.com/hashicorp/terraform/plugin"
	rs "github.com/terraform-providers/terraform-provider-rightscale/rightscale"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{ProviderFunc: rs.Provider})
}

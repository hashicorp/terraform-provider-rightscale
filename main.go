package main

import (
	"github.com/hashicorp/terraform/plugin"
	rs "github.com/rightscale/terraform-provider-rightscale/rightscale"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{ProviderFunc: rs.Provider})
}

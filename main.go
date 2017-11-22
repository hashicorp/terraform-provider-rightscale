package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/rightscale/terraform-provider-rs/rs"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{ProviderFunc: rs.Provider})
}

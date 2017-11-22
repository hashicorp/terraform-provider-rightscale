package rs

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/rightscale/terraform-provider-rs/rs/rsc"
)

// Provider returns the RightScale terraform provider.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_token": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("RS_API_TOKEN", nil),
				Description: "API token used to authenticate the RightScale API requests",
			},
			"project_id": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("RS_PROJECT_ID", nil),
				Description: "ID of RightScale project used to provision resources",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"rs_cm_deployment": resourceDeployment(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(r *schema.ResourceData) (interface{}, error) {
	token := r.Get("api_token").(string)
	if token == "" {
		return nil, errors.New("missing 'api_token'")
	}
	project, err := strconv.Atoi(r.Get("project_id").(string))
	if err != nil {
		return nil, fmt.Errorf("failed to parse project ID: %s", err)
	}

	return rsc.New(token, project)
}

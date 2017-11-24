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

		DataSourcesMap: map[string]*schema.Resource{
			"rs_cm_cloud":             dataSourceCMCloud(),
			"rs_cm_datacenter":        dataSourceCMDatacenter(),
			"rs_cm_instance":          dataSourceCMInstance(),
			"rs_cm_instance_type":     dataSourceCMInstanceType(),
			"rs_cm_multi_cloud_image": dataSourceCMMultiCloudImage(),
			"rs_cm_network":           dataSourceCMNetwork(),
			"rs_cm_security_group":    dataSourceCMSecurityGroup(),
			"rs_cm_server_template":   dataSourceCMServerTemplate(),
			"rs_cm_ssh_key":           dataSourceCMSSHKey(),
			"rs_cm_subnet":            dataSourceCMSubnet(),
			"rs_cm_volume":            dataSourceCMVolume(),
			"rs_cm_volume_snapshot":   dataSourceCMVolumeSnapshot(),
			"rs_cm_volume_type":       dataSourceCMVolumeType(),
		},

		ResourcesMap: map[string]*schema.Resource{
			"rs_cm_deployment": resourceCMDeployment(),
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

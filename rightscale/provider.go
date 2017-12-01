package rightscale

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/rightscale/terraform-provider-rightscale/rightscale/rsc"
)

// Provider returns the RightScale terraform provider.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"rightscale_api_token": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("RIGHTSCALE_API_TOKEN", nil),
				Description: "API token used to authenticate the RightScale API requests",
			},
			"rightscale_project_id": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("RIGHTSCALE_PROJECT_ID", nil),
				Description: "ID of RightScale project used to provision resources",
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			"rightscale_cm_cloud":             dataSourceCMCloud(),
			"rightscale_cm_datacenter":        dataSourceCMDatacenter(),
			"rightscale_cm_image":             dataSourceCMImage(),
			"rightscale_cm_instance":          dataSourceCMInstance(),
			"rightscale_cm_instance_type":     dataSourceCMInstanceType(),
			"rightscale_cm_multi_cloud_image": dataSourceCMMultiCloudImage(),
			"rightscale_cm_network":           dataSourceCMNetwork(),
			"rightscale_cm_security_group":    dataSourceCMSecurityGroup(),
			"rightscale_cm_server_template":   dataSourceCMServerTemplate(),
			"rightscale_cm_ssh_key":           dataSourceCMSSHKey(),
			"rightscale_cm_subnet":            dataSourceCMSubnet(),
			"rightscale_cm_volume":            dataSourceCMVolume(),
			"rightscale_cm_volume_snapshot":   dataSourceCMVolumeSnapshot(),
			"rightscale_cm_volume_type":       dataSourceCMVolumeType(),
		},

		ResourcesMap: map[string]*schema.Resource{
			"rightscale_cm_deployment":          resourceCMDeployment(),
			"rightscale_cm_instance":            resourceCMInstance(),
			"rightscale_cm_network":             resourceCMNetwork(),
			"rightscale_cm_network_gateway":     resourceCMNetworkGateway(),
			"rightscale_cm_route_table":         resourceCMRouteTable(),
			"rightscale_cm_security_group":      resourceCMSecurityGroup(),
			"rightscale_cm_security_group_rule": resourceCMSecurityGroupRule(),
			"rightscale_cm_server":              resourceCMServer(),
			"rightscale_cm_server_array":        resourceCMServerArray(),
			"rightscale_cm_ssh_key":             resourceCMSSHKey(),
			"rightscale_cm_subnet":              resourceCMSubnet(),
			"rightscale_cwf_process":            resourceCWFProcess(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(r *schema.ResourceData) (interface{}, error) {
	token := r.Get("rightscale_api_token").(string)
	if token == "" {
		return nil, errors.New("missing 'api_token'")
	}
	project, err := strconv.Atoi(r.Get("rightscale_project_id").(string))
	if err != nil {
		return nil, fmt.Errorf("failed to parse project ID: %s", err)
	}

	return rsc.New(token, project)
}

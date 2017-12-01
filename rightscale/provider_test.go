package rightscale

import (
	"os"
	"strings"
	"testing"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"gopkg.in/rightscale/rsc.v6/cm15"
	"gopkg.in/rightscale/rsc.v6/rsapi"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

const (
	credEnvVar     = "RIGHTSCALE_API_TOKEN"
	projectEnvVars = "RIGHTSCALE_PROJECT_ID"
)

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"rightscale": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ terraform.ResourceProvider = Provider()
}

func testAccPreCheck(t *testing.T) {
	if v := envSearch(credEnvVar); v == "" {
		t.Fatalf("%s must be set for acceptance tests", credEnvVar)
	}

	if v := envSearch(projectEnvVars); v == "" {
		t.Fatalf("%s must be set for acceptance tests", projectEnvVars)
	}
}

// getCMClient returns a low level API 1.5 client.
func getCMClient(s *terraform.State) *cm15.API {
	type cmClient interface {
		API() *rsapi.API
	}

	c := testAccProvider.Meta().(cmClient)
	return &cm15.API{API: c.API()}
}

// testAccPreCheck ensures at least one of the project env variables is set.
func getTestProjectFromEnv() string {
	return envSearch(projectEnvVars)
}

// testAccPreCheck ensures at least one of the credentials env variables is set.
func getTestCredsFromEnv() string {
	return envSearch(credEnvVar)
}

func getHrefFromID(id string) string {
	return strings.Split(id, ":")[1]
}

func envSearch(k string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return ""
}

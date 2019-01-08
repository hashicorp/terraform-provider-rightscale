package rightscale

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/terraform-providers/terraform-provider-rightscale/rightscale/rsc"
)

func TestAccRightScaleCWFProcess_basic(t *testing.T) {
	t.Parallel()

	const src = `
define main() return $out do
    $out = 42
end
`
	var process rsc.Process

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCWFProcessDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCWFProcess_basic(src),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCWFProcessExists("rightscale_cwf_process.foobar", &process),
					testAccCheckCWFProcessOutput(&process, []string{"$out"}, []interface{}{string("42")}),
					testAccCheckCWFProcessStatus(&process, "completed"),
				),
			},
		},
	})
}

func TestAccRightScaleCWFProcess_params(t *testing.T) {
	t.Parallel()
	var process rsc.Process

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCWFProcessDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCWFProcess_params(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCWFProcessExists("rightscale_cwf_process.foobar", &process),
					testAccCheckCWFProcessOutput(&process, []string{"$out1", "$out2", "$out3"}, []interface{}{"foobared", "42", "true"}),
					testAccCheckCWFProcessStatus(&process, "completed"),
				),
			},
		},
	})
}

func TestAccRightScaleCWFProcess_collection(t *testing.T) {
	t.Parallel()
	var process rsc.Process

	sgHref := os.Getenv("RIGHTSCALE_SECURITY_GROUP_HREF")

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCWFProcessDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCWFProcess_collection(sgHref),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCWFProcessExists("rightscale_cwf_process.collection", &process),
					testAccCheckCWFProcessOutput(&process, []string{"$out"}, []interface{}{sgHref}),
					testAccCheckCWFProcessStatus(&process, "completed"),
				),
			},
		},
	})
}

func TestAccRightScaleCWFProcess_error(t *testing.T) {
	t.Parallel()

	const src = `
define main() return $out do
    $out = 42
    raise "test error"
end
`

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCWFProcessDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccCWFProcess_basic(src),
				ExpectError: regexp.MustCompile("test error"),
			},
		},
	})
}

func testAccCheckCWFProcessExists(n string, p *rsc.Process) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}

		c := testAccProvider.Meta().(rsc.Client)
		proc, err := c.GetProcess(rs.Primary.ID)
		if err != nil {
			return err
		}

		*p = *proc

		return nil
	}
}

func testAccCheckCWFProcessOutput(p *rsc.Process, names []string, vals []interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if len(p.Outputs) != len(names) {
			return fmt.Errorf("got %d output(s), expected %d", len(p.Outputs), len(names))
		}
		for i, e := range names {
			out, ok := p.Outputs[e]
			if !ok {
				return fmt.Errorf("output %q is missing", e)
			}
			if out != vals[i] {
				return fmt.Errorf("output %q has value %T: %#v, expected %T: %#v", e, out, out, vals[i], vals[i])
			}
		}
		return nil
	}
}

func testAccCheckCWFProcessStatus(p *rsc.Process, status string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if p.Status != status {
			return fmt.Errorf("got status %q, expected %q", p.Status, status)
		}
		return nil
	}
}

func testAccCheckCWFProcessError(p *rsc.Process, err string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if !strings.Contains(p.Error.Error(), err) {
			return fmt.Errorf("got error %q, expected to contain %q", p.Error, err)
		}
		return nil
	}
}

func testAccCWFProcess_basic(src string) string {
	return fmt.Sprintf(`
resource "rightscale_cwf_process" "foobar" {
	source =  <<EOF
%s
EOF
}
`, src)
}

func testAccCWFProcess_params() string {
	return fmt.Sprintf(`
resource "rightscale_cwf_process" "foobar" {
	source     = <<EOF
	define main($s, $list, $bool) return $out1, $out2, $out3 do
	$out1 = $s + "ed"
	$out2 = $list[0] + $list[1]
	if $bool
	  $out3 = "true"
	end
end
EOF
	parameters = [
		{
			kind = "string"
			value = "foobar"
		},
		{
			kind = "array"
			value = "[ 11, 31 ]"
		},
		{
			kind = "boolean"
			value = "true"
		}
	]
}
`)
}

func testAccCWFProcess_collection(sgHref string) string {
	return fmt.Sprintf(`
		variable "sec_group" {
			type    = "map"
			default = {
			  "namespace" = "rs_cm"
			  "type" = "security_groups"
			  "hrefs" = ["%s"]
			  "details" = [{
					"description" = "A security group"
			  }]
			}
		  }

resource "rightscale_cwf_process" "collection" {
	source     = <<EOF
	define main(@collection) return $out do
		$json = to_object(@collection)
		$out = $json["hrefs"][0]
	end
EOF
	parameters = [
		{
			kind = "collection"
			value = "${jsonencode(var.sec_group)}"
		},
	]
}
`, sgHref)
}

func testAccCheckCWFProcessDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(rsc.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "rightscale_cwf_process" {
			continue
		}

		_, err := c.GetProcess(rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("process still exists")
		}
		if err != rsc.ErrNotFound {
			return fmt.Errorf("failed to get process: %s", err)
		}
	}
	return nil
}

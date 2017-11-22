package rs

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/rightscale/terraform-provider-rs/rs/rsc"
)

func resourceRead(d *schema.ResourceData, m interface{}) error {
	client := m.(rsc.Client)
	loc, err := locator(d)
	if err != nil {
		return err
	}
	res, err := client.Get(loc)
	if err != nil {
		return err
	}
	updateSchema(d, res)
	return nil
}

func resourceDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(rsc.Client)
	loc, err := locator(d)
	if err != nil {
		return err
	}
	return client.Delete(loc)
}

func resourceExists(d *schema.ResourceData, m interface{}) (bool, error) {
	client := m.(rsc.Client)
	loc, err := locator(d)
	if err != nil {
		return false, err
	}
	res, err := client.Get(loc)
	if err != nil {
		return false, err
	}
	return res != nil, nil
}

// locator builds a locator from a schema.
func locator(d *schema.ResourceData) (*rsc.Locator, error) {
	parts := strings.Split(d.Id(), ":")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid resource ID %q", d.Id())
	}
	return &rsc.Locator{Namespace: parts[0], Href: parts[1]}, nil
}

// updateSchema stores the resource information in the given schema.
func updateSchema(d *schema.ResourceData, res *rsc.Resource) {
	for k, v := range res.Fields {
		d.Set(k, v)
	}
}

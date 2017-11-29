package rs

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/rightscale/terraform-provider-rightscale/rightscale/rsc"
)

func resourceCreateFunc(namespace, typ string, fieldsFunc func(*schema.ResourceData) rsc.Fields) func(d *schema.ResourceData, m interface{}) error {
	return func(d *schema.ResourceData, m interface{}) error {
		client := m.(rsc.Client)
		res, err := client.Create(namespace, typ, fieldsFunc(d))
		if err != nil {
			return err
		}
		for k, v := range res.Fields {
			d.Set(k, v)
		}
		d.SetId(res.Locator.Namespace + ":" + res.Locator.Href)
		return nil
	}
}

func resourceUpdateFunc(fieldsFunc func(*schema.ResourceData) rsc.Fields) func(d *schema.ResourceData, m interface{}) error {
	return func(d *schema.ResourceData, m interface{}) error {
		client := m.(rsc.Client)
		loc, err := locator(d)
		if err != nil {
			return err
		}

		if err := client.Update(loc, fieldsFunc(d)); err != nil {
			return handleRSCError(d, err)
		}

		return nil
	}
}

func resourceRead(d *schema.ResourceData, m interface{}) error {
	client := m.(rsc.Client)
	loc, err := locator(d)
	if err != nil {
		return err
	}
	res, err := client.Get(loc)
	if err != nil {
		return handleRSCError(d, err)
	}
	for k, v := range res.Fields {
		d.Set(k, v)
	}
	return nil
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

func resourceDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(rsc.Client)
	loc, err := locator(d)
	if err != nil {
		return err
	}
	return client.Delete(loc)
}

// handleRSCError checks if the error is rsc.ErrNotFound and if so removes the
// id from the resource data fields to let Terraform know that the resource is
// gone. In this case it also returns nil instead of the original error for
// graceful handling by Terraform. Otherwise the original error is returned.
func handleRSCError(d *schema.ResourceData, err error) error {
	if err == rsc.ErrNotFound {
		d.SetId("")
		return nil
	}
	return err
}

// locator builds a locator from a schema.
func locator(d *schema.ResourceData) (*rsc.Locator, error) {
	parts := strings.Split(d.Id(), ":")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid resource ID %q", d.Id())
	}
	return &rsc.Locator{Namespace: parts[0], Href: parts[1]}, nil
}

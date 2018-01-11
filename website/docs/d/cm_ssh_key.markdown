---
layout: "rightscale"
page_title: "Rightscale: cm_ssh_key"
sidebar_current: "docs-rightscale-datasource-cm-ssh-key"
description: |-
  Defines an ssh key datasource to operate against. 
---

# rightscale_cm_ssh_key

Use this data source to get the ID of an existing ssh key for use in other
resources.

## Example Usage

```hcl
data "rightscale_cm_ssh_key" "infra-ssh-key" {
  filter {
    name = "infra"
  }
  cloud_href = ${data.rightscale_cm_cloud.ec2_us_east_1.id}
}
```

## Argument Reference

The following arguments are supported:

* `cloud_href` - (Required) The ID of the cloud with the ssh key you want.

The `filter` block supports:

* `name` - (Optional) SSH key name.  Pattern match. 

* `resource_uid` - (Optional) Href/ID of the SSH key.

## Attributes Reference

The following attributes are exported:

* `name` - Official cloud name as displayed in cm platform.

* `resource_uid` - Href/ID of the SSH key.

* `links` - Hrefs of related API resources.
---
layout: "rightscale"
page_title: "Rightscale: volume"
sidebar_current: "docs-rightscale-datasource-volume"
description: |-
  Defines a volume datasource to operate against.
---

# rightscale_volume

Use this data source to get the ID of an existing volume for use in other resources.

## Example Usage 1: Basic configuration of a volume data source

```hcl
data "rightscale_volume" "a_volume" {
  cloud_href = "/api/clouds/1"

  filter {
    name = "my_volume"
  }
}

output "volume name" {
  value = "${data.rightscale_volume.crunis_volume.name}"
}

output "volume ID" {
  value = "${data.rightscale_volume.crunis_volume.resource_uid}"
}
```

## Argument Reference

The following arguments are supported:

The `filter` block supports:

* `name` - The name of the volume

* `description` - The description of the volume

* `resource_uid` - The resource_uid of the volume

* `deployment_href` - The href of the deployment that contains de volume

* `datacenter_href` - The href of the datacenter that holds the volume

* `parent_volume_snapshot_href` - The href of snapshot the volume was created of

## Attributes Reference

The following attributes are exported:

* `name` - The name of the volume

* `description` - The description of the volume

* `links` - Hrefs of related API resources

* `resource_uid` - The resource_uid of the volume

* `size` - The volume size

* `status` - The volume Status

* `updated_at` - Last update of the volume

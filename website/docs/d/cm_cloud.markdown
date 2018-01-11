---
layout: "rightscale"
page_title: "Rightscale: cm_cloud"
sidebar_current: "docs-rightscale-datasource-cm-cloud"
description: |-
  Defines a cloud datasource to operate against. 
---

# rightscale_cm_cloud

Use this data source to get the ID of a registered clouds for use in other
resources.  Registration of clouds in a given RightScale account will need to have been executed ahead of time to define it as a cloud datasource. 

## Example Usage

```hcl
data "rightscale_cm_cloud" "ec2_us_oregon" {
  filter {
    name = "EC2 us-west-2"
    cloud_type = "amazon"
  }
}

data "rightscale_cm_cloud" "azure_us_east" {
  filter {
    name = "Azure East US"
    cloud_type = "azure"
  }
}
...
```

## Argument Reference

**Note - an empty config block IS valid and will return the first cloud object available in your account.**

The following arguments are supported:

The `filter` block supports:

* `name` - (Optional) Cloud name as displayed in cm platform.  Pattern match. 

* `description` - (Optional) Cloud description as displayed in cm platform.  Pattern match.

* `cloud_type` - (Optional) Cloud type as referenced in cm platform.  Common types include: amazon, google, azure, and vscale.  See  [supportedCloudTypes](https://github.com/rightscale/terraform-provider-rightscale/blob/master/rightscale/data_source_cm_cloud.go#L95) for complete list.

## Attributes Reference

The following attributes are exported:

* `name` - Official cloud name as displayed in cm platform.

* `display_name` - Display name for cloud as displayed in cm platform.

* `description` - Cloud description as displayed in cm platform.

* `cloud_type` - Cloud type as referenced in cm platform. 

* `links` - Hrefs of related API resources.
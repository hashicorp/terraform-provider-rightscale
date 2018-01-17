---
layout: "rightscale"
page_title: "Provider: Rightscale"
sidebar_current: "docs-rightscale-index"
description: |-
  The Rightscale provider is used to interact with the the RightScale Cloud Management Platform. The provider needs to be configured with the proper credentials before it can be used.
---

# RightScale Provider

The Rightscale provider is used to interact with the the RightScale Cloud Management Platform.

The provider needs to be configured with the proper credentials before it can be used.

Use the navigation to the left to read about the available data sources.

## Example Usage

```hcl
provider "rightscale" {
  rightscale_api_token = "${var.rightscale_api_token}"
  rightscale_project_id = "${var.rightscale_account_id}"
}

data "rightscale_cloud" "ec2_us_oregon" {
  filter {
    name = "EC2 us-west-2"
    cloud_type = "amazon"
  }
}

data "rightscale_cloud" "azure_us_east" {
  filter {
    name = "Azure East US"
    cloud_type = "azure"
  }
}

resource "rightscale_instance" "test-instance-oregon" {
  cloud_href = "${data.rightscale_cloud.ec2_us_oregon.id}"
  name = ...
  ...
}

resource "rightscale_instance" "test-instance-east" {
  cloud_href = "${data.rightscale_cloud.azure_us_east.id}"
  name = ...
  ...
}

```
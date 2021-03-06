---
layout: "rightscale"
page_title: "Rightscale: image"
sidebar_current: "docs-rightscale-datasource-image"
description: |-
  Defines an image datasource to operate against.
---

# rightscale_image

Use this data source to locate and extract info about an existing [image](http://reference.rightscale.com/api1.5/resources/ResourceImages.html) to pass to other rightscale resources. Sets default filter scope to own account, but allows for public searching if specified in filter block.

## Example Usage #1 - Finding specific AMI in own account based on resource_uid

```hcl
data "rightscale_image" "my_sweet_ami" {
  cloud_href = "${data.rightscale_cloud.ec2_us_oregon.href}"
  filter {
    resource_uid = "ami-abcdefg"
  }
}

data "rightscale_cloud" "ec2_us_oregon" {
  filter {
    name = "EC2 us-west-2"
    cloud_type = "amazon"
  }
}
...
```

## Example Usage #2 - Finding public image in cloud based on filters on name, description, etc.

Warning: The more images a cloud has public, the longer this filter call will take.  Consider multiple filters to narrow the scope.

```hcl
data "rightscale_cloud" "ec2_us_oregon" {
  filter {
    name = "EC2 us-west-2"
    cloud_type = "amazon"
  }
}

data "rightscale_image" "my_sweet_ami" {
  cloud_href = "${data.rightscale_cloud.ec2_us_oregon.href}"
  visibility = "public"
  filter {
    name = "My Super Great AMI"
    os_platform = "linux"
    description = "AMI Image built from CI that does great things"
  }
}
...
```

## Argument Reference

**Note - omitting the filter block IS valid and will return the first private image object available in the specific cloud and your account.  Probably not what you really want.**

The following arguments are supported:

* `cloud_href` - (Required) The Href of the cloud with the image you want.

* `filter` - (Optional) block supports:

  * `visibility` (Optional) Image visibility as displayed in cm platform.  Options are "private" or "public."  Defaults to "private."  A public search will greatly increase execution time and result set size, so care should be taken when toggling this argument.

  * `resource_uid` (Optional) Image resource_uid.  If this filter option is set, additional retry logic will be enabled to wait up to 5 minutes for cloud resources to be polled and populated for use.

  * `name` - (Optional) Image name as displayed in cm platform.  Pattern match.

  * `description` - (Optional) Image description as displayed in cm platform.  Pattern match.

  * `image_type` - (Optional) Image type as referenced in cm platform. This will be either "machine", "machine_azure", "ramdisk" or "kernel".

  * `os_platform` - (Optional) Image OS platform as referenced in cm platform.  This will either be "windows" or "linux."

  * `cpu_architecture` - (Optional) Image CPU architecture as referenced in cm platform.  Generally "x64_64", etc.  Pattern match.

## Attributes Reference

The following attributes are exported:

* `visibility` - Image visibility as displayed in cm platform.

* `resource_uid` - Image unique resource identifier as displayed in cm platform.

* `name` - Image name as displayed in cm platform.

* `description` - Image description as displayed in cm platform.

* `cpu_architecture` - Image CPU architecture as referenced in cm platform.

* `os_platform` - Image OS platform as referenced in cm platform.

* `root_device_storage` - Image root device storage as reported in cm platform.  Eg "volume" vs "instance", etc.

* `image_type` - Image type as referenced in cm platform.

* `virtualization_type` - Image virtualization type as referenced in cm platform. Eg "hvm" etc.

* `links` - Hrefs of related API resources.

* `href` - Href of the image.
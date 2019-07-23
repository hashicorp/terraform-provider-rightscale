---
layout: "rightscale"
page_title: "Rightscale: instance"
sidebar_current: "docs-rightscale-datasource-instance"
description: |-
  Defines an instance datasource to operate against.
---

# rightscale_instance

Use this data source to locate and extract info about an existing [instance](http://reference.rightscale.com/api1.5/resources/ResourceInstances.html) to pass to other rightscale resources.

Filter block is optional - ommitting it will result in the first available instance in a given cloud.

## Example Usage 1: Basic configuration of a instance data source

```hcl
data "rightscale_instance" "an_instance" {
  cloud_href = "${data.rightscale_cloud.ec2_us_oregon.href}"

  filter {
    name = "my_instance"
  }
}

output "instance name" {
  value = "${data.rightscale_instance.an_instance.name}"
}

output "instance href" {
  value = "${data.rightscale_instance.an_instance.href}"
}

data "rightscale_cloud" "ec2_us_oregon" {
  filter {
    name = "EC2 us-west-2"
    cloud_type = "amazon"
  }
}
```

## Argument Reference

The following arguments are supported:

* `cloud_href` (Required unless server_array_href specified) - The cloud_href the instance belongs to (mutually exclusive with server_array_href, specify only one)

* `server_array_href` (Required unless cloud_href specified) - The server_array_href the instance belongs to (mutually exclusive with cloud_href, specify only one)

* `filter` (Optional) - The filter block supports:

  * `name` - The name of the instance

  * `state` - The state of the instance (e.g.: operational, terminated, stranded, ...)

  * `os_platform` - The OS platform of the instance. One of "linux" or "windows"

  * `parent_href` - The Href of instance server or server array parent resource.

  * `server_template_href` - The Href of the instance server template resource

  * `public_dns_name` - The public DNS name of the instance

  * `private_dns_name` - The private DNS name of the instance

  * `public_ip` - The public IP of the instance

  * `private_ip` - The private IP of the instance

  * `resource_uid` - The resource_uid of the instance.  If this filter option is set, additional retry logic will be enabled to wait up to 5 minutes for cloud resources to be polled and populated for use.

  * `deployment_href` - The href of the [deployment](http://docs.rightscale.com/cm/dashboard/manage/deployments/) that contains the instance (e.g. /api/deployments/594684003)

  * `placement_group_href` - The href of the [placement_group](http://docs.rightscale.com/cm/dashboard/clouds/aws/ec2_placement_groups.html) that contains the instance (e.g. /api/placement_groups/512SV3FUJA7OO)

  * `datacenter_href` - The href of the [datacenter](http://docs.rightscale.com/cm/dashboard/clouds/generic/datacenter_zones_concepts.html) that holds the instance (e.g. /api/clouds/6/datacenters/6IHONC8ANOUHI)

## Attributes Reference

The following attributes are exported:

* `associate_public_ip_address` - Indicates if the instance will get a Public IP address

* `cloud_href` - The cloud_href the instance belongs to (mutually exclusive with server_array_href)

* `server_array_href` - The server_array_href the instance belongs to (mutually exclusive with cloud_href)

* `cloud_specific_attributes` - Attributes specific to the cloud the instance belongs to that have no specific rightscale abstraction.  This block includes:

  * `admin_username` - The user that will be granted administrative privileges. Supported by AzureRM cloud only.

  * `automatic_instance_store_mapping` - A flag indicating whether instance store mapping should be enabled.  Only available on clouds supporting automatic instance store mapping.

  * `availability_set` - Availability set for raw instance. Supported by Azure v2 cloud only.

  * `create_boot_volume` - If enabled, the instance will launch into volume storage. Otherwise, it will boot to local storage.  Only available on clouds supporting this option.

  * `create_default_port_forwarding_rules` - Automatically create default port forwarding rules (enabled by default). Supported by Azure cloud only.

  * `delete_boot_volume` - If enabled, the associated volume will be deleted when the instance is terminated.  Only available on clouds supporting this option.

  * `disk_gb` - The size of root disk. Supported by UCA cloud only.

  * `ebs_optimized` - Whether the instance is able to connect to IOPS-enabled volumes.  AWS clouds only.

  * `iam_instance_profile` - The name or ARN of the IAM Instance Profile (IIP) to associate with the instance. AWS clouds only.

  * `keep_alive_id` - The id of keep alive. Supported by UCA cloud only.

  * `local_ssd_count` - Additional local SSDs. Supported by GCE cloud only.

  * `local_ssd_interface` - The type of SSD(s) to be created. Supported by GCE cloud only.

  * `max_spot_price` - Specify the max spot price you will pay for. Required when 'pricing_type' is 'spot'. Only applies to clouds which support spot-pricing and when 'spot' is chosen as the 'pricing_type'. Should be a Float value >= 0.001, eg: 0.095, 0.123, 1.23, etc... AWS clouds only.

  * `memory_mb` - The size of instance memory. Supported by UCA cloud only.

  * `metadata` - Extra data used for configuration, in query string format. AWS clouds only.

  * `num_cores` - The number of instance cores. Supported by UCA cloud only.

  * `placement_tenancy` - The tenancy of the server you want to launch. A server with a tenancy of dedicated runs on single-tenant hardware and can only be launched into a VPC.  AWS clouds only.

  * `preemptible` - Launch a preemptible instance. A preemptible instance costs much less, but lasts only 24 hours. It can be terminated sooner due to system demands. Supported by GCE cloud only.

  * `pricing_type` - Specify whether or not you want to utilize 'fixed' (on-demand) or 'spot' pricing. Defaults to 'fixed' and only applies to clouds which support spot instances. Can only be set on when creating a new Instance, Server, or ServerArray, or when updating a Server or ServerArray's next_instance.  AWS clouds only.

  * `root_volume_performance` - The number of IOPS (I/O Operations Per Second) this root volume should support. Only available on clouds supporting performance provisioning.

  * `root_volume_size` - The size for root disk. Only available on clouds supporting dynamic resizing of root volume size.

  * `root_volume_type_uid` - The type of root volume for instance. Only available on clouds supporting root volume type.

  * `service_account` - Email of service account for instance. Scope will default to cloud-platform. Supported by GCE cloud only.

* `name` - The name of the instance

* `pricing_type` - Pricing type of the instance (e.g. fixed, spot)

* `resource_uid` - The resource_uid of the instance (e.g. e0bf62bc-4e35-11e8-9f1f-0242ac110003)

* `links` - Hrefs of related API resources

* `locked` - Whether instance is locked, a locked instance cannot be terminated or deleted

* `private_ip_addresses` - List of private IP addresses of the instance

* `public_ip_addresses` - List of public IP addresses of the instance

* `state` - The instance state (e.g. operational, terminated, stranded, ...)

* `created_ at` - Time of creation of the instance

* `updated_at` - Last update of the instance

* `id` - The instance ID (e.g. rs_cm:/api/clouds/1/instances/63NFHKF8B7RP4)

* `href` - Href of the instance (e.g. /api/clouds/1/instances/63NFHKF8B7RP4)

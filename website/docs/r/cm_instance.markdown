---
layout: "rightscale"
page_title: "Rightscale: instance"
sidebar_current: "docs-rightscale-resource-instance"
description: |-
  Create and maintain a RightScale instance.
---

# rightscale_instance

Use this resource to create, update or destroy RightScale [instances](http://reference.rightscale.com/api1.5/resources/ResourceInstances.html).

## Example Usage : Basic configuration of an instance resource

```hcl
resource "rightscale_instance" "an_instance" {
  cloud_href = "/api/clouds/6"
  image_href = "/api/clouds/6/images/3TRNL47PJB97N"
  instance_type_href = "/api/clouds/6/instance_types/8SCHNH0JBHE1R"
  deployment_href = "/api/deployments/934588004"
  name = "My Instance"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the instance.

* `cloud_href` - (Required) The cloud_href the instance belongs to.

* `image_href` - (Required) The href of the instance image.

* `instance_type_href` - (Required) The href of the instance type.

* `server_template_href` - (Optional) The href of the instance server template resource.

* `inputs` - (Optional) Inputs associated with an instance when incarnated from a [server](https://github.com/terraform-providers/terraform-provider-rightscale/blob/master/website/docs/r/cm_server.markdown) or [server_array](https://github.com/terraform-providers/terraform-provider-rightscale/blob/master/website/docs/r/cm_server_array.markdown).

* `associate_public_ip_address` - (Optional) Indicates if the instance will get a Public IP address.

* `datacenter_href` - (Optional) The href of the datacenter that holds the instance (e.g. /api/clouds/6/datacenters/6IHONC8ANOUHI).

* `deployment_href` - (Optional) The href of the deployment that contains the instance (e.g. /api/deployments/594684003).

* `ip_forwarding_enabled` - (Optional) Allows this Instance to send and receive network traffic when the source and destination IP addresses do not match the IP address of this Instance.

* `private_ip_address` - (Optional) The private ip address of this instance.

* `kernel_image_href` - (Optional) The href of the instance kernel image.

* `ramdisk_image_href` - (Optional) The href of the instance ramdisk image.

* `security_group_hrefs` - (Optional) The href of the instance security groups.

* `placement_group_href` - (Optional) The href of the [placement_group](http://docs.rightscale.com/cm/dashboard/clouds/aws/ec2_placement_groups.html) that contains the instance (e.g. /api/placement_groups/512SV3FUJA7OO).

* `ssh_key_href` - (Optional) The href of the SSH key to use.

* `subnet_hrefs` - (Optional) The hrefs of the instance subnet.

* `user_data` - (Optional) User data that RightScale automatically passes to your instance at boot time.

* `locked` - (Optional)  Whether instance is locked, a locked instance cannot be terminated or deleted.

* `cloud_specific_attributes` - (Optional) Attributes specific to the cloud the instance belongs to that have no specific rightscale abstraction.  This block supports:

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

  * `metadata"` - Extra data used for configuration, in query string format. AWS clouds only.

  * `num_cores` - The number of instance cores. Supported by UCA cloud only.

  * `placement_tenancy` - The tenancy of the server you want to launch. A server with a tenancy of dedicated runs on single-tenant hardware and can only be launched into a VPC.  AWS clouds only.

  * `preemptible` - Launch a preemptible instance. A preemptible instance costs much less, but lasts only 24 hours. It can be terminated sooner due to system demands. Supported by GCE cloud only.

  * `pricing_type` - Specify whether or not you want to utilize 'fixed' (on-demand) or 'spot' pricing. Defaults to 'fixed' and only applies to clouds which support spot instances. Can only be set on when creating a new Instance, Server, or ServerArray, or when updating a Server or ServerArray's next_instance.  AWS clouds only.

  * `root_volume_performance` - The number of IOPS (I/O Operations Per Second) this root volume should support. Only available on clouds supporting performance provisioning.

  * `root_volume_size` - The size for root disk. Only available on clouds supporting dynamic resizing of root volume size.

  * `root_volume_type_uid` - The type of root volume for instance. Only available on clouds supporting root volume type.

  * `service_account` - Email of service account for instance. Scope will default to cloud-platform. Supported by GCE cloud only.

## Attributes Reference

The following attributes are exported:

* `links` - Hrefs of related API resources

* `created_at` - Datestamp of instance creation.

* `updated_at` - Datestamp of when instance was updated last.

* `state` - The state of the instance (operational, terminating, pending, stranded, etc.)

* `href` - Href of the instance.

* `resource_uid` - Cloud resource_uid as reported by cm platform.

* `public_ip_addresses` - List of public IP addresses associated to the instance

* `private_ip_addresses` - List of private IP addresses associated to the instance

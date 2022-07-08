---
subcategory: "RDS"
layout: "alicloud"
page_title: "Alicloud: alicloud_rds_error_logs"
sidebar_current: "docs-alicloud-datasource-rds-error_logs"
description: |-
  Provides a list of Rds Error Logs to the user.
---

# alicloud\_rds\_error\_logs

This data source provides the Rds Error Logs of the current Alibaba Cloud user.

-> **NOTE:** Available in v1.175.0+.

## Example Usage

Basic Usage

```terraform
data "alicloud_rds_error_logs" "example" {
  db_instance_id = "example_value"
  start_time     = "2022-06-04T13:56Z"
  end_time       = "2022-06-08T13:56Z"
}
```

## Argument Reference

The following arguments are supported:

* `db_instance_id` - (Required, ForceNew) The db instance id.
* `end_time` - (Required, ForceNew) The end time.
* `output_file` - (Optional) File name where to save data source results (after running `terraform plan`).
* `start_time` - (Required, ForceNew) The start time.

## Argument Reference

The following attributes are exported in addition to the arguments listed above:

* `error_logs` - A list of Rds Error Logs. Each element contains the following attributes:
  * `error_info` - The error log information.
  * `create_time` - The time when the error log was generated. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
  
  
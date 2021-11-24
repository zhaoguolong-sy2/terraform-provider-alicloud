package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func init() {
	resource.AddTestSweepers("alicloud_rds_clone_db_instance", &resource.Sweeper{
		Name: "alicloud_rds_clone_db_instance",
		F:    testSweepDBInstances,
	})
}

func resourceCloneDBInstanceConfigDependence(name string) string {
	return fmt.Sprintf(`
variable "name" {
	default = "%s"
}
`, name)
}

var cloneInstanceBasicMap = map[string]string{
	"engine_version":       "13.0",
	"instance_type":        CHECKSET,
	"instance_storage":     CHECKSET,
	"zone_id":              CHECKSET,
	"instance_charge_type": "Postpaid",
	"connection_string":    CHECKSET,
	"port":                 CHECKSET,
}

func TestAccAlicloudRdsCloneDBInstancePostgreSQLSSL(t *testing.T) {
	var instance map[string]interface{}
	//var ips []map[string]interface{}

	resourceId := "alicloud_rds_clone_db_instance.default"
	ra := resourceAttrInit(resourceId, cloneInstanceBasicMap)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &instance, func() interface{} {
		return &RdsService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeDBInstance")
	rac := resourceAttrCheckInit(rc, ra)

	testAccCheck := rac.resourceAttrMapUpdateSet()
	name := "tf-testAccDBInstanceConfig"
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, resourceCloneDBInstanceConfigDependence)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},

		// module name
		IDRefreshName: resourceId,

		Providers:    testAccProviders,
		CheckDestroy: rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"source_db_instance_id":    "pgm-bp1t0u9s4861tcw3",
					"db_instance_storage_type": "cloud_ssd",
					"instance_charge_type":     "Postpaid",
					"restore_time":             "2021-11-18T10:00:00Z",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"source_db_instance_id":    "pgm-bp1t0u9s4861tcw3",
						"db_instance_storage_type": "cloud_ssd",
						"instance_charge_type":     "Postpaid",
						"restore_time":             "2021-11-18T10:00:00Z",
						"engine_version":           "13.0",
						"instance_type":            CHECKSET,
						"instance_storage":         CHECKSET,
						"zone_id":                  CHECKSET,
						"connection_string":        CHECKSET,
						"port":                     CHECKSET,
					}),
				),
			},
			{
				ResourceName:            resourceId,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"force_restart"},
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"instance_name": "tf-testAccDBInstance_instance_name",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"instance_name": "tf-testAccDBInstance_instance_name",
					}),
				),
			},
		},
	})
}

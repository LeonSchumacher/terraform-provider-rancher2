package rancher2

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func machineConfigV2NutanixFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"cloudInit": {
			Type:     schema.TypeString,
			Required: true,
		},
		"cluster": {
			Type:     schema.TypeString,
			Required: true,
		},
		"diskSize": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"project": {
			Type:     schema.TypeString,
			Required: true,
		},
		"storageContainer": {
			Type:     schema.TypeString,
			Required: true,
		},
		"vmCategories": {
			Type:     schema.TypeList,
			Required: false,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"vmCores": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"vmImage": {
			Type:     schema.TypeString,
			Required: true,
		},
		"vmImageSize": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"vmMemory": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"vmNetwork": {
			Type:     schema.TypeList,
			Required: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"vmSerialPort": {
			Type:     schema.TypeBool,
			Required: true,
		},
	}

	return s
}

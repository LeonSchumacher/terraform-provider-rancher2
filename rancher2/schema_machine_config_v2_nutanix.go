package rancher2

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func machineConfigV2NutanixFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"cloud_init": {
			Type:     schema.TypeString,
			Required: true,
		},
		"cluster": {
			Type:     schema.TypeString,
			Required: true,
		},
		"disk_size": {
			Type:     schema.TypeString,
			Required: true,
		},
		"project": {
			Type:     schema.TypeString,
			Required: true,
		},
		"storage_container": {
			Type:     schema.TypeString,
			Required: true,
		},
		"vm_categories": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"vm_cores": {
			Type:     schema.TypeString,
			Required: true,
		},
		"vm_cpus": {
			Type:     schema.TypeString,
			Required: true,
		},
		"vm_cpu_passthrough": {
			Type:     schema.TypeBool,
			Required: true,
		},
		"vm_image": {
			Type:     schema.TypeString,
			Required: true,
		},
		"vm_image_size": {
			Type:     schema.TypeString,
			Required: true,
		},
		"vm_memory": {
			Type:     schema.TypeString,
			Required: true,
		},
		"vm_network": {
			Type:     schema.TypeList,
			Required: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"vm_serial_port": {
			Type:     schema.TypeBool,
			Required: true,
		},
	}

	return s
}

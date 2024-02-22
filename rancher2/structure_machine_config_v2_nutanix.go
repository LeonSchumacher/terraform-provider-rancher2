package rancher2

import (
	norman "github.com/rancher/norman/types"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	machineConfigV2NutanixKind         = "NutanixConfig"
	machineConfigV2NutanixAPIVersion   = "rke-machine-config.cattle.io/v1"
	machineConfigV2NutanixAPIType      = "rke-machine-config.cattle.io.nutanixconfig"
	machineConfigV2NutanixClusterIDsep = "."
)

//Types

type machineConfigV2Nutanix struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	CloudInit         string   `json:"cloudInit,omitempty" yaml:"cloudInit,omitempty"`
	Cluster           string   `json:"cluster,omitempty" yaml:"cluster,omitempty"`
	DiskSize          string   `json:"diskSize,omitempty" yaml:"diskSize,omitempty"`
	Project           string   `json:"project,omitempty" yaml:"project,omitempty"`
	StorageContainer  string   `json:"storageContainer,omitempty" yaml:"storageContainer,omitempty"`
	VmCategories      []string `json:"vmCategories,omitempty" yaml:"vmCategories,omitempty"`
	VmCores           string   `json:"vmCores,omitempty" yaml:"vmCores,omitempty"`
	VmCPUs            string   `json:"vmCpus,omitempty" yaml:"vmCpus,omitempty"`
	VmCPUPassthrough  bool     `json:"vmCpuPassthrough,omitempty" yaml:"vmCpuPassthrough,omitempty"`
	VmImage           string   `json:"vmImage,omitempty" yaml:"vmImage,omitempty"`
	VmImageSize       string   `json:"vmImageSize,omitempty" yaml:"vmImageSize,omitempty"`
	VmMemory          string   `json:"vmMem,omitempty" yaml:"vmMem,omitempty"`
	VmNetwork         []string `json:"vmNetwork,omitempty" yaml:"vmNetwork,omitempty"`
	VmSerialPort      bool     `json:"vmSerialPort,omitempty" yaml:"vmSerialPort,omitempty"`
}

type MachineConfigV2Nutanix struct {
	norman.Resource
	machineConfigV2Nutanix
}

// Flatteners

func flattenMachineConfigV2Nutanix(in *MachineConfigV2Nutanix) []interface{} {
	if in == nil {
		return nil
	}

	obj := make(map[string]interface{})

	obj["cloud_init"] = in.CloudInit
	obj["cluster"] = in.Cluster
	obj["disk_size"] = in.DiskSize
	obj["project"] = in.Project
	obj["storage_container"] = in.StorageContainer
	obj["vm_categories"] = in.VmCategories
	obj["vm_cores"] = in.VmCores
	obj["vm_cpus"] = in.VmCPUs
	obj["vm_cpu_passthrough"] = in.VmCPUPassthrough
	obj["vm_image"] = in.VmImage
	obj["vm_image_size"] = in.VmImageSize
	obj["vm_memory"] = in.VmMemory
	obj["vm_network"] = in.VmNetwork
	obj["vm_serial_port"] = in.VmSerialPort

	return []interface{}{obj}
}

// Expanders

func expandMachineConfigV2Nutanix(p []interface{}, source *MachineConfigV2) *MachineConfigV2Nutanix {
	if p == nil || len(p) == 0 || p[0] == nil {
		return nil
	}
	obj := &MachineConfigV2Nutanix{}

	if len(source.ID) > 0 {
		obj.ID = source.ID
	}
	in := p[0].(map[string]interface{})

	obj.TypeMeta.Kind = machineConfigV2NutanixKind
	obj.TypeMeta.APIVersion = machineConfigV2NutanixAPIVersion
	source.TypeMeta = obj.TypeMeta
	obj.ObjectMeta = source.ObjectMeta

	if v, ok := in["cloud_init"].(string); ok && len(v) > 0 {
		obj.CloudInit = v
	}
	if v, ok := in["cluster"].(string); ok && len(v) > 0 {
		obj.Cluster = v
	}
	if v, ok := in["disk_size"].(string); ok {
		obj.DiskSize = v
	}
	if v, ok := in["project"].(string); ok && len(v) > 0 {
		obj.Project = v
	}
	if v, ok := in["storage_container"].(string); ok && len(v) > 0 {
		obj.StorageContainer = v
	}
	if v, ok := in["vm_categories"].([]interface{}); ok && len(v) > 0 {
		obj.VmCategories = toArrayString(v)
	}
	if v, ok := in["vm_cores"].(string); ok {
		obj.VmCores = v
	}
	if v, ok := in["vm_cpus"].(string); ok {
		obj.VmCPUs = v
	}
	if v, ok := in["vm_cpu_passthrough"].(bool); ok {
		obj.VmCPUPassthrough = v
	}
	if v, ok := in["vm_image"].(string); ok && len(v) > 0 {
		obj.VmImage = v
	}
	if v, ok := in["vm_image_size"].(string); ok {
		obj.VmImageSize = v
	}
	if v, ok := in["vm_memory"].(string); ok {
		obj.VmMemory = v
	}
	if v, ok := in["vm_network"].([]interface{}); ok && len(v) > 0 {
		obj.VmNetwork = toArrayString(v)
	}
	if v, ok := in["vm_serial_port"].(bool); ok {
		obj.VmSerialPort = v
	}

	return obj
}

package rancher2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// Types

type nutanixCredentialConfig struct {
	Password string `json:"password,omitempty" yaml:"password,omitempty"`
	Username string `json:"username,omitempty" yaml:"username,omitempty"`
	Endpoint string `json:"endpoint,omitempty" yaml:"endpoint,omitempty"`
	Port     string `json:"nutanixPort,omitempty" yaml:"nutanixPort,omitempty"`
}

// Flatteners

func flattenCloudCredentialNutanix(in *nutanixCredentialConfig, p []interface{}) []interface{} {
	var obj map[string]interface{}
	if len(p) == 0 || p[0] == nil {
		obj = make(map[string]interface{})
	} else {
		obj = p[0].(map[string]interface{})
	}

	if in == nil {
		return []interface{}{}
	}

	if len(in.Password) > 0 {
		obj["password"] = in.Password
	}
	if len(in.Username) > 0 {
		obj["username"] = in.Username
	}
	if len(in.Endpoint) > 0 {
		obj["endpoint"] = in.Endpoint
	}
	if len(in.Port) > 0 {
		obj["nutanix_port"] = in.Port
	}

	return []interface{}{obj}
}

// Expanders

func expandCloudCredentialNutanix(p []interface{}) *nutanixCredentialConfig {
	obj := &nutanixCredentialConfig{}
	if len(p) == 0 || p[0] == nil {
		return obj
	}
	in := p[0].(map[string]interface{})

	if v, ok := in["password"].(string); ok && len(v) > 0 {
		obj.Password = v
	}
	if v, ok := in["username"].(string); ok && len(v) > 0 {
		obj.Username = v
	}
	if v, ok := in["endpoint"].(string); ok && len(v) > 0 {
		obj.Endpoint = v
	}
	if v, ok := in["nutanix_port"].(string); ok && len(v) > 0 {
		obj.Port = v
	}

	return obj
}

// Schemas

func cloudCredentialNutanixFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"password": {
			Type:        schema.TypeString,
			Required:    true,
			Sensitive:   true,
			Description: "Nutanix password",
		},
		"username": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Nutanix username",
		},
		"endpoint": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Nutanix IP/hostname for Prism Center",
		},
		"nutanix_port": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "443",
			Description: "Nutanix Port for Prism Center",
		},
	}

	return s
}


package rancher2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// Types

type nutnaixCredentialConfig struct {
	Password string `json:"password,omitempty" yaml:"password,omitempty"`
	Username string `json:"username,omitempty" yaml:"username,omitempty"`
	Endpoint string `json:"endpoint,omitempty" yaml:"endpoint,omitempty"`
	Port     string `json:"nutanixPort,omitempty" yaml:"nutnaixPort,omitempty"`
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
			Description: "Nutnaix Port for Prism Center",
		},
	}

	return s
}

package rancher2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

const (
	nutanixConfigDriver = "nutanix"
)

//Types

type nutanixConfig struct {
	Cloudinit        string   `json:"cloudinit,omitempty" yaml:"cloudinit,omitempty"`
	Password         string   `json:"password,omitempty" yaml:"password,omitempty"`
	Username         string   `json:"username,omitempty" yaml:"username,omitempty"`
	Insecure         bool     `json:"insecure,omitempty" yaml:"insecure,omitempty"`
	Cluster          string   `json:"cluster,omitempty" yaml:"cluster,omitempty"`
	DiskSize         string   `json:"diskSize,omitempty" yaml:"diskSize,omitempty"`
	Endpoint         string   `json:"endpoint,omitempty" yaml:"endpoint,omitempty"`
	Port             string   `json:"port,omitempty" yaml:"port,omitempty"`
	Project          string   `json:"project,omitempty" yaml:"project,omitempty"`
	StorageContainer string   `json:"storageContainer,omitempty" yaml:"storageContainer,omitempty"`
	VmCategories     []string `json:"vmCategories,omitempty" yaml:"vmCategories,omitempty"`
	VmCores          string   `json:"vmCores,omitempty" yaml:"vmCores,omitempty"`
	VmCpuPassthrough bool     `json:"vmCpuPassthrough,omitempty" yaml:"vmCpuPassthrough,omitempty"`
	VmCpus           string   `json:"vmCpus,omitempty" yaml:"vmCpus,omitempty"`
	VmImage          string   `json:"vmImage,omitempty" yaml:"vmImage,omitempty"`
	VmImageSize      string   `json:"vmImageSize,omitempty" yaml:"vmImageSize,omitempty"`
	VmMem            string   `json:"vmMem,omitempty" yaml:"vmMem,omitempty"`
	VmNetwork        []string `json:"vmNetwork,omitempty" yaml:"vmNetwork,omitempty"`
	VmSerialPort     bool     `json:"vmSerialPort,omitempty" yaml:"vmSerialPort,omitempty"`
}

//Schemas

func nutanixConfigFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"cloudinit": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Cloud-init configuration passed to the Host",
		},
		"password": {
			Type:        schema.TypeString,
			Optional:    true,
			Sensitive:   true,
			Description: "Password for the Nutanix Cluster",
		},
		"username": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Username for the Nutanix Cluster",
		},
		"insecure": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: "Whether to skip TLS validation for the Nutanix Cluster",
		},
		"cluster": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Name of the Nutanix Cluster",
		},
		"diskSize": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Disk Size for the primary disk of the Virtual Machine",
		},
		"endpoint": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "FQDN of the Nutnaix Prism Endpoint",
		},
		"port": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Network port of the Nutanix Prism Endpoint",
		},
		"project": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Name of the Nutanix Project",
		},
		"storageContainer": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The storage container UUID of the additional disk to add to the VM",
		},
		"vmCategories": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "The name of the categories who will be applied to the newly created VM",
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"vmCores": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The number of cores to assign to the newly created VM",
		},
		"vmCpuPassthrough": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Enable passthrough the host's CPU features to the newly created VM",
		},
		"vmCpus": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The number of cpus in the newly created VM (core)",
		},
		"vmImage": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The name of the image to use for the newly created VM",
		},
		"vmImageSize": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Image will be inflated to this size",
		},
		"vmMem": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Memory size for the newly created VM",
		},
		"vmNetwork": {
			Type:        schema.TypeList,
			Optional:    false,
			Description: "The network(s) to which the VM is attached to",
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"vmSerialPort": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Attach a serial port to the newly created VM",
		},
	}
	return s
}

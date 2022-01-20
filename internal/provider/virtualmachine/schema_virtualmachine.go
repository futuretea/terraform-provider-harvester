package virtualmachine

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/harvester/terraform-provider-harvester/internal/util"
	"github.com/harvester/terraform-provider-harvester/pkg/constants"
)

func ResourceSchema() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		constants.FieldVirtualMachineMachineType: {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		constants.FieldVirtualMachineHostname: {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		constants.FieldVirtualMachineStart: {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		constants.FieldVirtualMachineCPU: {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  1,
		},
		constants.FieldVirtualMachineMemory: {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "1Gi",
		},
		constants.FieldVirtualMachineSSHKeys: {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		constants.FieldVirtualMachineCloudInit: {
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: resourceCloudInitSchema(),
			},
		},
		constants.FieldVirtualMachineDisk: {
			Type:     schema.TypeList,
			Required: true,
			MinItems: 1,
			Elem: &schema.Resource{
				Schema: resourceDiskSchema(),
			},
		},
		constants.FieldVirtualMachineNetworkInterface: {
			Type:     schema.TypeList,
			Required: true,
			MinItems: 1,
			Elem: &schema.Resource{
				Schema: resourceNetworkInterfaceSchema(),
			},
		},
		constants.FieldVirtualMachineInstanceNodeName: {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
	util.ResourceNamespacedSchemaWrap(s, false)
	return s
}

func DataSourceSchema() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		constants.FieldVirtualMachineMachineType: {
			Type:     schema.TypeString,
			Computed: true,
		},
		constants.FieldVirtualMachineHostname: {
			Type:     schema.TypeString,
			Computed: true,
		},
		constants.FieldVirtualMachineStart: {
			Type:     schema.TypeBool,
			Computed: true,
		},
		constants.FieldVirtualMachineCPU: {
			Type:     schema.TypeInt,
			Computed: true,
		},
		constants.FieldVirtualMachineMemory: {
			Type:     schema.TypeString,
			Computed: true,
		},
		constants.FieldVirtualMachineSSHKeys: {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		constants.FieldVirtualMachineCloudInit: {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: dataSourceCloudInitSchema(),
			},
		},
		constants.FieldVirtualMachineDisk: {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: dataSourceDiskSchema(),
			},
		},
		constants.FieldVirtualMachineNetworkInterface: {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: dataSourceNetworkInterfaceSchema(),
			},
		},
		constants.FieldVirtualMachineInstanceNodeName: {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
	util.DataSourceNamespacedSchemaWrap(s, false)
	return s
}

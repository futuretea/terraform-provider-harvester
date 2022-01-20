package clusternetwork

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/harvester/terraform-provider-harvester/internal/util"
	"github.com/harvester/terraform-provider-harvester/pkg/constants"
)

func ResourceSchema() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		constants.FieldClusterNetworkEnable: {
			Type:     schema.TypeBool,
			Required: true,
		},
		constants.FieldClusterNetworkDefaultPhysicalNIC: {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
	util.ResourceNonNamespacedSchemaWrap(s)
	return s
}

func DataSourceSchema() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		constants.FieldClusterNetworkEnable: {
			Type:     schema.TypeBool,
			Computed: true,
		},
		constants.FieldClusterNetworkDefaultPhysicalNIC: {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
	util.DataSourceNonNamespacedSchemaWrap(s)
	return s
}

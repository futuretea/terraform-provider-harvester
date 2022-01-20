package network

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/harvester/terraform-provider-harvester/internal/util"
	"github.com/harvester/terraform-provider-harvester/pkg/constants"
)

func ResourceSchema() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		constants.FieldNetworkVlanID: {
			Type:         schema.TypeInt,
			Required:     true,
			ValidateFunc: validation.IntBetween(1, 4094),
			Description:  "eg.1-4094",
		},
		constants.FieldNetworkConfig: {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
	util.ResourceNamespacedSchemaWrap(s, false)
	return s
}

func DataSourceSchema() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		constants.FieldNetworkVlanID: {
			Type:     schema.TypeInt,
			Computed: true,
		},
		constants.FieldNetworkConfig: {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
	util.DataSourceNamespacedSchemaWrap(s, false)
	return s
}

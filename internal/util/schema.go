package util

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/apimachinery/pkg/util/validation"

	"github.com/harvester/terraform-provider-harvester/pkg/constants"
)

func ResourceNamespacedSchemaWrap(s map[string]*schema.Schema, system bool) {
	var namespace = constants.NamespaceDefault
	if system {
		namespace = constants.NamespaceHarvesterSystem
	}
	ResourceNonNamespacedSchemaWrap(s)
	s[constants.FieldCommonNamespace] = &schema.Schema{
		Type:         schema.TypeString,
		ForceNew:     true,
		Optional:     true,
		Default:      namespace,
		ValidateFunc: IsValidName,
	}
}

func ResourceNonNamespacedSchemaWrap(s map[string]*schema.Schema) {
	s[constants.FieldCommonName] = &schema.Schema{
		Type:         schema.TypeString,
		Required:     true,
		ForceNew:     true,
		ValidateFunc: IsValidName,
		Description:  "A unique name",
	}
	s[constants.FieldCommonDescription] = &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Any text you want that better describes this resource",
	}
	s[constants.FieldCommonTags] = &schema.Schema{
		Type:     schema.TypeMap,
		Optional: true,
	}
	s[constants.FieldCommonState] = &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	}
}

func DataSourceNamespacedSchemaWrap(s map[string]*schema.Schema, system bool) {
	var namespace = constants.NamespaceDefault
	if system {
		namespace = constants.NamespaceHarvesterSystem
	}
	DataSourceNonNamespacedSchemaWrap(s)
	s[constants.FieldCommonNamespace] = &schema.Schema{
		Type:         schema.TypeString,
		Optional:     true,
		Default:      namespace,
		ValidateFunc: IsValidName,
	}
}

func DataSourceNonNamespacedSchemaWrap(s map[string]*schema.Schema) {
	s[constants.FieldCommonName] = &schema.Schema{
		Type:         schema.TypeString,
		Required:     true,
		ValidateFunc: IsValidName,
		Description:  "A unique name",
	}
	s[constants.FieldCommonDescription] = &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	}
	s[constants.FieldCommonTags] = &schema.Schema{
		Type:     schema.TypeMap,
		Computed: true,
	}
	s[constants.FieldCommonState] = &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	}
}

func IsValidName(i interface{}, k string) ([]string, []error) {
	v, ok := i.(string)
	if !ok {
		return nil, []error{fmt.Errorf("expected type of %q to be string", k)}
	}

	if errs := validation.IsDNS1123Subdomain(v); len(errs) > 0 {
		return nil, []error{fmt.Errorf("expected %q to not be an kubernetes valid name", k)}
	}

	return nil, nil
}

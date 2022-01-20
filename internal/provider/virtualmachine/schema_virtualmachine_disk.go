package virtualmachine

import (
	"github.com/harvester/harvester/pkg/builder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/harvester/terraform-provider-harvester/internal/util"
	"github.com/harvester/terraform-provider-harvester/pkg/constants"
)

func resourceDiskSchema() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		constants.FieldDiskName: {
			Type:     schema.TypeString,
			Required: true,
		},
		constants.FieldDiskType: {
			Type:     schema.TypeString,
			Optional: true,
			Default:  builder.DiskTypeDisk,
			ValidateFunc: validation.StringInSlice([]string{
				builder.DiskTypeDisk,
				builder.DiskTypeCDRom,
			}, false),
		},
		constants.FieldDiskSize: {
			Type:     schema.TypeString,
			Optional: true,
		},
		constants.FieldDiskBus: {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
			ValidateFunc: validation.StringInSlice([]string{
				builder.DiskBusVirtio,
				builder.DiskBusSata,
				builder.DiskBusScsi,
				"",
			}, false),
		},
		constants.FieldDiskBootOrder: {
			Type:         schema.TypeInt,
			Optional:     true,
			Default:      0,
			ValidateFunc: validation.IntAtLeast(0),
		},
		constants.FieldVolumeImage: {
			Type:     schema.TypeString,
			Optional: true,
		},
		constants.FieldDiskExistingVolumeName: {
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: util.IsValidName,
		},
		constants.FieldDiskContainerImageName: {
			Type:     schema.TypeString,
			Optional: true,
		},
		constants.FieldDiskAutoDelete: {
			Type:     schema.TypeBool,
			Optional: true,
			Computed: true,
		},
		constants.FieldDiskHotPlug: {
			Type:     schema.TypeBool,
			Optional: true,
			Computed: true,
		},
		constants.FieldVolumeStorageClassName: {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			ValidateFunc: util.IsValidName,
		},
		constants.FieldVolumeMode: {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
			ValidateFunc: validation.StringInSlice([]string{
				builder.PersistentVolumeModeBlock,
				builder.PersistentVolumeModeFilesystem,
			}, false),
		},
		constants.FieldVolumeAccessMode: {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
			ValidateFunc: validation.StringInSlice([]string{
				builder.PersistentVolumeAccessModeReadWriteOnce,
				builder.PersistentVolumeAccessModeReadOnlyMany,
				builder.PersistentVolumeAccessModeReadWriteMany,
			}, false),
		},
		constants.FieldDiskVolumeName: {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			ValidateFunc: util.IsValidName,
		},
	}
	return s
}

func dataSourceDiskSchema() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		constants.FieldDiskName: {
			Type:     schema.TypeString,
			Computed: true,
		},
		constants.FieldDiskType: {
			Type:     schema.TypeString,
			Computed: true,
		},
		constants.FieldDiskSize: {
			Type:     schema.TypeString,
			Computed: true,
		},
		constants.FieldDiskBus: {
			Type:     schema.TypeString,
			Computed: true,
		},
		constants.FieldDiskBootOrder: {
			Type:     schema.TypeInt,
			Computed: true,
		},
		constants.FieldVolumeImage: {
			Type:     schema.TypeString,
			Computed: true,
		},
		constants.FieldDiskExistingVolumeName: {
			Type:     schema.TypeString,
			Computed: true,
		},
		constants.FieldDiskContainerImageName: {
			Type:     schema.TypeString,
			Computed: true,
		},
		constants.FieldDiskAutoDelete: {
			Type:     schema.TypeBool,
			Computed: true,
		},
		constants.FieldDiskHotPlug: {
			Type:     schema.TypeBool,
			Computed: true,
		},
		constants.FieldVolumeStorageClassName: {
			Type:     schema.TypeString,
			Computed: true,
		},
		constants.FieldVolumeMode: {
			Type:     schema.TypeString,
			Computed: true,
		},
		constants.FieldVolumeAccessMode: {
			Type:     schema.TypeString,
			Computed: true,
		},
		constants.FieldDiskVolumeName: {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
	return s
}

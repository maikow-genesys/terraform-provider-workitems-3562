package routing_utilization

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"strings"
	"terraform-provider-genesyscloud/genesyscloud/provider"
	resourceExporter "terraform-provider-genesyscloud/genesyscloud/resource_exporter"
	"time"
)

const resourceName = "genesyscloud_routing_utilization"

type MediaUtilization struct {
	MaximumCapacity         int32    `json:"maximumCapacity"`
	InterruptableMediaTypes []string `json:"interruptableMediaTypes"`
	IncludeNonAcd           bool     `json:"includeNonAcd"`
}

type LabelUtilization struct {
	MaximumCapacity      int32    `json:"maximumCapacity"`
	InterruptingLabelIds []string `json:"interruptingLabelIds"`
}

type OrgUtilizationWithLabels struct {
	Utilization       map[string]MediaUtilization `json:"utilization"`
	LabelUtilizations map[string]LabelUtilization `json:"labelUtilizations"`
}

var (
	// Map of SDK media type name to schema media type name
	UtilizationMediaTypes = map[string]string{
		"call":     "call",
		"callback": "callback",
		"chat":     "chat",
		"email":    "email",
		"message":  "message",
	}

	UtilizationSettingsResource = &schema.Resource{
		Schema: map[string]*schema.Schema{
			"maximum_capacity": {
				Description:  "Maximum capacity of conversations of this media type. Value must be between 0 and 25.",
				Type:         schema.TypeInt,
				Required:     true,
				ValidateFunc: validation.IntBetween(0, 25),
			},
			"interruptible_media_types": {
				Description: fmt.Sprintf("Set of other media types that can interrupt this media type (%s).", strings.Join(getSdkUtilizationTypes(), " | ")),
				Type:        schema.TypeSet,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"include_non_acd": {
				Description: "Block this media type when on a non-ACD conversation.",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
		},
	}

	UtilizationLabelResource = &schema.Resource{
		Schema: map[string]*schema.Schema{
			"label_id": {
				Description: "Id of the label being configured.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"maximum_capacity": {
				Description:  "Maximum capacity of conversations with this label. Value must be between 0 and 25.",
				Type:         schema.TypeInt,
				Required:     true,
				ValidateFunc: validation.IntBetween(0, 25),
			},
			"interrupting_label_ids": {
				Description: "Set of other labels that can interrupt this label.",
				Type:        schema.TypeSet,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
)

func ResourceRoutingUtilization() *schema.Resource {
	return &schema.Resource{
		Description: "Genesys Cloud Org-wide Routing Utilization Settings.",

		CreateContext: provider.CreateWithPooledClient(createRoutingUtilization),
		ReadContext:   provider.ReadWithPooledClient(readRoutingUtilization),
		UpdateContext: provider.UpdateWithPooledClient(updateRoutingUtilization),
		DeleteContext: provider.DeleteWithPooledClient(deleteRoutingUtilization),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		SchemaVersion: 1,
		Timeouts: &schema.ResourceTimeout{
			Update: schema.DefaultTimeout(8 * time.Minute),
			Read:   schema.DefaultTimeout(8 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"call": {
				Description: "Call media settings. If not set, this reverts to the default media type settings.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				Elem:        UtilizationSettingsResource,
			},
			"callback": {
				Description: "Callback media settings. If not set, this reverts to the default media type settings.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				Elem:        UtilizationSettingsResource,
			},
			"message": {
				Description: "Message media settings. If not set, this reverts to the default media type settings.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				Elem:        UtilizationSettingsResource,
			},
			"email": {
				Description: "Email media settings. If not set, this reverts to the default media type settings.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				Elem:        UtilizationSettingsResource,
			},
			"chat": {
				Description: "Chat media settings. If not set, this reverts to the default media type settings.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				Elem:        UtilizationSettingsResource,
			},
			"label_utilizations": {
				Description: "Label utilization settings. If not set, default label settings will be applied. This is in PREVIEW and should not be used unless the feature is available to your organization.",
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Elem:        UtilizationLabelResource,
			},
		},
	}
}

func RoutingUtilizationExporter() *resourceExporter.ResourceExporter {
	return &resourceExporter.ResourceExporter{
		GetResourcesFunc: provider.GetAllWithPooledClient(getAllRoutingUtilization),
		RefAttrs:         map[string]*resourceExporter.RefAttrSettings{}, // No references
		AllowZeroValues:  []string{"maximum_capacity"},
	}
}

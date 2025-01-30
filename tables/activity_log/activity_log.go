package activity_log

import (
	"time"

	"github.com/turbot/tailpipe-plugin-sdk/schema"
)

type ActivityLog struct {
	schema.CommonFields

	AuthorizationInfo    *ActivityLogAuthorization `json:"authorization_info,omitempty"`
	Caller               *string                   `json:"caller"`
	Category             *string                   `json:"category"`
	Claims               *map[string]*string       `json:"claims,omitempty" parquet:"type=JSON"`
	CorrelationID        *string                   `json:"correlation_id"`
	Description          *string                   `json:"description"`
	EventDataID          *string                   `json:"event_data_id"`
	EventName            *string                   `json:"event_name"`
	EventTimestamp       *time.Time                `json:"event_timestamp"`
	HttpRequest          *ActivityLogHttpRequest   `json:"http_request,omitempty"`
	ID                   *string                   `json:"id"`
	Level                *string                   `json:"level"`
	OperationID          *string                   `json:"operation_id"`
	OperationName        *string                   `json:"operation_name"`
	Properties           *map[string]*string       `json:"properties,omitempty" parquet:"type=JSON"`
	ResourceGroupName    *string                   `json:"resource_group_name"`
	ResourceID           *string                   `json:"resource_id"`
	ResourceProviderName *string                   `json:"resource_provider_name"`
	ResourceType         *string                   `json:"resource_type"`
	Status               *string                   `json:"status"`
	SubStatus            *string                   `json:"sub_status"`
	SubmissionTimestamp  *time.Time                `json:"submission_timestamp"`
	SubscriptionID       *string                   `json:"subscription_id"`
	TenantID             *string                   `json:"tenant_id"`
}

type ActivityLogAuthorization struct {
	Action *string `json:"action"`
	Scope  *string `json:"scope"`
	Role   *string `json:"role"`
}

type ActivityLogHttpRequest struct {
	ClientIpAddress *string `json:"client_ip_address,omitempty"`
	ClientRequestId *string `json:"client_request_id,omitempty"`
	Method          *string `json:"method,omitempty"`
	Uri             *string `json:"uri,omitempty"`
}

// Storage account log structure

type Identity struct {
	Authorization Authorization           `json:"authorization,omitempty"`
	Claims        *map[string]interface{} `json:"claims,omitempty"` // Dynamic structure
}

type Authorization struct {
	Scope    *string  `json:"scope,omitempty"`
	Action   *string  `json:"action,omitempty"`
	Evidence Evidence `json:"evidence,omitempty"`
}

type Evidence struct {
	Role                *string `json:"role,omitempty"`
	RoleAssignmentScope *string `json:"roleAssignmentScope,omitempty"`
	RoleAssignmentID    *string `json:"roleAssignmentId,omitempty"`
	RoleDefinitionID    *string `json:"roleDefinitionId,omitempty"`
	PrincipalID         *string `json:"principalId,omitempty"`
	PrincipalType       *string `json:"principalType,omitempty"`
}

type AzureStorageAccountLog struct {
	RoleLocation    *string                 `json:"RoleLocation,omitempty"`
	Stamp           *string                 `json:"Stamp,omitempty"`
	ReleaseVersion  *string                 `json:"ReleaseVersion,omitempty"`
	Time            *time.Time              `json:"time,omitempty"`
	ResourceID      *string                 `json:"resourceId,omitempty"`
	OperationName   *string                 `json:"operationName,omitempty"`
	Category        *string                 `json:"category,omitempty"`
	ResultType      *string                 `json:"resultType,omitempty"`
	ResultSignature *string                 `json:"resultSignature,omitempty"`
	CallerIPAddress *string                 `json:"callerIpAddress,omitempty"`
	CorrelationID   *string                 `json:"correlationId,omitempty"`
	Identity        *Identity               `json:"identity,omitempty"`
	Level           *string                 `json:"level,omitempty"`
	Properties      *map[string]interface{} `json:"properties,omitempty"` // Dynamic structure
	TenantID        *string                 `json:"tenantId,omitempty"`
}

func (a *ActivityLog) GetColumnDescriptions() map[string]string {
	return map[string]string{
		"authorization_info":     "Details about the authorization context for the event, including the actor, scope, and role assignments.",
		"caller":                 "The identity of the user, service principal, or managed identity that performed the operation.",
		"category":               "The category of the event, such as 'Administrative', 'Security', 'Policy', or 'Alert'.",
		"claims":                 "Claims related to the authentication context, including tokens and identity assertions.",
		"correlation_id":         "A unique identifier for correlating related events across Azure services.",
		"description":            "A textual description of the event, providing additional details.",
		"event_data_id":          "A unique identifier for the event data entry.",
		"event_name":             "The name of the operation or action that was performed.",
		"event_timestamp":        "The date and time when the event occurred, in ISO 8601 format.",
		"http_request":           "Details about the HTTP request associated with the event, if applicable.",
		"id":                     "A globally unique identifier for the event.",
		"level":                  "The severity level of the event (e.g., 'Informational', 'Warning', 'Error', 'Critical').",
		"operation_id":           "A unique identifier for tracking the operation associated with the event.",
		"operation_name":         "The full name of the operation that was performed, often in a 'provider/action' format.",
		"properties":             "Additional metadata and custom properties related to the event, in JSON format.",
		"resource_group_name":    "The name of the Azure resource group that the affected resource belongs to.",
		"resource_id":            "The full Azure Resource Manager (ARM) ID of the affected resource.",
		"resource_provider_name": "The name of the Azure resource provider that handled the operation.",
		"resource_type":          "The type of the affected resource, such as 'Microsoft.Compute/virtualMachines'.",
		"status":                 "The final status of the operation (e.g., 'Succeeded', 'Failed').",
		"sub_status":             "Additional details about the operation’s status, providing finer granularity.",
		"submission_timestamp":   "The date and time when the event was submitted to Azure Monitor, in ISO 8601 format.",
		"subscription_id":        "The Azure Subscription ID under which the event occurred.",
		"tenant_id":              "The Azure Active Directory tenant ID associated with the event.",

		// override table specific tp_* column descriptions
		"tp_akas":  "Resource IDs associated with the event.",
		"tp_index": "The Azure Subscription ID under which the event occurred.",
	}
}

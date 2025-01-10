package rows

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
	Authorization Authorization `json:"authorization,omitempty"`
	Claims        *map[string]interface{} `json:"claims,omitempty"` // Dynamic structure
}

type Authorization struct {
	Scope              *string      `json:"scope,omitempty"`
	Action             *string      `json:"action,omitempty"`
	Evidence           Evidence    `json:"evidence,omitempty"`
}

type Evidence struct {
	Role                  *string `json:"role,omitempty"`
	RoleAssignmentScope   *string `json:"roleAssignmentScope,omitempty"`
	RoleAssignmentID      *string `json:"roleAssignmentId,omitempty"`
	RoleDefinitionID      *string `json:"roleDefinitionId,omitempty"`
	PrincipalID           *string `json:"principalId,omitempty"`
	PrincipalType         *string `json:"principalType,omitempty"`
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
	DurationMs      *string                 `json:"durationMs,omitempty"`
	CallerIPAddress *string                 `json:"callerIpAddress,omitempty"`
	CorrelationID   *string                 `json:"correlationId,omitempty"`
	Identity        *Identity               `json:"identity,omitempty"`
	Level           *string                 `json:"level,omitempty"`
	Properties      *map[string]interface{} `json:"properties,omitempty"` // Dynamic structure
	TenantID        *string                 `json:"tenantId,omitempty"`
}

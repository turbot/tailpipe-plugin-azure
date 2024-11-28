package rows

import (
	"time"

	"github.com/turbot/tailpipe-plugin-sdk/enrichment"
)

type ActivityLog struct {
	enrichment.CommonFields

	Authorization        *ActivityLogAuthorization `json:"authorization,omitempty"`
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

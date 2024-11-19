package rows

import (
	"time"

	"github.com/turbot/tailpipe-plugin-sdk/enrichment"
	"github.com/turbot/tailpipe-plugin-sdk/types"
)

type ActivityLog struct {
	enrichment.CommonFields

	AuthorizationAction            *string           `json:"authorization_action"`
	AuthorizationRole              *string           `json:"authorization_role"`
	AuthorizationScope             *string           `json:"authorization_scope"`
	Caller                         *string           `json:"caller"`
	Category                       *string           `json:"category"`
	Claims                         *types.JSONString `json:"claims"`
	CorrelationID                  *string           `json:"correlation_id"`
	Description                    *string           `json:"description"`
	EventDataID                    *string           `json:"event_data_id"`
	EventName                      *string           `json:"event_name"`
	EventTimestamp                 *time.Time        `json:"event_timestamp"`
	HttpRequestInfoClientIpAddress *string           `json:"http_request_info_client_ip_address"`
	HttpRequestInfoClientRequestId *string           `json:"http_request_info_client_request_id"`
	HttpRequestInfoMethod          *string           `json:"http_request_info_method"`
	HttpRequestInfoUri             *string           `json:"http_request_info_uri"`
	ID                             *string           `json:"id"`
	Level                          *string           `json:"level"`
	OperationID                    *string           `json:"operation_id"`
	OperationName                  *string           `json:"operation_name"`
	Properties                     *types.JSONString `json:"properties"`
	ResourceGroupName              *string           `json:"resource_group_name"`
	ResourceID                     *string           `json:"resource_id"`
	ResourceProviderName           *string           `json:"resource_provider_name"`
	ResourceType                   *string           `json:"resource_type"`
	Status                         *string           `json:"status"`
	SubStatus                      *string           `json:"sub_status"`
	SubmissionTimestamp            *time.Time        `json:"submission_timestamp"`
	SubscriptionID                 *string           `json:"subscription_id"`
	TenantID                       *string           `json:"tenant_id"`
}

func NewActivityLog() *ActivityLog {
	return &ActivityLog{}
}

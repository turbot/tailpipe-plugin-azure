package mappers

import (
	"context"
	"fmt"
	"github.com/turbot/tailpipe-plugin-sdk/table"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
	"github.com/turbot/tailpipe-plugin-azure/rows"
)

type ActivityLogMapper struct{}

func (m *ActivityLogMapper) Identifier() string {
	return "azure_activity_log_mapper"
}

func (m *ActivityLogMapper) Map(_ context.Context, a any, _ ...table.MapOption[*rows.ActivityLog]) (*rows.ActivityLog, error) {
	logEntry, ok := a.(*armmonitor.EventData)
	if !ok {
		return nil, fmt.Errorf("invalid row type: %T, expected *armmonitor.EventData", a)
	}

	var row rows.ActivityLog

	if logEntry.Authorization != nil {
		row.AuthorizationInfo = &rows.ActivityLogAuthorization{
			Action: logEntry.Authorization.Action,
			Scope:  logEntry.Authorization.Scope,
			Role:   logEntry.Authorization.Role,
		}
	}

	if logEntry.HTTPRequest != nil {
		row.HttpRequest = &rows.ActivityLogHttpRequest{
			ClientIpAddress: logEntry.HTTPRequest.ClientIPAddress,
			ClientRequestId: logEntry.HTTPRequest.ClientRequestID,
			Method:          logEntry.HTTPRequest.Method,
			Uri:             logEntry.HTTPRequest.URI,
		}
	}

	row.Caller = logEntry.Caller
	if logEntry.Category != nil {
		row.Category = logEntry.Category.Value
	}
	row.CorrelationID = logEntry.CorrelationID
	row.Description = logEntry.Description
	row.EventDataID = logEntry.EventDataID

	if logEntry.EventName != nil {
		row.EventName = logEntry.EventName.Value
	}

	row.EventTimestamp = logEntry.EventTimestamp
	row.ID = logEntry.ID
	row.Level = (*string)(logEntry.Level)
	row.OperationID = logEntry.OperationID
	if logEntry.OperationName != nil {
		row.OperationName = logEntry.OperationName.Value
	}
	row.ResourceGroupName = logEntry.ResourceGroupName
	row.ResourceID = logEntry.ResourceID
	if logEntry.ResourceProviderName != nil {
		row.ResourceProviderName = logEntry.ResourceProviderName.Value
	}
	if logEntry.ResourceType != nil {
		row.ResourceType = logEntry.ResourceType.Value
	}
	if logEntry.Status != nil {
		row.Status = logEntry.Status.Value
	}
	if logEntry.SubStatus != nil {
		row.SubStatus = logEntry.SubStatus.Value
	}
	row.SubmissionTimestamp = logEntry.SubmissionTimestamp
	row.SubscriptionID = logEntry.SubscriptionID
	row.TenantID = logEntry.TenantID

	if logEntry.Claims != nil {
		row.Claims = &logEntry.Claims
	}

	if logEntry.Properties != nil {
		row.Properties = &logEntry.Properties
	}

	return &row, nil
}

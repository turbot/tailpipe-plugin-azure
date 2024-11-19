package mappers

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"

	"github.com/turbot/tailpipe-plugin-azure/rows"
	"github.com/turbot/tailpipe-plugin-sdk/table"
	"github.com/turbot/tailpipe-plugin-sdk/types"
)

type ActivityLogMapper struct{}

func NewActivityLogMapper() table.Mapper[*rows.ActivityLog] {
	return &ActivityLogMapper{}
}

func (m *ActivityLogMapper) Identifier() string {
	return "azure_activity_log_mapper"
}

func (m *ActivityLogMapper) Map(_ context.Context, a any) ([]*rows.ActivityLog, error) {
	logEntry, ok := a.(*armmonitor.EventData)
	if !ok {
		return nil, fmt.Errorf("invalid row type: %T, expected *armmonitor.EventData", a)
	}

	row := rows.NewActivityLog()

	if logEntry.Authorization != nil {
		row.AuthorizationAction = logEntry.Authorization.Action
		row.AuthorizationScope = logEntry.Authorization.Scope
		row.AuthorizationRole = logEntry.Authorization.Role
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
		// marshal the claims to JSON
		claims, err := json.Marshal(logEntry.Claims)
		if err != nil {
			return nil, fmt.Errorf("error marshalling claims: %w", err)
		}
		claimsJson := types.JSONString(claims)
		row.Claims = &claimsJson
	}

	if logEntry.Properties != nil {
		// marshal the properties to JSON
		properties, err := json.Marshal(logEntry.Properties)
		if err != nil {
			return nil, fmt.Errorf("error marshalling properties: %w", err)
		}
		propertiesJson := types.JSONString(properties)
		row.Properties = &propertiesJson
	}

	return []*rows.ActivityLog{row}, nil
}

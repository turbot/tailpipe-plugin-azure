package activity_log

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"

	"github.com/turbot/tailpipe-plugin-sdk/table"
)

type ActivityLogMapper struct{}

func (m *ActivityLogMapper) Identifier() string {
	return "azure_activity_log_mapper"
}

func (m *ActivityLogMapper) Map(_ context.Context, a any, _ ...table.MapOption[*ActivityLog]) (*ActivityLog, error) {

	var row ActivityLog

	switch v := a.(type) {
	case string: // Activity log storage account artifact source
		var storageAccLog AzureStorageAccountLog
		err := json.Unmarshal([]byte(v), &storageAccLog)
		if err != nil {
			return nil, err
		}
		claims := make(map[string]*string)
		if storageAccLog.Identity != nil && storageAccLog.Identity.Claims != nil {
			for key, value := range *storageAccLog.Identity.Claims {
				if val, ok := value.(string); ok {
					claims[key] = &val
				}
			}
		}

		properties := make(map[string]*string)
		if storageAccLog.Properties != nil {
			for key, value := range *storageAccLog.Properties {
				if val, ok := value.(string); ok {
					properties[key] = &val
				}
			}
		}

		activityLog := &ActivityLog{
			Caller:               claims["name"],
			Category:             storageAccLog.Category,
			Claims:               &claims,
			CorrelationID:        storageAccLog.CorrelationID,
			Description:          storageAccLog.OperationName,
			EventDataID:          storageAccLog.OperationName, // Assuming EventDataID is tied to OperationName
			EventName:            storageAccLog.OperationName,
			EventTimestamp:       storageAccLog.Time,
			ID:                   storageAccLog.CorrelationID, // Assuming ID relates to CorrelationID
			Level:                storageAccLog.Level,
			OperationID:          storageAccLog.CorrelationID, // Assuming OperationID ties to CorrelationID
			OperationName:        storageAccLog.OperationName,
			Properties:           &properties,
			ResourceGroupName:    extractResourceGroup(*storageAccLog.ResourceID), // Extract from ResourceID
			ResourceID:           storageAccLog.ResourceID,
			ResourceProviderName: extractResourceProvider(*storageAccLog.ResourceID), // Extract from ResourceID
			ResourceType:         extractResourceType(*storageAccLog.ResourceID),     // Extract from ResourceID
			Status:               storageAccLog.ResultType,
			SubStatus:            storageAccLog.ResultSignature,
			SubmissionTimestamp:  storageAccLog.Time,
			SubscriptionID:       extractSubscriptionID(*storageAccLog.ResourceID), // Extract from ResourceID
			TenantID:             storageAccLog.TenantID,
			HttpRequest: &ActivityLogHttpRequest{
				ClientIpAddress: storageAccLog.CallerIPAddress,
			},
		}

		if storageAccLog.Identity != nil {
			activityLog.AuthorizationInfo = &ActivityLogAuthorization{
				Action: storageAccLog.Identity.Authorization.Action,
				Scope:  storageAccLog.Identity.Authorization.Scope,
				Role:   storageAccLog.Identity.Authorization.Evidence.Role,
			}
		}

		row = *activityLog

	case *armmonitor.EventData: // Activity log API source
		logEntry, ok := a.(*armmonitor.EventData)
		if !ok {
			return nil, fmt.Errorf("invalid row type: %T, expected *armmonitor.EventData", a)
		}
		if logEntry.Authorization != nil {
			row.AuthorizationInfo = &ActivityLogAuthorization{
				Action: logEntry.Authorization.Action,
				Scope:  logEntry.Authorization.Scope,
				Role:   logEntry.Authorization.Role,
			}
		}

		if logEntry.HTTPRequest != nil {
			row.HttpRequest = &ActivityLogHttpRequest{
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
	case []byte:
		err := json.Unmarshal(v, &row)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("expected *armmonitor.EventData, string or []byte, got %T", a)
	}

	return &row, nil
}

func extractResourceGroup(resourceID string) *string {
	// Extract Resource Group Name from the Resource ID
	parts := strings.Split(resourceID, "/")
	for i, part := range parts {
		if strings.EqualFold(part, "resourceGroups") && i+1 < len(parts) {
			return &parts[i+1]
		}
	}
	return nil
}

func extractResourceProvider(resourceID string) *string {
	// Extract Resource Provider Name from the Resource ID
	parts := strings.Split(resourceID, "/")
	for i, part := range parts {
		if strings.EqualFold(part, "providers") && i+1 < len(parts) {
			return &parts[i+1]
		}
	}
	return nil
}

func extractResourceType(resourceID string) *string {
	// Extract Resource Type from the Resource ID
	parts := strings.Split(resourceID, "/")
	for i := 0; i < len(parts)-1; i++ {
		if strings.EqualFold(parts[i], "providers") && i+2 < len(parts) {
			return &parts[i+2]
		}
	}
	return nil
}

func extractSubscriptionID(resourceID string) *string {
	// Extract Subscription ID from the Resource ID
	parts := strings.Split(resourceID, "/")
	for i, part := range parts {
		if strings.EqualFold(part, "subscriptions") && i+1 < len(parts) {
			return &parts[i+1]
		}
	}
	return nil
}

package azure_collection

import (
	"fmt"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
	"github.com/rs/xid"
	"github.com/turbot/tailpipe-plugin-azure/azure_types"
	"github.com/turbot/tailpipe-plugin-sdk/collection"
	"github.com/turbot/tailpipe-plugin-sdk/enrichment"
	"github.com/turbot/tailpipe-plugin-sdk/hcl"
	"github.com/turbot/tailpipe-plugin-sdk/helpers"
)

type ActivityLogCollection struct {
	collection.CollectionBase[*ActivityLogCollectionConfig]
}

func NewActivityLogCollection() collection.Collection {
	return &ActivityLogCollection{}
}

func (c *ActivityLogCollection) Identifier() string {
	return "azure_activity_log"
}

func (c *ActivityLogCollection) GetRowSchema() any {
	return azure_types.ActivityLogRow{}
}

func (c *ActivityLogCollection) GetConfigSchema() hcl.Config {
	return &ActivityLogCollectionConfig{}
}

func (c *ActivityLogCollection) EnrichRow(row any, sourceEnrichmentFields *enrichment.CommonFields) (any, error) {
	logEntry, ok := row.(*armmonitor.EventData)
	if !ok {
		return nil, fmt.Errorf("invalid row type: %T, expected *armmonitor.EventData", row)
	}

	if sourceEnrichmentFields == nil || sourceEnrichmentFields.TpConnection == "" {
		return nil, fmt.Errorf("source must provide connection in sourceEnrichmentFields")
	}

	record := &azure_types.ActivityLogRow{CommonFields: *sourceEnrichmentFields}

	if logEntry.Authorization != nil {
		record.AuthorizationAction = logEntry.Authorization.Action
		record.AuthorizationScope = logEntry.Authorization.Scope
		record.AuthorizationRole = logEntry.Authorization.Role
	}
	record.Caller = logEntry.Caller
	if logEntry.Category != nil {
		record.Category = logEntry.Category.Value
	}
	record.CorrelationID = logEntry.CorrelationID
	record.Description = logEntry.Description
	record.EventDataID = logEntry.EventDataID
	if logEntry.EventName != nil {
		record.EventName = logEntry.EventName.Value
	}
	record.EventTimestamp = logEntry.EventTimestamp
	record.ID = logEntry.ID
	record.Level = (*string)(logEntry.Level)
	record.OperationID = logEntry.OperationID
	if logEntry.OperationName != nil {
		record.OperationName = logEntry.OperationName.Value
	}
	record.ResourceGroupName = logEntry.ResourceGroupName
	record.ResourceID = logEntry.ResourceID
	if logEntry.ResourceProviderName != nil {
		record.ResourceProviderName = logEntry.ResourceProviderName.Value
	}
	if logEntry.ResourceType != nil {
		record.ResourceType = logEntry.ResourceType.Value
	}
	if logEntry.Status != nil {
		record.Status = logEntry.Status.Value
	}
	if logEntry.SubStatus != nil {
		record.SubStatus = logEntry.SubStatus.Value
	}
	record.SubmissionTimestamp = logEntry.SubmissionTimestamp
	record.SubscriptionID = logEntry.SubscriptionID
	record.TenantID = logEntry.TenantID

	// Record Standardization
	record.TpID = xid.New().String()
	record.TpSourceType = c.Identifier()
	record.TpTimestamp = helpers.UnixMillis(logEntry.EventTimestamp.UnixNano() / int64(time.Millisecond))
	record.TpIngestTimestamp = helpers.UnixMillis(time.Now().UnixNano() / int64(time.Millisecond))

	// TODO: #enrichment process more Tp fields from the logEntry

	// Hive Fields
	record.TpCollection = c.Identifier()
	record.TpYear = int32(logEntry.EventTimestamp.Year())
	record.TpMonth = int32(logEntry.EventTimestamp.Month())
	record.TpDay = int32(logEntry.EventTimestamp.Day())

	return record, nil
}

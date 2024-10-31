package tables

import (
	"fmt"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
	"github.com/rs/xid"
	"github.com/turbot/tailpipe-plugin-azure/models"
	"github.com/turbot/tailpipe-plugin-sdk/enrichment"
	"github.com/turbot/tailpipe-plugin-sdk/helpers"
	"github.com/turbot/tailpipe-plugin-sdk/parse"
	"github.com/turbot/tailpipe-plugin-sdk/table"
)

type ActivityLogTable struct {
	table.TableBase[*ActivityLogTableConfig]
}

func NewActivityLogTable() table.Table {
	return &ActivityLogTable{}
}

func (c *ActivityLogTable) Identifier() string {
	return "azure_activity_log"
}

func (c *ActivityLogTable) GetRowSchema() any {
	return models.ActivityLog{}
}

func (c *ActivityLogTable) GetConfigSchema() parse.Config {
	return &ActivityLogTableConfig{}
}

func (c *ActivityLogTable) EnrichRow(row any, sourceEnrichmentFields *enrichment.CommonFields) (any, error) {
	logEntry, ok := row.(*armmonitor.EventData)
	if !ok {
		return nil, fmt.Errorf("invalid row type: %T, expected *armmonitor.EventData", row)
	}

	if sourceEnrichmentFields == nil || sourceEnrichmentFields.TpIndex == "" {
		return nil, fmt.Errorf("source must provide connection in sourceEnrichmentFields")
	}

	record := &models.ActivityLog{CommonFields: *sourceEnrichmentFields}

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
	record.TpPartition = c.Identifier()
	record.TpDate = logEntry.EventTimestamp.Format("2006-01-02")

	return record, nil
}

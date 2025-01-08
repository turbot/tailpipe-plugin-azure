package tables

import (
	"strings"
	"time"

	"github.com/rs/xid"
	"github.com/turbot/tailpipe-plugin-azure/mappers"
	"github.com/turbot/tailpipe-plugin-azure/rows"
	"github.com/turbot/tailpipe-plugin-azure/sources"
	"github.com/turbot/tailpipe-plugin-sdk/schema"
	"github.com/turbot/tailpipe-plugin-sdk/table"
)

const ActivityLogTableIdentifier = "azure_activity_log"

// register the table from the package init function
func init() {
	// Register the table, with type parameters:
	// 1. row struct
	// 2. table config struct
	// 3. table implementation
	table.RegisterTable[*rows.ActivityLog, *ActivityLogTable]()
}

type ActivityLogTable struct {
}

func (c *ActivityLogTable) Identifier() string {
	return ActivityLogTableIdentifier
}

func (c *ActivityLogTable) GetSourceMetadata() []*table.SourceMetadata[*rows.ActivityLog] {
	return []*table.SourceMetadata[*rows.ActivityLog]{
		{
			SourceName: sources.ActivityLogAPISourceIdentifier,
			Mapper:     &mappers.ActivityLogMapper{},
		},
	}
}

func (c *ActivityLogTable) EnrichRow(row *rows.ActivityLog, sourceEnrichmentFields schema.SourceEnrichment) (*rows.ActivityLog, error) {
	row.CommonFields = sourceEnrichmentFields.CommonFields

	// Record Standardization
	row.TpID = xid.New().String()
	row.TpTimestamp = *row.EventTimestamp
	row.TpDate = row.EventTimestamp.Truncate(24 * time.Hour)
	row.TpIngestTimestamp = time.Now()
	if row.SubscriptionID != nil {
		subId := strings.ToLower(*row.SubscriptionID)
		row.TpIndex = subId
	} else {
		row.TpIndex = "default"
	}

	if row.HttpRequest != nil {
		if row.HttpRequest.ClientIpAddress != nil {
			row.TpSourceIP = row.HttpRequest.ClientIpAddress
			row.TpIps = append(row.TpIps, *row.HttpRequest.ClientIpAddress)
		}
	}

	if row.ResourceID != nil {
		row.TpAkas = append(row.TpAkas, *row.ResourceID)
	}

	return row, nil
}

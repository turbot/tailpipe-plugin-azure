package tables

import (
	"fmt"
	"time"

	"github.com/rs/xid"

	"github.com/turbot/tailpipe-plugin-azure/mappers"
	"github.com/turbot/tailpipe-plugin-azure/rows"
	"github.com/turbot/tailpipe-plugin-azure/sources"
	"github.com/turbot/tailpipe-plugin-sdk/enrichment"
	"github.com/turbot/tailpipe-plugin-sdk/table"
)

const ActivityLogTableIdentifier = "azure_activity_log"

// register the table from the package init function
func init() {
	// Register the table, with type parameters:
	// 1. row struct
	// 2. table config struct
	// 3. table implementation
	table.RegisterTable[*rows.ActivityLog, *ActivityLogTableConfig, *ActivityLogTable]()
}

type ActivityLogTable struct {
}

func (c *ActivityLogTable) Identifier() string {
	return ActivityLogTableIdentifier
}

func (c *ActivityLogTable) SupportedSources(_ *ActivityLogTableConfig) []*table.SourceMetadata[*rows.ActivityLog] {
	return []*table.SourceMetadata[*rows.ActivityLog]{
		{
			SourceName: sources.ActivityLogAPISourceIdentifier,
			MapperFunc: mappers.NewActivityLogMapper,
		},
	}
}

func (c *ActivityLogTable) EnrichRow(row *rows.ActivityLog, sourceEnrichmentFields *enrichment.CommonFields) (*rows.ActivityLog, error) {

	// we expect sourceEnrichmentFields to be set
	if sourceEnrichmentFields == nil {
		return nil, fmt.Errorf("ActivityLogTable EnrichRow called with nil sourceEnrichmentFields")
	}
	// we expect name to be set by the Source
	if sourceEnrichmentFields.TpSourceName == nil {
		return nil, fmt.Errorf("ActivityLogTable EnrichRow called with TpSourceName unset in sourceEnrichmentFields")
	}

	row.CommonFields = *sourceEnrichmentFields

	// Record Standardization
	row.TpID = xid.New().String()
	row.TpTimestamp = *row.EventTimestamp
	row.TpIngestTimestamp = time.Now()
	row.TpPartition = c.Identifier()
	row.TpIndex = *row.SubscriptionID

	// Hive Fields
	row.TpDate = row.EventTimestamp.Truncate(24 * time.Hour)

	return row, nil
}

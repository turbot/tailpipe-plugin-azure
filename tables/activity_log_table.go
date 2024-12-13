package tables

import (
	"strings"
	"time"

	"github.com/rs/xid"

	"github.com/turbot/pipe-fittings/utils"
	"github.com/turbot/tailpipe-plugin-azure/mappers"
	"github.com/turbot/tailpipe-plugin-azure/rows"
	"github.com/turbot/tailpipe-plugin-azure/sources"
	"github.com/turbot/tailpipe-plugin-sdk/artifact_source"
	"github.com/turbot/tailpipe-plugin-sdk/artifact_source_config"
	"github.com/turbot/tailpipe-plugin-sdk/enrichment"
	"github.com/turbot/tailpipe-plugin-sdk/row_source"
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

func (c *ActivityLogTable) GetSourceMetadata(_ *ActivityLogTableConfig) []*table.SourceMetadata[*rows.ActivityLog] {

	defaultArtifactConfig := &artifact_source_config.ArtifactSourceConfigBase{
		FileLayout: utils.ToStringPointer("/resourceId=/SUBSCRIPTIONS/[A-F0-9-]+/y=\\d{4}/m=\\d{2}/d=\\d{2}/h=\\d{2}/m=\\d{2}/PT\\d+H\\.json"),
	}

	return []*table.SourceMetadata[*rows.ActivityLog]{
		{
			SourceName: sources.ActivityLogAPISourceIdentifier,
			Mapper:     &mappers.ActivityLogMapper{},
		},
		{
			SourceName: sources.AzureBlobStorageSourceIdentifier,
			Mapper:     &mappers.ActivityLogMapper{},
			Options: []row_source.RowSourceOption{
				artifact_source.WithDefaultArtifactSourceConfig(defaultArtifactConfig),
				artifact_source.WithRowPerLine(),
			},
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

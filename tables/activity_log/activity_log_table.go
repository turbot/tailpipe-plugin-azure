package activity_log

import (
	"strings"
	"time"

	"github.com/rs/xid"

	"github.com/turbot/pipe-fittings/v2/utils"
	"github.com/turbot/tailpipe-plugin-azure/sources/activity_log_api"
	"github.com/turbot/tailpipe-plugin-azure/sources/blob_storage"
	"github.com/turbot/tailpipe-plugin-sdk/artifact_source"
	"github.com/turbot/tailpipe-plugin-sdk/artifact_source_config"
	"github.com/turbot/tailpipe-plugin-sdk/row_source"
	"github.com/turbot/tailpipe-plugin-sdk/schema"
	"github.com/turbot/tailpipe-plugin-sdk/table"
)

const ActivityLogTableIdentifier = "azure_activity_log"

type ActivityLogTable struct {
}

func (c *ActivityLogTable) Identifier() string {
	return ActivityLogTableIdentifier
}

func (c *ActivityLogTable) GetSourceMetadata() []*table.SourceMetadata[*ActivityLog] {

	defaultArtifactConfig := &artifact_source_config.ArtifactSourceConfigImpl{
		FileLayout: utils.ToStringPointer("/SUBSCRIPTIONS/%{DATA:subscription_id}/y=%{YEAR:year}/m=%{MONTHNUM:month}/d=%{MONTHDAY:day}/h=%{HOUR:hour}/m=%{MINUTE:minute}/%{DATA:filename}.json"),
	}

	return []*table.SourceMetadata[*ActivityLog]{
		{
			SourceName: activity_log_api.ActivityLogAPISourceIdentifier,
			Mapper:     &ActivityLogMapper{},
		},
		{
			SourceName: blob_storage.AzureBlobStorageSourceIdentifier,
			Mapper:     &ActivityLogMapper{},
			Options: []row_source.RowSourceOption{
				artifact_source.WithDefaultArtifactSourceConfig(defaultArtifactConfig),
				artifact_source.WithRowPerLine(),
			},
		},
	}
}

func (c *ActivityLogTable) EnrichRow(row *ActivityLog, sourceEnrichmentFields schema.SourceEnrichment) (*ActivityLog, error) {
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

func (c *ActivityLogTable) GetDescription() string {
	return "Azure Activity Logs record management operations and user actions performed on Azure resources, providing insight into administrative changes and service health."
}

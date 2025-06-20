package cost_and_usage_actual

import (
	"time"

	"github.com/rs/xid"

	"github.com/turbot/pipe-fittings/v2/utils"
	"github.com/turbot/tailpipe-plugin-sdk/artifact_source"
	"github.com/turbot/tailpipe-plugin-sdk/artifact_source_config"
	"github.com/turbot/tailpipe-plugin-sdk/constants"
	"github.com/turbot/tailpipe-plugin-sdk/row_source"
	"github.com/turbot/tailpipe-plugin-sdk/schema"
	"github.com/turbot/tailpipe-plugin-sdk/table"
)

const CostAndUsageActualTableIdentifier = "azure_cost_and_usage_actual"

// CostAndUsageActualTable - table for Azure Cost and Usage Actual data
type CostAndUsageActualTable struct{}

// Identifier implements table.Table
func (t *CostAndUsageActualTable) Identifier() string {
	return CostAndUsageActualTableIdentifier
}

func (t *CostAndUsageActualTable) GetSourceMetadata() ([]*table.SourceMetadata[*CostAndUsageActual], error) {
	defaultArtifactConfig := &artifact_source_config.ArtifactSourceConfigImpl{
		// Grok pattern to match Azure Cost Management export files
		// Pattern matches files like:  part_0_0001.csv.gz / part_1_0001.csv / part_2_0001.csv.zip
		FileLayout: utils.ToStringPointer("part_%{INT:part_number}_%{INT:file_number}.csv.(?:gz|zip)"),
	}

	return []*table.SourceMetadata[*CostAndUsageActual]{
		{
			// any artifact source
			SourceName: constants.ArtifactSourceIdentifier,
			Mapper:     NewCostAndUsageActualMapper(),
			Options: []row_source.RowSourceOption{
				artifact_source.WithDefaultArtifactSourceConfig(defaultArtifactConfig),
				artifact_source.WithRowPerLine(),
				artifact_source.WithHeaderRowNotification(","),
			},
		},
	}, nil
}

// EnrichRow implements table.Table
func (t *CostAndUsageActualTable) EnrichRow(row *CostAndUsageActual, sourceEnrichmentFields schema.SourceEnrichment) (*CostAndUsageActual, error) {
	// initialize the enrichment fields to any fields provided by the source
	row.CommonFields = sourceEnrichmentFields.CommonFields

	// Record standardization
	row.TpID = xid.New().String()
	row.TpIngestTimestamp = time.Now()

	// Set TpTimestamp and TpDate based on available date fields using priority-based function
	setTimestampByPriority(row)

	// Set TpIndex to default index
	row.TpIndex = schema.DefaultIndex

	// Set TpAkas to resource ID if available
	if row.ResourceId != nil {
		row.TpAkas = append(row.TpAkas, *row.ResourceId)
	}

	return row, nil
}

func (t *CostAndUsageActualTable) GetDescription() string {
	return "Azure Cost and Usage data provides detailed information about Azure resource usage and costs, including subscription charges, resource consumption, pricing details, and billing information. This table enables cost analysis, budget tracking, and optimization insights across Azure subscriptions."
}

// setTimestampByPriority sets the timestamp based on priority order of available date fields
func setTimestampByPriority(row *CostAndUsageActual) {
	// Priority order: Date > ServicePeriodStartDate > ServicePeriodEndDate > BillingPeriodStartDate > BillingPeriodEndDate
	dateFields := []*time.Time{
		row.Date,
		row.ServicePeriodStartDate,
		row.ServicePeriodEndDate,
		row.BillingPeriodStartDate,
		row.BillingPeriodEndDate,
	}

	for _, dateField := range dateFields {
		if dateField != nil {
			truncatedDate := dateField.Truncate(24 * time.Hour)
			row.TpTimestamp = *dateField
			row.TpDate = truncatedDate
			return
		}
	}
}

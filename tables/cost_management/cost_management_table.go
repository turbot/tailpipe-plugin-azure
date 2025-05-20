package cost_management

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

const CostManagementTableIdentifier = "azure_cost_management"

// CostManagementTable - table for Azure Cost Management data
type CostManagementTable struct{}

// Identifier implements table.Table
func (t *CostManagementTable) Identifier() string {
	return CostManagementTableIdentifier
}

func (t *CostManagementTable) GetSourceMetadata() ([]*table.SourceMetadata[*CostManagement], error) {
	defaultArtifactConfig := &artifact_source_config.ArtifactSourceConfigImpl{
		// Grok pattern to match Azure Cost Management export files
		// Pattern matches files like:
		// - part_0_0001.csv.gz
		// - part_1_0001.csv.gz
		// This pattern can be adjusted based on the actual file naming convention
		FileLayout: utils.ToStringPointer("part_%{INT:part_number}_%{INT:file_number}.csv.(?:gz|zip)"),
	}

	return []*table.SourceMetadata[*CostManagement]{
		{
			// any artifact source
			SourceName: constants.ArtifactSourceIdentifier,
			Mapper:     NewCostManagementMapper(),
			Options: []row_source.RowSourceOption{
				artifact_source.WithDefaultArtifactSourceConfig(defaultArtifactConfig),
				artifact_source.WithRowPerLine(),
				artifact_source.WithHeaderRowNotification(","),
			},
		},
	}, nil
}

// EnrichRow implements table.Table
func (t *CostManagementTable) EnrichRow(row *CostManagement, sourceEnrichmentFields schema.SourceEnrichment) (*CostManagement, error) {
	// initialize the enrichment fields to any fields provided by the source
	row.CommonFields = sourceEnrichmentFields.CommonFields

	// Record standardization
	row.TpID = xid.New().String()
	row.TpIngestTimestamp = time.Now()

	// Set TpTimestamp based on available date fields
	if row.Date != nil {
		row.TpTimestamp = *row.Date
		row.TpDate = row.Date.Truncate(24 * time.Hour)
	} else if row.ServicePeriodStartDate != nil {
		row.TpTimestamp = *row.ServicePeriodStartDate
		row.TpDate = row.ServicePeriodStartDate.Truncate(24 * time.Hour)
	} else if row.ServicePeriodEndDate != nil {
		row.TpTimestamp = *row.ServicePeriodEndDate
		row.TpDate = row.ServicePeriodEndDate.Truncate(24 * time.Hour)
	} else if row.BillingPeriodStartDate != nil {
		row.TpTimestamp = *row.BillingPeriodStartDate
		row.TpDate = row.BillingPeriodStartDate.Truncate(24 * time.Hour)
	} else if row.BillingPeriodEndDate != nil {
		row.TpTimestamp = *row.BillingPeriodEndDate
		row.TpDate = row.BillingPeriodEndDate.Truncate(24 * time.Hour)
	}

	row.TpIndex = schema.DefaultIndex

	// Set TpAkas to resource ID if available
	if row.ResourceId != nil {
		row.TpAkas = append(row.TpAkas, *row.ResourceId)
	}

	return row, nil
}

func (t *CostManagementTable) GetDescription() string {
	return "Azure Cost Management data provides detailed information about Azure resource usage and costs, including subscription charges, resource consumption, pricing details, and billing information. This table enables cost analysis, budget tracking, and optimization insights across Azure subscriptions."
}

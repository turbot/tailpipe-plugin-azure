package tables

import (
	"context"
	"fmt"
	"time"

	"github.com/rs/xid"

	"github.com/turbot/tailpipe-plugin-azure/config"
	"github.com/turbot/tailpipe-plugin-azure/mappers"
	"github.com/turbot/tailpipe-plugin-azure/rows"
	"github.com/turbot/tailpipe-plugin-sdk/enrichment"
	"github.com/turbot/tailpipe-plugin-sdk/parse"
	"github.com/turbot/tailpipe-plugin-sdk/table"
	"github.com/turbot/tailpipe-plugin-sdk/types"
)

type ActivityLogTable struct {
	table.TableImpl[*rows.ActivityLog, *ActivityLogTableConfig, *config.AzureConnection]
}

func NewActivityLogTable() table.Table {
	return &ActivityLogTable{}
}

func (c *ActivityLogTable) Init(ctx context.Context, connectionSchemaProvider table.ConnectionSchemaProvider, req *types.CollectRequest) error {
	// call base init
	if err := c.TableImpl.Init(ctx, connectionSchemaProvider, req); err != nil {
		return err
	}

	c.initMapper()
	return nil
}

func (c *ActivityLogTable) initMapper() {
	// TODO switch on source

	c.Mapper = mappers.NewActivityLogMapper()
}

func (c *ActivityLogTable) Identifier() string {
	return "azure_activity_log"
}

func (c *ActivityLogTable) GetRowSchema() any {
	return rows.ActivityLog{}
}

func (c *ActivityLogTable) GetConfigSchema() parse.Config {
	return &ActivityLogTableConfig{}
}

func (c *ActivityLogTable) EnrichRow(row *rows.ActivityLog, sourceEnrichmentFields *enrichment.CommonFields) (*rows.ActivityLog, error) {

	if sourceEnrichmentFields == nil || sourceEnrichmentFields.TpIndex == "" {
		return nil, fmt.Errorf("source must provide connection in sourceEnrichmentFields")
	}
	// Record Standardization
	row.TpID = xid.New().String()
	row.TpTimestamp = *row.EventTimestamp
	row.TpIngestTimestamp = time.Now()
	row.TpPartition = c.Identifier()
	row.TpIndex = *row.SubscriptionID

	// TODO: #enrichment process more Tp fields from the row

	// Hive Fields
	row.TpDate = row.EventTimestamp.Format("2006-01-02")

	return row, nil
}

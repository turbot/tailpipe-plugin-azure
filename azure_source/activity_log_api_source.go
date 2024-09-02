package azure_source

import (
	"context"
	"fmt"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
	"github.com/turbot/tailpipe-plugin-sdk/enrichment"
	"github.com/turbot/tailpipe-plugin-sdk/parse"
	"github.com/turbot/tailpipe-plugin-sdk/row_source"
	"github.com/turbot/tailpipe-plugin-sdk/types"
)

const ActivityLogAPISourceIdentifier = "azure_activity_log_api"

type ActivityLogAPISource struct {
	row_source.RowSourceBase[*ActivityLogAPISourceConfig]
}

func NewActivityLogAPISource() row_source.RowSource {
	return &ActivityLogAPISource{}
}

func (s *ActivityLogAPISource) Identifier() string {
	return ActivityLogAPISourceIdentifier
}

func (s *ActivityLogAPISource) Collect(ctx context.Context) error {
	// TODO: #paging implement paging
	client, err := s.getClient()
	if err != nil {
		return err
	}
	// client doesn't have s Close() method

	sourceEnrichmentFields := &enrichment.CommonFields{
		TpSourceType: ActivityLogAPISourceIdentifier,
		TpIndex:      s.Config.SubscriptionId,
		// TODO: #enrichment can we add more source fields?
	}

	// TODO: #config move these to config
	startTime := time.Now().Add(-2159 * time.Hour) // 89 days 23 hours => { "code" : "BadRequest", "message" : "The start time cannot be more than 90 days in the past."}
	endTime := time.Now()

	filter := fmt.Sprintf("eventTimestamp ge '%s' and eventTimestamp le '%s'", startTime.Format(time.RFC3339Nano), endTime.Format(time.RFC3339Nano))
	pager := client.NewListPager(filter, nil)

	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return fmt.Errorf("failed to get next page: %w", err)
		}

		for _, logEntry := range page.Value {
			row := &types.RowData{
				Data:     logEntry,
				Metadata: sourceEnrichmentFields,
			}

			if err := s.OnRow(ctx, row, nil); err != nil {
				return fmt.Errorf("failed to processing row: %w", err)
			}
		}
	}

	return nil
}

func (s *ActivityLogAPISource) GetConfigSchema() parse.Config {
	return &ActivityLogAPISourceConfig{}
}

func (s *ActivityLogAPISource) getClient() (*armmonitor.ActivityLogsClient, error) {
	// TODO: #authentication support other authentication methods
	cred, err := azidentity.NewClientSecretCredential(s.Config.TenantId, s.Config.ClientId, s.Config.ClientSecret, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create client secret credential: %w", err)
	}

	client, err := armmonitor.NewActivityLogsClient(s.Config.SubscriptionId, cred, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create activity logs client: %w", err)
	}

	return client, nil
}

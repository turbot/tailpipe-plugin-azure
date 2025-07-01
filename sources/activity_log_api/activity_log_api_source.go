package activity_log_api

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"

	"github.com/turbot/tailpipe-plugin-azure/config"
	"github.com/turbot/tailpipe-plugin-sdk/collection_state"
	"github.com/turbot/tailpipe-plugin-sdk/row_source"
	"github.com/turbot/tailpipe-plugin-sdk/schema"
	"github.com/turbot/tailpipe-plugin-sdk/types"
)

const ActivityLogAPISourceIdentifier = "azure_activity_log_api"

type ActivityLogAPISource struct {
	row_source.RowSourceImpl[*ActivityLogAPISourceConfig, *config.AzureConnection]
}

func (s *ActivityLogAPISource) Init(ctx context.Context, params *row_source.RowSourceParams, opts ...row_source.RowSourceOption) error {
	// set collection order to reverse
	s.CollectionOrder = collection_state.CollectionOrderReverse

	// set the collection state ctor
	s.NewCollectionStateFunc = collection_state.NewTimeRangeCollectionState

	// call base init
	if err := s.RowSourceImpl.Init(ctx, params, opts...); err != nil {
		return err
	}

	return nil
}

func (s *ActivityLogAPISource) Identifier() string {
	return ActivityLogAPISourceIdentifier
}

func (s *ActivityLogAPISource) Collect(ctx context.Context) error {

	client, err := s.getClient(ctx) // client doesn't have a Close() method, nothing to defer
	if err != nil {
		return err
	}

	tpSource := ActivityLogAPISourceIdentifier
	sourceEnrichmentFields := &schema.SourceEnrichment{
		CommonFields: schema.CommonFields{
			TpSourceType: ActivityLogAPISourceIdentifier,
			TpSourceName: &tpSource,
		},
	}

	// set the collection time range
	toTime := s.CollectionTimeRange.EndTime()
	fromTime := s.CollectionTimeRange.StartTime()
	// limit 'from' to 90 days in the past, as per Azure API limits
	if time.Since(fromTime) > 2160*time.Hour {
		slog.Warn("the start time is more than 90 days in the past, adjusting to 90 days ago")
		fromTime = toTime.Add(-2160 * time.Hour) // 90 days in the past
	}

	filter := fmt.Sprintf("eventTimestamp ge '%s' and eventTimestamp le '%s'", fromTime.Format(time.RFC3339Nano), toTime.Format(time.RFC3339Nano))
	pager := client.NewListPager(filter, nil)

	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return fmt.Errorf("failed to get next page: %w", err)
		}

		for _, logEntry := range page.Value {

			// check if we've hit previous item - return false if we have, return from function
			if !s.CollectionState.ShouldCollect(*logEntry.ID, *logEntry.EventTimestamp) {
				return nil
			}

			row := &types.RowData{
				Data:             logEntry,
				SourceEnrichment: sourceEnrichmentFields,
			}

			// update collection state
			err := s.CollectionState.OnCollected(*logEntry.ID, *logEntry.EventTimestamp)
			if err != nil {
				return fmt.Errorf("failed to update collection state: %w", err)
			}

			if err := s.OnRow(ctx, row); err != nil {
				return fmt.Errorf("failed to processing row: %w", err)
			}
		}
	}

	return nil
}

func (s *ActivityLogAPISource) getClient(_ context.Context) (*armmonitor.ActivityLogsClient, error) {
	sess, err := s.Connection.GetSession()
	if err != nil {
		return nil, fmt.Errorf("failed to get session: %w", err)
	}

	client, err := armmonitor.NewActivityLogsClient(sess.SubscriptionID, sess.Credential, sess.ClientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to create activity logs client: %w", err)
	}

	return client, nil
}

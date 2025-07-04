package activity_log_api

import (
	"context"
	"fmt"
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
	// (note that although we collect backwards, 'from' is still the lower boundary time and 'to' is the upper boundary time,
	// as we poass them top the filter in the correct order)
	toTime := s.CollectionTimeRange.UpperBoundary
	fromTime := s.CollectionTimeRange.LowerBoundary
	// limit 'from' to 90 days in the past, as per Azure API limits
	if time.Since(fromTime) > 2160*time.Hour {
		return fmt.Errorf("from time %s is more than 90 days in the past, which exceeds Azure API limits", fromTime.Format(time.RFC3339Nano))
	}

	filter := fmt.Sprintf("eventTimestamp ge '%s' and eventTimestamp le '%s'", fromTime.Format(time.RFC3339Nano), toTime.Format(time.RFC3339Nano))
	pager := client.NewListPager(filter, nil)

	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return fmt.Errorf("failed to get next page: %w", err)
		}

		for _, logEntry := range page.Value {

			// check if we've hit previous item - continue if we have
			if !s.CollectionState.ShouldCollect(*logEntry.ID, *logEntry.EventTimestamp) {
				continue
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

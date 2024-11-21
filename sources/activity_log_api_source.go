package sources

import (
	"context"
	"fmt"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"

	"github.com/turbot/tailpipe-plugin-azure/config"
	"github.com/turbot/tailpipe-plugin-sdk/collection_state"
	"github.com/turbot/tailpipe-plugin-sdk/config_data"
	"github.com/turbot/tailpipe-plugin-sdk/enrichment"
	"github.com/turbot/tailpipe-plugin-sdk/row_source"
	"github.com/turbot/tailpipe-plugin-sdk/types"
)

const ActivityLogAPISourceIdentifier = "azure_activity_log_api"

// register the source from the package init function
func init() {
	row_source.RegisterRowSource[*ActivityLogAPISource]()
}

type ActivityLogAPISource struct {
	row_source.RowSourceImpl[*ActivityLogAPISourceConfig, *config.AzureConnection]
}

func (s *ActivityLogAPISource) Init(ctx context.Context, configData config_data.ConfigData, connectionData config_data.ConfigData, opts ...row_source.RowSourceOption) error {
	// set the collection state ctor
	s.NewCollectionStateFunc = collection_state.NewTimeRangeCollectionState

	// call base init
	return s.RowSourceImpl.Init(ctx, configData, connectionData, opts...)
}

func (s *ActivityLogAPISource) Identifier() string {
	return ActivityLogAPISourceIdentifier
}

func (s *ActivityLogAPISource) Collect(ctx context.Context) error {
	// NOTE: The API only allows fetching from newest to oldest, so we need to collect in reverse order until we've hit a previously obtain item.
	collectionState := s.CollectionState.(*collection_state.TimeRangeCollectionState[*ActivityLogAPISourceConfig])
	collectionState.IsChronological = false
	collectionState.HasContinuation = true
	collectionState.StartCollection() // sets previous state to current state as we manipulate the current state

	client, err := s.getClient(ctx) // client doesn't have a Close() method, nothing to defer
	if err != nil {
		return err
	}

	tpSource := ActivityLogAPISourceIdentifier
	sourceEnrichmentFields := &enrichment.CommonFields{
		TpSourceType: ActivityLogAPISourceIdentifier,
		TpSourceName: &tpSource,
	}

	endTime := time.Now()
	startTime := endTime.Add(-2160 * time.Hour) // 2160hr == 90 days => { "code" : "BadRequest", "message" : "The start time cannot be more than 90 days in the past."}

	if !collectionState.IsEmpty() {
		latestEndTime := collectionState.GetLatestEndTime()
		if latestEndTime != nil && latestEndTime.After(startTime) {
			startTime = *latestEndTime
		}
	}

	filter := fmt.Sprintf("eventTimestamp ge '%s' and eventTimestamp le '%s'", startTime.Format(time.RFC3339Nano), endTime.Format(time.RFC3339Nano))
	pager := client.NewListPager(filter, nil)

	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return fmt.Errorf("failed to get next page: %w", err)
		}

		for _, logEntry := range page.Value {

			// check if we've hit previous item - return false if we have, return from function
			if !collectionState.ShouldCollectRow(*logEntry.EventTimestamp, *logEntry.ID) {
				return nil
			}

			row := &types.RowData{
				Data:     logEntry,
				Metadata: sourceEnrichmentFields,
			}

			// update collection state
			collectionState.Upsert(*logEntry.EventTimestamp, *logEntry.ID, nil)
			collectionStateJSON, err := s.GetCollectionStateJSON()
			if err != nil {
				return fmt.Errorf("error serialising collectionState data: %w", err)
			}

			if err := s.OnRow(ctx, row, collectionStateJSON); err != nil {
				// TODO #errorHandling - this does not bubble up
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

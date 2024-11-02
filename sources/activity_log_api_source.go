package sources

import (
	"context"
	"fmt"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
	"github.com/turbot/tailpipe-plugin-sdk/collection_state"
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

func (s *ActivityLogAPISource) Init(ctx context.Context, configData *types.ConfigData, opts ...row_source.RowSourceOption) error {
	// set the collection state ctor
	s.NewCollectionStateFunc = collection_state.NewTimeRangeCollectionState

	// call base init
	return s.RowSourceBase.Init(ctx, configData, opts...)
}

func (s *ActivityLogAPISource) Identifier() string {
	return ActivityLogAPISourceIdentifier
}

func (s *ActivityLogAPISource) Collect(ctx context.Context) error {
	// NOTE: The API only allows fetching from newest to oldest, so we need to collect in reverse order until we've hit a previously obtain item.
	collectionState := s.CollectionState.(*collection_state.TimeRangeCollectionState[*ActivityLogAPISourceConfig])
	// TODO: #config the below should be settable via a config option
	collectionState.IsChronological = false
	collectionState.HasContinuation = true
	// TODO: #collectionState is there a way we can call StartCollection/EndCollection from elsewhere to enforce it?
	collectionState.StartCollection() // sets previous state to current state as we manipulate the current state

	client, err := s.getClient() // client doesn't have a Close() method, nothing to defer
	if err != nil {
		return err
	}

	sourceEnrichmentFields := &enrichment.CommonFields{
		TpSourceType: ActivityLogAPISourceIdentifier,
		TpIndex:      s.Config.SubscriptionId,
		// TODO: #enrichment can we add more source fields?
	}

	// TODO: #config should we move these to config?
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

package azure_source

import (
	"github.com/turbot/tailpipe-plugin-sdk/collection_state"
	"time"
)

type ActivityLogAPICollectionState struct {
	collection_state.CollectionStateBase

	StartTime time.Time `json:"start_time,omitempty"` // oldest record timestamp
	EndTime   time.Time `json:"end_time,omitempty"`   // newest record timestamp

	prevTime time.Time `json:"-"`
	// NOTE: Whilst haven't seen evidence of this, it may be possible to have 2 records in the same timestamp; to combat this we may wish to store map of IDs at the timestamps akin to ArtifactCollectionState.
}

func NewActivityLogAPICollectionState() collection_state.CollectionState[*ActivityLogAPISourceConfig] {
	return &ActivityLogAPICollectionState{}
}

func (s *ActivityLogAPICollectionState) Init(*ActivityLogAPISourceConfig) error {
	return nil
}

func (s *ActivityLogAPICollectionState) IsEmpty() bool {
	return s.StartTime.IsZero() // && s.EndTime.IsZero()
}

func (s *ActivityLogAPICollectionState) Upsert(createdAt time.Time) {
	if s.StartTime.IsZero() || createdAt.Before(s.StartTime) {
		s.StartTime = createdAt
	}

	if s.EndTime.IsZero() || createdAt.After(s.EndTime) {
		s.EndTime = createdAt
	}
}

// StartCollection stores the current state as previous state
func (s *ActivityLogAPICollectionState) StartCollection() {
	s.prevTime = s.EndTime
}

func (s *ActivityLogAPICollectionState) ShouldCollectRow(createdAt time.Time) bool {
	if !s.prevTime.IsZero() && createdAt.Equal(s.prevTime) {
		return false
	}

	return true
}

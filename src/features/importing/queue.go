package importing

import "github.com/nokevah/sonicary/src/music"

// QueueItemType represents the type of item in the queue

const (
	ManualReview    music.QueueItemType = "manual_review"
	Duplicate       music.QueueItemType = "duplicate"
	FailedImport    music.QueueItemType = "failed_import"
	MissingMetadata music.QueueItemType = "missing_metadata"
)

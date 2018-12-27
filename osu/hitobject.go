package osu

import "github.com/oklog/ulid"

type HitObject interface {
	// GetULID returns a unique identifier for this HitObject. This doesn't
	// necessarily need to persist between different instances of the editor
	// (i.e. doesn't need to be saved to disk)
	GetULID() ulid.ULID

	GetStartTime() Timestamp
}

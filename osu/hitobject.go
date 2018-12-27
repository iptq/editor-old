package osu

import "github.com/oklog/ulid"

type HitObject interface{}

type hoHeap []HitObject

// HitObjects is a container for multiple hitobjects. They will be stored in
// both a heap indexed by timestamp, as well as a hashmap (default Golang impl)
// simultaneously. This allows for both O(log(n)) reordering and O(1) access.
type HitObjects struct {
	objheap hoHeap
	objmap  map[ulid.ULID]*HitObject
}

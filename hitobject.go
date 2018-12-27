package editor

import (
	"container/heap"
	"editor/osu"

	"github.com/oklog/ulid"
)

// hobjHeap is a heap of HitObjects
type hobjHeap []*osu.HitObject

func (h hobjHeap) Len() int {
	return len(h)
}

func (h hobjHeap) Less(i, j int) bool {
	a := (*h[i]).GetStartTime()
	b := (*h[j]).GetStartTime()
	return a.Milliseconds() < b.Milliseconds()
}

func (h hobjHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h hobjHeap) Push(item interface{}) {
	obj := item.(*osu.HitObject)
	h = append(h, obj)
}

func (h hobjHeap) Pop() interface{} {
	n := len(h) - 1
	item := h[n]
	h = h[:n]
	return item
}

// HitObjects is a container for multiple hitobjects. They will be stored in
// both a heap indexed by timestamp, as well as a hashmap (default Golang impl)
// simultaneously. This allows for both O(log(n)) reordering and O(1) access.
type HitObjects struct {
	objheap hobjHeap
	objmap  map[ulid.ULID]*osu.HitObject
}

func HitObjectsFrom(objs []*osu.HitObject) *HitObjects {
	h := hobjHeap(objs)
	heap.Init(h)

	m := make(map[ulid.ULID]*osu.HitObject)
	for _, obj := range objs {
		m[(*obj).GetULID()] = obj
	}

	return &HitObjects{
		objheap: h,
		objmap:  m,
	}
}

package editor

import (
	"editor/osu"
	"fmt"
	"math/rand"

	"github.com/oklog/ulid"
)

type htNode struct {
	ulid  ulid.ULID
	time  osu.Timestamp
	obj   *osu.HitObject
	left  *htNode
	right *htNode
}

func bstInsert(n *htNode, obj *osu.HitObject) *htNode {
	if n == nil {
		return &htNode{
			ulid:  (*obj).GetULID(),
			time:  (*obj).GetStartTime(),
			obj:   obj,
			left:  nil,
			right: nil,
		}
	} else {
		time := (*obj).GetStartTime()
		if n.time.Milliseconds() < time.Milliseconds() {
			n.left = bstInsert(n.left, obj)
		} else {
			n.right = bstInsert(n.right, obj)
		}
		return n
	}
}

// HitObjects is a container for multiple hitobjects. They will be stored in
// both a BST indexed by timestamp, as well as a hashmap (default Golang impl)
// simultaneously. This allows for both O(log(n)) insertion and O(n) traversal.
type HitObjects struct {
	objBst *htNode
	objMap map[ulid.ULID]*osu.HitObject
}

func HitObjectsFrom(objs []*osu.HitObject) *HitObjects {
	// shuffle to get a (hopefully!) even tree
	rand.Shuffle(len(objs), func(i, j int) {
		objs[i], objs[j] = objs[j], objs[i]
	})

	var b *htNode = nil
	for _, obj := range objs {
		b = bstInsert(b, obj)
	}
	fmt.Println(b)

	m := make(map[ulid.ULID]*osu.HitObject)
	for _, obj := range objs {
		m[(*obj).GetULID()] = obj
	}

	return &HitObjects{
		objBst: b,
		objMap: m,
	}
}

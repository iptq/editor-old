package editor

import "editor/osu"

// EditorBeatmap is a wrapper around the traditional beatmap struct giving it
// more features and functions
type EditorBeatmap struct {
	*osu.Beatmap
	objects *HitObjects
}

func EditorBeatmapFrom(b *osu.Beatmap) *EditorBeatmap {
	objects := HitObjectsFrom(b.HitObjects)
	return &EditorBeatmap{
		Beatmap: b,
		objects: objects,
	}
}

// GetVisibleObjects returns all objects that can be seen at a particular
// instant in time.
func (m EditorBeatmap) GetVisibleObjects(timestamp int) {

}

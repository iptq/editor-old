package editor

import "editor/osu"

// EditorBeatmap is a wrapper around the traditional beatmap struct giving it
// more features and functions
type EditorBeatmap struct {
	inner *osu.Beatmap
}

// GetVisibleObjects returns all objects that can be seen at a particular
// instant in time.
func (m EditorBeatmap) GetVisibleObjects(timestamp int) {

}

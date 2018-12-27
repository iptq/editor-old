package osu

import (
	"fmt"
	"sort"
)

var (
	// allow objects to be up to 2 milliseconds off
	ESTIMATE_THRESHOLD = 2.0

	// list of snappings that the editor uses
	SNAPPINGS = []int{1, 2, 3, 4, 6, 8, 12, 16}
)

type Timestamp interface {
	Milliseconds() int
}

type TimestampAbsolute struct {
	inner int
}

func (t *TimestampAbsolute) Milliseconds() int {
	return t.inner
}

type snapping struct {
	Num   int
	Denom int
	Delta float64
}

type snappings []snapping

func (s snappings) Len() int {
	return len(s)
}

func (s snappings) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s snappings) Less(i, j int) bool {
	return s[i].Delta < s[j].Delta
}

// IntoRelative attempts to convert an absolute timestamp into a relative one
func (t *TimestampAbsolute) IntoRelative(to *TimingPoint) (*TimestampRelative, error) {
	bpm := (*to).GetBPM()
	meter := (*to).GetMeter()

	msPerBeat := 60000.0 / bpm
	msPerMeasure := msPerBeat * float64(meter)

	base := t.inner
	cur := t.inner

	measures := int(float64(cur-base) / msPerMeasure)
	measureStart := float64(base) + float64(measures)*msPerMeasure
	offset := float64(cur) - measureStart

	snapTimes := make([]snapping, len(SNAPPINGS)*16)
	for _, denom := range SNAPPINGS {
		for i := 0; i < denom; i++ {
			var snapAt float64

			snapAt = msPerMeasure * float64(i) / float64(denom)
			snapTimes = append(snapTimes, snapping{
				Num:   i,
				Denom: denom,
				Delta: offset - snapAt,
			})

			snapAt = msPerMeasure * float64(i+denom) / float64(denom)
			snapTimes = append(snapTimes, snapping{
				Num:   i + denom,
				Denom: denom,
				Delta: offset - snapAt,
			})
		}
	}
	sort.Sort(snappings(snapTimes))

	first := snapTimes[0]
	if first.Delta > ESTIMATE_THRESHOLD {
		return nil, fmt.Errorf("Could not find accurate snapping.")
	}

	t2 := &TimestampRelative{
		previous: to,
		measures: measures,
	}
	return t2, nil
}

type TimestampRelative struct {
	previous *TimingPoint
	measures int
	num      int
	denom    int
}

func (t *TimestampRelative) Milliseconds() int {
	base := (*t.previous).GetTimestamp().Milliseconds()
	bpm := (*t.previous).GetBPM()

	msPerBeat := 60000.0 / bpm
	measures := float64(t.measures) + float64(t.num)/float64(t.denom)

	return int(float64(base) + measures*msPerBeat)
}

type TimingPoint interface {
	// Get the timestamp
	GetTimestamp() Timestamp

	// Get the BPM of the nearest uninherited timing section to which this belongs
	GetBPM() float64

	// Get the meter of the nearest uninherited timing section to which this belongs
	GetMeter() int
}

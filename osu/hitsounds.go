package osu

type SampleSet = int

var SAMPLE_SETS = map[int]string{1: "Normal", 2: "Soft", 3: "Drum"}
var SAMPLE_SETS_INV = map[string]int{"normal": 1, "soft": 2, "drum": 3}

const (
	SAMPLE_NORMAL = 1
	SAMPLE_SOFT   = 2
	SAMPLE_DRUM   = 3
)

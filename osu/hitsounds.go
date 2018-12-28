package osu

import (
	"fmt"
	"strconv"
	"strings"
)

type SampleSet = int

var SAMPLE_SETS = map[int]string{1: "Normal", 2: "Soft", 3: "Drum"}
var SAMPLE_SETS_INV = map[string]int{"normal": 1, "soft": 2, "drum": 3}

const (
	SAMPLE_NORMAL = 1
	SAMPLE_SOFT   = 2
	SAMPLE_DRUM   = 3
)

type Hitsound = int

const (
	HITSOUND_NORMAL  = 1
	HITSOUND_WHISTLE = 2
	HITSOUND_FINISH  = 4
	HITSOUND_CLAP    = 8
)

type Extras struct {
	SampleSet    int
	AdditionSet  int
	CustomIndex  int
	SampleVolume int
	Filename     string
}

func ParseExtras(line string) (*Extras, error) {
	parts := strings.Split(line, ":")
	if strings.Count(line, ":") == 0 {
		// technically the extras field is optional, so if it's blank, assume "0:0:0:0:"
		return &Extras{}, nil
	} else if len(parts) != 5 {
		return nil, fmt.Errorf("len(extras) = %d != 5", len(parts))
	}

	sampleSet, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil, err
	}

	additionSet, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, err
	}

	customIndex, err := strconv.Atoi(parts[2])
	if err != nil {
		return nil, err
	}

	sampleVolume, err := strconv.Atoi(parts[3])
	if err != nil {
		return nil, err
	}

	extras := &Extras{sampleSet, additionSet, customIndex, sampleVolume, parts[4]}
	return extras, nil
}

func (extras Extras) String() string {
	return fmt.Sprintf("%d:%d:%d:%d:%s",
		extras.SampleSet,
		extras.AdditionSet,
		extras.CustomIndex,
		extras.SampleVolume,
		extras.Filename,
	)
}

package osu

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Mode = int

const (
	MODE_STD   = 0
	MODE_TAIKO = 1
	MODE_CTB   = 2
	MODE_MANIA = 3
)

type Beatmap struct {
	Version int

	AudioFilename     string
	AudioLeadIn       int
	PreviewTime       int
	Countdown         bool
	SampleSet         SampleSet
	StackLeniency     float64
	Mode              Mode
	LetterboxInBreaks int

	Title          string
	TitleUnicode   string
	Artist         string
	ArtistUnicode  string
	Creator        string
	DifficultyName string
	Source         string
	Tags           []string

	HPDrainRate       float64
	CircleSize        float64
	OverallDifficulty float64
	ApproachRate      float64
	SliderMultiplier  float64
	SliderTickRate    int

	Colors       []Color
	TimingPoints []TimingPoint
	HitObjects   []HitObject

	// TODO: events
}

var (
	FILE_FORMAT_PATTERN = regexp.MustCompile("^osu file format v(\\d+)$")
	SECTION_PATTERN     = regexp.MustCompile("^\\[([[:alpha:]]+)\\]$")
	KEY_VALUE_PATTERN   = regexp.MustCompile("^([A-Za-z]+)\\s*:\\s*(.*)$")
)

func ParseBeatmap(contents string) (*Beatmap, error) {
	// Largely based on https://github.com/natsukagami/go-osu-parser/blob/master/parser.go
	// TODO: read in a stream
	var section string

	m := &Beatmap{}
	lines := strings.Split(contents, "\n")

	for _, line := range lines {
		line = strings.Trim(line, " \r\n")
		if len(line) == 0 {
			// empty line
			continue
		}

		// check for osu file format header
		if match := FILE_FORMAT_PATTERN.FindStringSubmatch(line); match != nil {
			if n, err := strconv.Atoi(match[1]); err == nil {
				m.Version = n
			}
			continue
		}

		// update current section
		if match := SECTION_PATTERN.FindStringSubmatch(line); match != nil {
			section = match[1]
			continue
		}

		// yay all other sections
		switch strings.ToLower(section) {
		case "general":
			fallthrough
		case "editor":
			fallthrough
		case "metadata":
			fallthrough
		case "difficulty":
			if match := KEY_VALUE_PATTERN.FindStringSubmatch(line); match != nil {
				key, value := match[1], match[2]
				switch strings.ToLower(key) {
				case "audiofilename":
					m.AudioFilename = value
				}
			} else {
				return nil, fmt.Errorf("failed to match: '%+v'", line)
			}
		case "events":
			// TODO
		case "timingpoints":
			// TODO
		case "colours":
			// TODO
		case "hitobjects":
			// TODO
		default:
			return nil, fmt.Errorf("unknown section '%s'", section)
		}
	}

	return nil, fmt.Errorf("%#v", m)
}

func (m *Beatmap) Serialize() (string, error) {
	return "", nil
}

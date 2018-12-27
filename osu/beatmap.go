package osu

import (
	"bufio"
	"errors"
	"fmt"
	"io"
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

var (
	FILE_FORMAT_PATTERN = regexp.MustCompile("^osu file format v(\\d+)$")
	SECTION_PATTERN     = regexp.MustCompile("^\\[([[:alpha:]]+)\\]$")
	KEY_VALUE_PATTERN   = regexp.MustCompile("^([A-Za-z]+)\\s*:\\s*(.*)$")
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
	TimingPoints []*TimingPoint
	HitObjects   []*HitObject

	// TODO: events
}

func ParseBeatmap(reader io.Reader) (*Beatmap, error) {
	// Largely based on https://github.com/natsukagami/go-osu-parser/blob/master/parser.go
	var section string
	var buf []byte
	var err error

	m := &Beatmap{}
	bufreader := bufio.NewReader(reader)

	// compatibility for older versions
	approachSet := false
	artistUnicodeSet := false
	titleUnicodeSet := false

	for ; err == nil; buf, _, err = bufreader.ReadLine() {
		line := strings.Trim(string(buf), " \r\n")
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
					// check that its extension is mp3
					if !strings.HasSuffix(strings.ToLower(value), ".mp3") {
						return nil, errors.New("AudioFilename does not have the .mp3 extension")
					}

					m.AudioFilename = value
				case "audioleadin":
					if val, err := strconv.Atoi(value); err == nil {
						m.AudioLeadIn = val
					}
				case "previewtime":
					if val, err := strconv.Atoi(value); err == nil {
						m.PreviewTime = val
					}
				case "countdown":
					if val, err := strconv.Atoi(value); err == nil {
						m.Countdown = val > 0
					}
				case "sampleset":
					var val int
					switch strings.ToLower(value) {
					case "normal":
						val = SAMPLE_NORMAL
					case "soft":
						val = SAMPLE_SOFT
					case "drum":
						val = SAMPLE_DRUM
					default:
						return nil, fmt.Errorf("unknown sample set '%s'", value)
					}
					m.SampleSet = val
				default:
					// return nil, fmt.Errorf("unknown key '%s'", key)
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

	// compatibility for older versions
	if !approachSet {
		// AR used to be set by OD
		m.ApproachRate = m.OverallDifficulty
	}
	if !artistUnicodeSet {
		m.ArtistUnicode = m.Artist
	}
	if !titleUnicodeSet {
		m.TitleUnicode = m.Title
	}

	return nil, fmt.Errorf("%#v", m)
}

func (m *Beatmap) Serialize(writer io.Writer) error {
	fmt.Fprintf(writer, "osu file format v%d\n", m.Version)
	return nil
}

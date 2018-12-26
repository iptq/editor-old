package osu

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Beatmap struct {
	Version int
}

var (
	FILE_FORMAT_PATTERN = regexp.MustCompile("osu file format v(?P<version>\\d+)")
	SECTION_PATTERN     = regexp.MustCompile("\\[(?P<section>[[:alpha:]]+)\\]")
)

func ParseBeatmap(contents string) (*Beatmap, error) {
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

		fmt.Println("section:", section)
	}

	return nil, fmt.Errorf("%+v", m)
}

func (m *Beatmap) Serialize() (string, error) {
	return "", nil
}

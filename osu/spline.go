package osu

import (
	"strconv"
	"strings"
)

const (
	SPLINE_LINEAR  = 'L'
	SPLINE_PERFECT = 'P'
	SPLINE_BEZIER  = 'B'
	SPLINE_CATMULL = 'C'
)

func ParseControlPoints(line string) (points []IntPoint, err error) {
	pointsStr := strings.Split(line, "|")

	for _, s := range pointsStr {
		var x, y int
		pair := strings.Split(s, ":")

		x, err = strconv.Atoi(pair[0])
		if err != nil {
			return
		}

		y, err = strconv.Atoi(pair[1])
		if err != nil {
			return
		}

		points = append(points, IntPoint{x, y})
	}

	return
}

func SplineFrom(points []IntPoint) {

}

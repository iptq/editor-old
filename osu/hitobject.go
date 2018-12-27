package osu

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/oklog/ulid"
)

type HitObject interface {
	// GetULID returns a unique identifier for this HitObject. This doesn't
	// necessarily need to persist between different instances of the editor
	// (i.e. doesn't need to be saved to disk)
	GetULID() ulid.ULID

	GetStartTime() Timestamp

	Serialize() (string, error)
}

type ObjCircle struct {
	ulid ulid.ULID

	startTime int
}

func ParseHitCircle(params parseParameters, parts []string) (ObjCircle, error) {
	obj := ObjCircle{
		ulid: NewULID(),
	}
	return obj, nil
}

func (obj ObjCircle) GetULID() ulid.ULID {
	return obj.ulid
}

func (obj ObjCircle) GetStartTime() Timestamp {
	return TimestampAbsolute(obj.startTime)
}

func (obj ObjCircle) Serialize() (string, error) {
	// TODO:
	return "", nil
}

type ObjSlider struct {
	ulid ulid.ULID

	startTime int
}

func (obj *ObjSlider) GetULID() ulid.ULID {
	return obj.ulid
}

type ObjSpinner struct {
	ulid ulid.ULID

	startTime int
}

func (obj ObjSpinner) GetULID() ulid.ULID {
	return obj.ulid
}

type parseParameters struct {
	x, y      int
	startTime int
	newCombo  bool
}

func ParseHitObject(line string) (HitObject, error) {
	parts := strings.Split(line, ":")

	x, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil, err
	}

	y, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, err
	}

	startTime, err := strconv.Atoi(parts[2])
	if err != nil {
		return nil, err
	}

	ty, err := strconv.Atoi(parts[3])
	if err != nil {
		return nil, err
	}

	newCombo := (ty & 4) > 0
	params := parseParameters{x, y, startTime, newCombo}

	switch {
	case (ty & 1) > 0:
		return ParseHitCircle(params, parts)
	default:
		return nil, fmt.Errorf("unknown hitobject type: %+v", ty)
	}

}

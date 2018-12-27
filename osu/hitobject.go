package osu

import (
	"errors"
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
	ulid      ulid.ULID
	x, y      int
	startTime Timestamp
	newCombo  bool
	hitsound  Hitsound
}

func ParseHitCircle(params commonParameters, parts []string) (ObjCircle, error) {
	obj := ObjCircle{
		ulid:      NewULID(),
		x:         params.X,
		y:         params.Y,
		startTime: TimestampAbsolute(params.StartTime),
		newCombo:  params.NewCombo,
		hitsound:  Hitsound(params.Hitsound),
	}
	return obj, nil
}

func (obj ObjCircle) GetULID() ulid.ULID {
	return obj.ulid
}

func (obj ObjCircle) GetStartTime() Timestamp {
	return obj.startTime
}

func (obj ObjCircle) Serialize() (string, error) {
	return fmt.Sprintf("%d,%d,%d,%d,%d,%s",
		obj.x,
		obj.y,
		obj.startTime,
		1|(WHAT_THE_FUCK[obj.newCombo]<<2),
		0,
		"0:0:0:0:",
	), nil
}

type ObjSlider struct {
	ulid ulid.ULID

	startTime int
}

func ParseSlider(params commonParameters, parts []string) (ObjSlider, error) {
	obj := ObjSlider{
		ulid: NewULID(),
	}
	return obj, nil
}

func (obj ObjSlider) GetULID() ulid.ULID {
	return obj.ulid
}

func (obj ObjSlider) GetStartTime() Timestamp {
	return TimestampAbsolute(obj.startTime)
}

func (obj ObjSlider) Serialize() (string, error) {
	// TODO:
	return "", errors.New("unimplemented")
}

type ObjSpinner struct {
	ulid ulid.ULID

	startTime int
}

func ParseSpinner(params commonParameters, parts []string) (ObjSpinner, error) {
	obj := ObjSpinner{
		ulid: NewULID(),
	}
	return obj, errors.New("unimplemented")
}

func (obj ObjSpinner) GetULID() ulid.ULID {
	return obj.ulid
}

func (obj ObjSpinner) GetStartTime() Timestamp {
	return TimestampAbsolute(obj.startTime)
}

func (obj ObjSpinner) Serialize() (string, error) {
	// TODO:
	return "", errors.New("unimplemented")
}

type commonParameters struct {
	X, Y      int
	StartTime int
	NewCombo  bool
	Hitsound  int
}

func ParseHitObject(line string) (HitObject, error) {
	parts := strings.Split(line, ",")

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

	hitsound, err := strconv.Atoi(parts[3])
	if err != nil {
		return nil, err
	}

	newCombo := (ty & 4) > 0
	params := commonParameters{x, y, startTime, newCombo, hitsound}

	switch {
	case (ty & 1) > 0:
		return ParseHitCircle(params, parts)
	case (ty & 2) > 0:
		return ParseSlider(params, parts)
	case (ty & 8) > 0:
		return ParseSpinner(params, parts)
	default:
		return nil, fmt.Errorf("unknown hitobject type: %+v", ty)
	}

}

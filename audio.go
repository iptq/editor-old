package editor

import (
	"os"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
)

type AudioManager struct {
	stream beep.StreamSeekCloser
	format beep.Format
}

func (am *AudioManager) Open(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}

	stream, format, _ := mp3.Decode(f)
	am.stream = stream
	am.format = format

	return nil
}

// GetSongLength gets the length of the song in milliseconds
func (am *AudioManager) GetSongLength() int {
	return int(float64(am.stream.Len()) * 1000.0 / float64(am.format.SampleRate))
}

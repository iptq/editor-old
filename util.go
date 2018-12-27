package editor

import (
	"fmt"
)

func FormatTimestamp(milliseconds int) string {
	var minutes, seconds int
	var negSign string

	if milliseconds < 0 {
		milliseconds *= -1
		negSign = "-"
	}

	seconds = milliseconds / 1000
	milliseconds %= 1000
	minutes = seconds / 60
	seconds %= 60

	return fmt.Sprintf("%s%02d:%02d:%03d", negSign, minutes, seconds, milliseconds)
}

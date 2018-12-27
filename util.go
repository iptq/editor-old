package editor

import "fmt"

func FormatTimestamp(milliseconds int) string {
	var minutes, seconds int

	seconds = milliseconds / 1000
	milliseconds %= 1000
	minutes = seconds / 60
	seconds %= 60

	return fmt.Sprintf("%02d:%02d:%03d", minutes, seconds, milliseconds)
}

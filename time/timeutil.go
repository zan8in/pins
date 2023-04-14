package timeutil

import "time"

func Format(format string) string {
	return time.Now().Format(format)
}

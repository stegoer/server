package util

import "time"

// FormatDate returns time formatted as RFC3339 (2022-04-12T13:23:12+50:00)
func FormatDate(t time.Time) string {
	return t.Format(time.RFC3339)
}

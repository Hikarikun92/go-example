package util

import "time"

func TimeToIso(t time.Time) string {
	return t.Format("2006-01-02T15:04:05")
}

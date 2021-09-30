package utils

import "time"

func TimeStamp2Time(tStamp int64) time.Time {
	return time.Unix(tStamp, 0)
}

func Time2TimeStamp(t time.Time) int64 {
	return t.Unix()
}

func ParseDay(tStr string) (time.Time, error) {
	layout := "2006-01-02 15:04:05"
	return time.Parse(layout, tStr)
}

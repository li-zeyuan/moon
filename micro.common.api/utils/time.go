package utils

import "time"

func TimeStamp2Time(tStamp int64) time.Time {
	return time.Unix(tStamp, 0)
}

func Time2TimeStamp(t time.Time) int64 {
	return t.Unix()
}

package utils

import "time"

func TimeStamp2Time(tStamp int64) time.Time {
	return time.Unix(tStamp, 0)
}

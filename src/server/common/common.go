package common

import (
	"time"
)

func GetNowMillisecond() int64 {
	return time.Now().UnixNano() / 1000000
}
func GetNowSecond() int64 {
	return time.Now().Unix()
}

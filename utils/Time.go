package utils

import (
	"time"
)

func CurrentTimeMillis() int64 {
	return time.Now().UnixNano() / 1000000
}

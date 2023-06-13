package time

import (
	"time"
)

func Exec(fn func()) time.Duration {
	startTime := time.Now()
	fn()
	execTime := time.Since(startTime)
	return execTime
}

package httpx

import (
	"time"

	"github.com/guonaihong/gout"
	"github.com/guonaihong/gout/filter"
)

const maxBackOff int = 62

type BackOff func() (next time.Duration, stop bool)

func WithConstInterval(interval time.Duration, maxTimes int) BackOff {
	n := maxTimes
	return func() (next time.Duration, stop bool) {
		if n > 0 {
			n--
			return interval, false
		}
		return 0, true
	}
}

func WithPowerInterval(base time.Duration, maxTimes int) BackOff {
	n := 0
	return func() (next time.Duration, stop bool) {
		if n >= maxBackOff {
			// (1<<63 will be overflow)
			return 0, true
		}
		if n < maxTimes {
			n++
			return base * time.Duration(1<<(n+1)), false
		}
		return 0, true
	}
}

func convertToGoutRetryFunc(fn BackOff) func(c *gout.Context) error {
	return func(c *gout.Context) error {
		next, stop := fn()
		if stop {
			return filter.ErrRetryFail
		}
		time.Sleep(next)
		return filter.ErrRetry
	}
}

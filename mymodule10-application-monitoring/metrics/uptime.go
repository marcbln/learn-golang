package metrics

import (
	"time"
)

type Uptime struct {
	startTime time.Time
}

func (u Uptime) GetValueWithUnit() ValueWithUnit {
	return ValueWithUnit{
		Value: uint64(time.Since(u.startTime).Milliseconds()),
		Unit:  Milliseconds,
	}
}

func NewUpTime(startTime time.Time) Uptime {
	return Uptime{
		startTime: startTime,
	}
}

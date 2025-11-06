package avsyncer

import (
	"time"
)

type Clock struct {
	Rate int
	BaseWall time.Time
	BaseRTP uint32
}

func NewClock(rate int) *Clock {
	return &Clock{
		Rate: rate,
		BaseWall: time.Now(),
		BaseRTP: 0,
	}
}

func (c *Clock) SetBase(wall time.Time, rtp uint32){
	c.BaseWall = wall
	c.BaseRTP = rtp
}
func (c *Clock) RTPToDuration(ts uint32) time.Duration {
	delta := int64(ts) - int64(c.BaseRTP)

	if delta < -(1<<31) {
		delta += 1 << 32
	} else if delta > (1<<31) {
		delta -= 1 << 32
	}

	nanos := delta * 1_000_000_000 / int64(c.Rate)
	return time.Duration(nanos) * time.Nanosecond
}

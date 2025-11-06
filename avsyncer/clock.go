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

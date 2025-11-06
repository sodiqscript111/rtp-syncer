package main

import (
	"fmt"
	"time"
	"rtp-syncer/avsyncer"
)

func main() {
	// 1. Make a video clock (90kHz)
	clock := avsyncer.NewClock(90000)

	// 2. Fake: "At Unix time 1000, RTP timestamp was 9000"
	clock.SetBase(time.Unix(1000, 0), 9000)

	// 3. Print something to confirm
	fmt.Printf("Rate: %d, BaseRTP: %d, BaseWall: %v\n", 
    clock.Rate, clock.BaseRTP, clock.BaseWall)
}

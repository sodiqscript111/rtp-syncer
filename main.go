package main

import (
	"fmt"
	"time"
	"rtp-syncer/avsyncer"
)

func main() {
	// Normal test (9000 ticks later → 100ms)
	clock := avsyncer.NewClock(90000)
	clock.SetBase(time.Unix(1000, 0), 9000)
	dur := clock.RTPToDuration(18000)
	fmt.Printf("Play delay: %v\n", dur)

	// Rollover test: base near max uint32, next packet after wrap
	clock2 := avsyncer.NewClock(90000)
	clock2.SetBase(time.Unix(0, 0), 4_294_967_295) // 2^32 - 1
	durWrap := clock2.RTPToDuration(100)           // after rollover
	fmt.Printf("\nRollover delay: %v\n", durWrap)
}
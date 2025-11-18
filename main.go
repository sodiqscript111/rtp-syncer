package main

import (
	"fmt"
	"rtp-syncer/avsyncer"
	"time"
)

func main() {

	clock := avsyncer.NewClock(90000)
	clock.SetBase(time.Unix(1000, 0), 9000)
	dur := clock.RTPToDuration(18000)
	fmt.Printf("Play delay: %v\n", dur)

	clock2 := avsyncer.NewClock(90000)
	clock2.SetBase(time.Unix(0, 0), 4_294_967_295) // 2^32 - 1
	durWrap := clock2.RTPToDuration(100)           // after rollover
	fmt.Printf("\nRollover delay: %v\n", durWrap)
}

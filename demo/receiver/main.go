package main

import (
	"fmt"
	"log"
	"net"
	"rtp-syncer/avsyncer"
	"time"

	"github.com/pion/rtp"
)

func main() {
	
	addr, _ := net.ResolveUDPAddr("udp", ":5000")
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	fmt.Println("Receiver listening on :5000...")

	clock := avsyncer.NewClock(90000)
	clock.SetBase(time.Now(), 0)
	
	jb := avsyncer.NewJitterBuffer(clock, 200*time.Millisecond)


	go func() {
		buf := make([]byte, 1500)
		for {
			n, _, err := conn.ReadFromUDP(buf)
			if err != nil {
				continue
			}

			
			pkt := &rtp.Packet{}
			if err := pkt.Unmarshal(buf[:n]); err != nil {
				fmt.Println("Error parsing RTP:", err)
				continue
			}


			jb.Push(pkt)
		}
	}()

	
	ticker := time.NewTicker(10 * time.Millisecond)
	defer ticker.Stop()

	fmt.Println("Player started...")

	for range ticker.C {
		pkt := jb.Pop()
		if pkt != nil {
			fmt.Printf("🎥 RENDER: Seq %d (TS: %d)\n", pkt.SequenceNumber, pkt.Timestamp)
		}
	}
}

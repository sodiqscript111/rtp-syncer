package main

import (
	"fmt"
	"math/rand"
	"net"
	"time"

	"github.com/pion/rtp"
)

func main() {

	addr, _ := net.ResolveUDPAddr("udp", "localhost:5000")
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Println("Sender started. Streaming to localhost:5000...")

	seq := uint16(0)
	ts := uint32(0)

	// 30 FPS = 33.33ms per frame
	// 90kHz clock = 3000 ticks per frame
	frameDuration := 33 * time.Millisecond
	tsStep := uint32(3000)

	for {
		// 2. Create Packet
		pkt := &rtp.Packet{
			Header: rtp.Header{
				Version:        2,
				PayloadType:    96,
				SequenceNumber: seq,
				Timestamp:      ts,
			},
			Payload: []byte{0x00, 0x01, 0x02, 0x03}, // Fake video data
		}

		// 3. Marshal
		buf, _ := pkt.Marshal()

		// 4. Send
		conn.Write(buf)
		fmt.Printf("Sent Seq %d\n", seq)

		// 5. Update state
		seq++
		ts += tsStep

		// 6. Sleep with JITTER!
		// Normal sleep is 33ms.
		// We will sleep anywhere from 10ms to 60ms to simulate bad network.
		jitter := time.Duration(rand.Intn(50)-25) * time.Millisecond // +/- 25ms
		time.Sleep(frameDuration + jitter)
	}
}

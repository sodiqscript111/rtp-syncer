package avsyncer

import (
	"container/heap"
	"sync"
	"time"

	"github.com/pion/rtp"
)

type pktItem struct {
	pkt *rtp.Packet
	idx int
}
type JitterBuffer struct {
	lock sync.Mutex
	heap pktHeap

	clock       *Clock
	bufferTime  time.Duration
	lastPopSeq  uint16
	initialized bool
}
type pktHeap []*pktItem

func (h pktHeap) Len() int            { return len(h) }
func (h pktHeap) Less(i, j int) bool  { return h[i].pkt.Timestamp < h[j].pkt.Timestamp }
func (h pktHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *pktHeap) Push(x interface{}) { *h = append(*h, x.(*pktItem)) }
func (h *pktHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func NewJitterBuffer(clock *Clock, bufferTime time.Duration) *JitterBuffer {
	return &JitterBuffer{
		clock:      clock,
		bufferTime: bufferTime,
		heap:       make(pktHeap, 0),
	}
}

func (jb *JitterBuffer) Push(pkt *rtp.Packet) {
	jb.lock.Lock()
	defer jb.lock.Unlock()

	idx := len(jb.heap)
	heap.Push(&jb.heap, &pktItem{pkt: pkt, idx: idx})
}

func (jb *JitterBuffer) Pop() *rtp.Packet {
	jb.lock.Lock()
	defer jb.lock.Unlock()

	if len(jb.heap) == 0 {
		return nil
	}

	item := jb.heap[0]

	offset := jb.clock.RTPToDuration(item.pkt.Timestamp)
	playTime := jb.clock.BaseWall.Add(offset).Add(jb.bufferTime)

	if time.Now().Before(playTime) {
		return nil
	}

	x := heap.Pop(&jb.heap)
	return x.(*pktItem).pkt
}

func (jb *JitterBuffer) GetStats() int {
	jb.lock.Lock()
	defer jb.lock.Unlock()
	return len(jb.heap)
}

package gen

import (
	"errors"
	"sync"
	"time"
)
// Constants for bit lengths
const (
	workerIDBits = 10
	sequenceBits = 12
	maxWorkerID = -1 ^ (-1 << workerIDBits) // 1023
	maxSequence = -1 ^ (-1 << sequenceBits) // 4095

	workerIDShift = sequenceBits
	timeStampShift = sequenceBits + workerIDBits

	customEpoch = int64(1577863800000) // Jan 1 2020 UTC in milliseconds
)
// Gen generates unique IDs
type Gen struct {
	mu 		sync.Mutex
	lastTimestamp	int64
	sequence 	int64
	workerID 	int64
}
// NewGen creates a new ID generator for a given worker ID
func NewGen(workerID int64) (*Gen, error) {
	if workerID < 0 || workerID > maxWorkerID {
		return nil, errors.New("worker ID out of range")
	}
	return &Gen{
		workerID: workerID,
	}, nil
}

// NextID generates the next unique ID
func (g *Gen) NextID() (int64, error) {
	g.mu.Lock()
	defer g.mu.Unlock()

	now := time.Now().UnixMilli()
	if now < g.lastTimestamp {
		return 0, errors.New("clock moved backwards")
	}
	if now == g.lastTimestamp {
		g.sequence = (g.sequence + 1) & maxSequence
		if g.sequence == 0 {
			// Sequence exhausted in this millisecond; wait for next ms.
			now = g.waitUntilNextMillis(now)
		}
	} else {
		g.sequence = 0
	}

	g.lastTimestamp = now

	id := ((now-customEpoch) << timeStampShift) | 
		(g.workerID << workerIDShift) |
		g.sequence
	return id, nil
}

// waitUntilNextMillis spins until the next millisecond
func (g *Gen) waitUntilNextMillis(lastTs int64) int64 {
	for {
		now := time.Now().UnixMilli()
		if now > lastTs {
			return now
		}
	}
}

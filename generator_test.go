package gen

import (
	"sync"
	"testing"
)

func TestGenSequential(t *testing.T) {
	g, err := NewGen(1)
	if err != nil {
		t.Fatal(err)
	}
	var prev int64 = -1
	for i := 0; i < 1000; i++ {
		id, err := g.NextID()
		if err != nil {
			t.Fatal(err)
		}
		if id <= prev {
			t.Errorf("ID not increasing: prev%d current=%d", prev, id)
		}
		prev = id
	}
}

func TestGenConcurrent(t *testing.T){
	g, err := NewGen(1)
	if err != nil {
		t.Fatal(err)
	}
	const numIDs = 100000	
	ids := make(map[int64]struct{})
	mu := sync.Mutex{}

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < numIDs/10; j++{
				id, err := g.NextID()
				if err != nil {
					t.Fatal(err)
				}
				mu.Lock()
				if _, exists := ids[id]; exists {
					t.Errorf("Duplicate ID found: %d", id)
				}
				ids[id] = struct{}{}
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
}

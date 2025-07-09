package main

// concurrently testing id generator
import (
	"fmt"
	"log"
	"sync"
	"time"

	gen "github.com/RicardoTlatelpa/uniqueidgen"
)

func main() {
	g, err := gen.NewGen(1)
	if err != nil {
		log.Fatal(err)
	}

	const totalIDs = 1_000_000
	const concurrency = 10
	start := time.Now()

	var wg sync.WaitGroup
	wg.Add(concurrency)
	
	for i := 0; i < concurrency; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < totalIDs/concurrency; j++ {
				_, err := g.NextID()
				if err != nil {
					log.Fatal(err)
				}
		}	
	}()
}

	wg.Wait()
	elapsed := time.Since(start)
	idsPerSec := float64(totalIDs) / elapsed.Seconds()

	fmt.Printf("Generated %d IDs in %s\n", totalIDs, elapsed)
	fmt.Printf("Throughput: %.2f IDs/second\n", idsPerSec)
}

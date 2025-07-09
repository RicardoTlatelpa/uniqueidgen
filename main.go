package main

import (
	"fmt"
	"log"
	"time"

	"github.com/RicardoTlatelpa/uniqueidgen/gen"
)

func main() {
	// create a new generator with worker ID 1

	g, err := gen.NewGen(1)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 10; i++ {
		id, err := g.NextID()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id)
		time.Sleep(1 * time.Millisecond)
	}
}

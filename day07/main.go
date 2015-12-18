package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	c, err := NewCircuit(file)
	if err != nil {
		log.Fatal(err)
	}

	start := time.Now()
	val := c.Read("a")
	fmt.Printf("a is %d (%s)\n", val, time.Since(start))
}

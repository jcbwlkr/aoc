package main

import (
	"fmt"
	"time"
)

func main() {
	var input = "yzbqklnj"

	run(mineOriginal, "original serial implementation", input)
	run(mineSerial, "improved serial implementation", input)
	run(mineConcurrent, "concurrent implementation", input)
}

type miner func(string, string) int

func run(fn miner, label, input string) {
	var (
		t0    = time.Now()
		coin  int
		start time.Time
	)

	fmt.Println("Starting", label)

	fmt.Print("5-zero coin: ")
	start = time.Now()
	coin = fn(input, "00000")
	fmt.Printf("%d (%v)\n", coin, time.Since(start))

	fmt.Print("6-zero coin: ")
	start = time.Now()
	coin = fn(input, "000000")
	fmt.Printf("%d (%v)\n", coin, time.Since(start))

	//fmt.Print("Failure case: ")
	//start = time.Now()
	//coin = fn(input, "xxxx")
	//fmt.Printf("%d (%v)\n", coin, time.Since(start))

	fmt.Printf("Total time: %v\n\n", time.Since(t0))
}

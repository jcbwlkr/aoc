package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("AOC Day 10")

	run("Part 1", 40)
	run("Part 2", 50)
}

func run(label string, iterations int) {
	var (
		start    time.Time
		duration time.Duration
		val      int
		input    = "1113122113"
	)

	fmt.Printf("%s: %d iterations: ", label, iterations)
	start = time.Now()
	val = len(lookAndSay(input, iterations))
	duration = time.Since(start)
	fmt.Println(val, duration)
}

package main

import (
	"fmt"
	"log"
	"time"
)

type logger struct{}

func (logger) Println(args ...interface{}) {
	log.Println(args...)
}

func main() {
	t0 := time.Now()
	fmt.Println("AOC Day 10 Part 2")
	fmt.Println("50 iterations: ")
	fmt.Println(len(lookAndSay("1113122113", 50, logger{})))
	fmt.Println("That took", time.Since(t0))
}

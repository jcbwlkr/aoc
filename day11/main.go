package main

import "fmt"

func main() {
	var pw = "vzbxkghb"

	fmt.Println("AOC Day 11")
	fmt.Println("Starting with password", pw)
	for validCount := 0; validCount < 2; {
		pw = increment(pw)
		if isValid(pw) {
			validCount++
			fmt.Println("Next valid pw is", pw)
		}
	}
}

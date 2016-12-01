package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
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

	a := c.Read("a")
	fmt.Printf("Part 1: a is %d\n", a)

	c.Replace("b", strconv.Itoa(int(a)))

	a = c.Read("a")
	fmt.Printf("Part 2: a is %d\n", a)
}

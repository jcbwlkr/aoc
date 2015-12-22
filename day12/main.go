package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
)

func main() {
	nums := regexp.MustCompile("([-0-9]+)")

	input, err := ioutil.ReadFile("input.json")
	if err != nil {
		log.Fatal(err)
	}

	var sum int
	matches := nums.FindAllStringSubmatch(string(input), -1)
	fmt.Println(matches)
	for _, m := range matches {
		for i := 1; i < len(m); i++ {
			n, err := strconv.Atoi(m[i])
			if err != nil {
				log.Fatal(err)
			}
			sum += n
		}
	}

	fmt.Println(sum)
}

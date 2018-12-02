package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:	day03 <target>")
		os.Exit(0)
	}
	target := Atoi(os.Args[1])
	fmt.Println("== Part 1 ==")
	part1(target)
	fmt.Println("== Part 2 ==")
	part2(target)
}

func part1(target int) {
	idx := 1
	lyrSz := 1
	sizeSquare := int(math.Pow(float64(lyrSz), 2))
	for sizeSquare < target {
		lyrSz += 2
		idx += 1
		sizeSquare = int(math.Pow(float64(lyrSz), 2))
	}
	diff := sizeSquare - target
	lyrJmp := lyrSz - 1
	// Find the side it's on
	side := int(diff / lyrJmp)
	perps := make([]int, 4)
	perps[0] = sizeSquare - (lyrJmp / 2)
	perps[1] = perps[0] - lyrJmp
	perps[2] = perps[1] - lyrJmp
	perps[3] = perps[2] - lyrJmp
	fmt.Println("Distance: ", int(math.Abs(float64(perps[side]-target)))+(idx-1))
}

const (
	DIR_E = iota
	DIR_N
	DIR_W
	DIR_S
	DIR_ERROR
)

var memMap map[string]int
var lastX, lastY int
var lastDir int

func part2(target int) {
	lastDir = DIR_S
	lastX, lastY = 0, 0

	val := 1
	memMap = make(map[string]int)
	memMap[getMapKey(lastX, lastY)] = val
	for val <= target {
		// Find the coordinate for the next one
		lastX, lastY, lastDir = findNextMapPos()
		val = getAdjacentSum(lastX, lastY)
		memMap[getMapKey(lastX, lastY)] = val
	}
	fmt.Println("Result:", val, "at", getMapKey(lastX, lastY))
}

func getAdjacentSum(posX, posY int) int {
	res := 0
	for _, x := range []int{-1, 0, 1} {
		for _, y := range []int{-1, 0, 1} {
			if x == 0 && y == 0 {
				continue
			}
			if i, ok := memMap[getMapKey((posX+x), (posY+y))]; ok {
				res += i
			}
		}
	}
	return res
}

func findNextMapPos() (int, int, int) {
	// First check if the current pos is empty
	if _, ok := memMap[getMapKey(lastX, lastY)]; !ok {
		return lastX, lastY, lastDir
	}
	// Check if the wrapping direction is free
	nextDir := (lastDir + 1) % DIR_ERROR
	testX, testY := getXYInDir(nextDir)
	if _, ok := memMap[getMapKey(testX, testY)]; ok {
		testX, testY = getXYInDir(lastDir)
		nextDir = lastDir
	}
	// Otherwise continue in lastDir
	return testX, testY, nextDir
}

func getXYInDir(dir int) (int, int) {
	testX, testY := lastX, lastY
	switch dir {
	case DIR_E:
		testX++
	case DIR_N:
		testY--
	case DIR_W:
		testX--
	case DIR_S:
		testY++
	}
	return testX, testY
}

func getMapKey(xPos, yPos int) string {
	return fmt.Sprintf("%d;%d", xPos, yPos)
}

func Atoi(i string) int {
	var ret int
	var err error
	if ret, err = strconv.Atoi(i); err != nil {
		log.Fatal("Invalid Atoi")
	}
	return ret
}
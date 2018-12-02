package main

import "fmt"

var banks = [...]int{4, 10, 4, 1, 8, 4, 9, 14, 5, 1, 14, 15, 0, 15, 3, 5}

//var banks = [...]int{0, 2, 7, 0}

var steps = make(map[string]struct{})

func main() {
	fmt.Println("Balancing", len(banks), "banks")

	fmt.Println("Starting:", banks)
	for {
		balance()
		fmt.Println("Balanced:", banks)
		if _, seen := steps[fmt.Sprint(banks)]; seen {
			break
		}
		steps[fmt.Sprint(banks)] = struct{}{}
	}

	fmt.Println(len(steps) + 1)
}

func balance() {

	// Identify the target with the most blocks
	var blocks, pos int
	for i, v := range banks {
		if v > blocks {
			pos = i
			blocks = v
		}
	}

	// Remove blocks from the target then redistribute equal shares to the rest
	banks[pos] = 0
	share := blocks / (len(banks) - 1)
	if share == 0 {
		share = 1
	}

	//fmt.Printf("Distributing the %2d blocks from %2d to others %2d at a time: ", blocks, pos, share)

	// Start redistributing from pos forward
	for i := pos + 1; i < len(banks) && blocks > 0; i++ {
		banks[i] += share
		blocks -= share
	}
	// Wrap around and redistribute from 0 up to pos
	for i := 0; i < pos && blocks > 0; i++ {
		banks[i] += share
		blocks -= share
	}

	// If there are any blocks left they go back to the target cell
	banks[pos] += blocks
}

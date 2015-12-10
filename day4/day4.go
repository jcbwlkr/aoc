package main

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math"
	"strconv"
	"strings"
	"sync"
	"time"
)

func main() {
	var input = "ckczppom"

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

func mineOriginal(input string, prefix string) int {
	for i := 1; i < math.MaxInt32; i++ {
		shaft := input + fmt.Sprintf("%d", i)
		ore := fmt.Sprintf("%x", md5.Sum([]byte(shaft)))

		if strings.HasPrefix(ore, prefix) {
			return i
		}
	}

	return -1
}

func mineSerial(input string, prefix string) int {
	var (
		// l is the min length of the md5 byte slice that is needed to generate a
		// hex string at least as long as the desired prefix
		l = int(math.Ceil(float64(len(prefix)) / float64(2)))

		// in is the input cast as a []byte. Do this before the loop so we don't
		// have to keep doing it.
		in = []byte(input)

		p = []byte(prefix)

		ore = make([]byte, hex.EncodedLen(l))
	)

	for i := 1; i < math.MaxInt32; i++ {
		// Combine the bytes of in plus the bytes of the string version of i then
		// get the md5 sum as another []byte
		sum := md5.Sum(append(
			in,
			[]byte(strconv.Itoa(i))...,
		))

		// Get the hex encoding of just the first part of the md5 sum slice
		hex.Encode(ore, sum[:l])

		if bytes.Equal(p, ore[:len(p)]) {
			return i
		}
	}

	return -1
}

func mineConcurrent(input string, prefix string) int {
	var (
		workers = 8                              // how many workers to start
		work    = make(chan int, workers*100000) // chan where feeder provides workers input
		scrap   = make(chan int, workers*100000) // chan where workers will report failure
		jackpot = make(chan int, 1)              // chan where workers will report success
		done    = make(chan struct{})            // signal chan to stop all goroutines
		wg      sync.WaitGroup                   // waitgroup workers will close when done

		in = []byte(input) // pre-do the work of casting input to []byte
		p  = []byte(prefix)
		l  = int(math.Ceil(float64(len(prefix)) / float64(2))) // min byte length needed to make sufficient hex
	)

	// Kick off workers to mine for the coin
	wg.Add(workers)
	for j := 0; j < workers; j++ {
		go func() {
			ore := make([]byte, hex.EncodedLen(l))
			for {
				select {
				case <-done:
					wg.Done()
					return
				case i := <-work:
					sum := md5.Sum(append(in, []byte(strconv.Itoa(i))...))
					hex.Encode(ore, sum[:l])

					if bytes.Equal(p, ore[:len(p)]) {
						jackpot <- i
					} else {
						scrap <- i
					}
				}
			}
		}()
	}

	// Kick off a goroutine to feed prospects to our workers as they're available
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i < math.MaxInt32; i++ {
			select {
			case <-done:
				return
			case work <- i:
			}
		}
	}()

	// Kick off goroutine to look for the failure case which is where all the
	// prospects were examined and none of them worked
	wg.Add(1)
	go func() {
		scrapFound := 0
		for {
			select {
			case <-done:
				wg.Done()
				return
			case <-scrap:
				scrapFound++
				if scrapFound == math.MaxInt32-1 {
					close(jackpot)
				}
			}
		}
	}()

	// Pull the first value we find off jackpot then signal workers to quit by
	// closing `done`.
	gold, found := <-jackpot
	close(done)

	// Wait for all goroutines to quit so we can be sure we can be sure to clean
	// up after ourselves.
	wg.Wait()

	if found {
		return gold
	}
	return -1
}

package main

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"math"
	"strconv"
	"sync"
)

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

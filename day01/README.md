# AoC Day 1
## Not Quite Lisp

This puzzle feeds a stream of `(` and `)` values to represent instructions for
Santa to go up or down a floor respectively in an infinitely tall apartment
building. He starts on floor 0.

## Objectives
* Part 1: Find the floor Santa is on at the end of the instructions
* Part 2: Find the position of the first instruction that brings Santa to the
  first basement level (-1).

## Answers For My Input
* Part 1: 138
* Part 2: 1771

## My Approach
I wrote separate functions for parts 1 and 2: `floorCalculator` and
`basementFinder`. For each func I consume the interface `io.Reader` which
allowed me to pass an `os.File` from my main goroutine and a
`strings.NewReader` from my tests. Inside the functions I use a `bufio.Scanner`
to inspect the contents of the reader.

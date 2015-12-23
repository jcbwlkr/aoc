# AoC Day 4
## The Ideal Stocking Stuffer

Santa needs help mining some AdventCoins (very similar to bitcoins) to use as
gifts for all the economically forward-thinking little girls and boys.

To do this, he needs to find MD5 hashes which, in hexadecimal, start with at
least five zeroes. The input to the MD5 hash is some secret key (your puzzle
input, given below) followed by a number in decimal. To mine AdventCoins, you
must find Santa the lowest positive number (no leading zeroes: 1, 2, 3, ...)
that produces such a hash.

## Objectives
* Part 1: Find the first value that generates a 5-zero coin.
* Part 2: Find the first value that generates a 6-zero coin.

## Answers For My Input
* Part 1: 282749
* Part 2: 9962624

## My Approach

This was actually the first exercise that got me interested in AoC. A friend
posted his implementation which is included here as `mineOriginal`. It worked
but it ran for about 4.8 seconds on my machine and I was curious to see if we
could do better. Working with friends and coworkers I came up with two more
implementations: an improved serial method `mineSerial` and a concurrent
implementation `mineConcurrent`. In both cases I maintained the signature of
`mineOriginal` and the behavior of returning -1 if a match was not find between
1 and `math.MaxInt32`.

The serial implementation clocks in at around 1.1 seconds. Speed improvements
mostly came from avoiding `fmt.Sprintf` for conversions and instead using the
`strconv` and `hex` packages. Additionally I avoided casting to strings instead
preferring operating on byte slices. Finally some speed was realized by only
hex encoding just enough of the md5 bytes to check for the desired prefix.

I had a theory that I could get even faster results by doing the work in
parallel. `mineConcurrent` kicks off 8 goroutines to mine for coins, 1 to feed
them input, 1 to report failure if the workers and feeder come up empty, and
then the main goroutine reports the results. These are all orchestrated with
channels and I use a `sync.WaitGroup` to ensure I clean up after myself before
returning the answer. As easy as Go makes concurrency this implementation still
is slower than `mineSerial` (around 3.0 seconds) because of the costs involved
in orchestration.

This problem was a fun experience as I learned a lot about concurrency,
profiling, and benchmarking.

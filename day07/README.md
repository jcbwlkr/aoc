# AoC Day 7
## Some Assembly Required

This project provides an input file of instructions which represent a circuit
of wires which are connected as bitwise logic gates. Each instruction
represents sending a value on a particular wire or connecting two or more wires
with some bitwise operation.


## Objectives
* Part 1: Parse the circuit then determine the value of wire `a`.
* Part 2: Set the result from part 1 as the value on wire `b` then recalculate
  the value of wire `a`.

## Answers For My Input
* Part 1: 956
* Part 2: 40149

## My Approach

I created a type `Circuit` which holds each of the
instructions in a `map[string]string`.  Wrapping my head
around this was much easier when I realized that each wire
is only written to once. Once the entire circuit is
assembled I recursively work backwards from the desired
wire `a` until I find it's value. This was extremely slow
at first until I implemented a cache of values for each
wire once they were calculated. Adding the cache was easy
by using `defer`.

Implementing part 2 was easy because I added a function to
replace a particular wire's value (which also reset the
cache).

Like the other days I implemented this using TDD and
worked around the `io.Reader` interface.

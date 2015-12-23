# AoC Day 6
## Probably a Fire Hazard

Imagine we have a grid of one million christmas lights arranged in a 1000x1000
square. Instructions are read from a file in the form of

* `turn on 123,456 through 456,789`
* `turn off 400,400 through 600,600`
* `toggle 450,50 through 550,950`

## Objectives
* Part 1: Count the number of lights on at the end.
* Part 2: Interpret the commands to increase/decrease brightness of the lights
  measure as an int. Sum the total brightness of all lights at the end.

## Answers For My Input
* Part 1: 569999
* Part 2: 17836115

## My Approach

I had a bit of fun with this one. In addition to generating the final answers
for the puzzle I generated animated gifs representing the steps in the
instructions. I also experimented with adding methods to non-struct types such
as `type Board map[Point]Light`.

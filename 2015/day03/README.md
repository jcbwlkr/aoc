# AoC Day 3
## Perfectly Spherical Houses in a Vacuum

Santa is delivering presents to an infinite two-dimensional grid of houses.
Directions are fed in as one house to the north (`^`), south (`v`), east (`>`),
or west (`<`).

## Objectives
* Part 1: How many houses are visited at least once.
* Part 2: Every other command is given to a robot Santa. Now how many houses
  are visited at least once.

## Answers For My Input
* Part 1: 2081
* Part 2: 2341

## My Approach
I created a type `House` to represent a point on the grid and a
type `Worker` that tracked its position and knew how to respond to
a direction to move (a rune). I then used the `ReadRune` method
from `bufio.Reader` to feed the stream of commands to one or more
workers. A `map[House]int` tracked the number visits.

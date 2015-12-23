# AoC Day 2
## I Was Told There Would Be No Math

Given a list of dimensions for present boxes the elves want to calculate the
exact amount of wrapping paper and ribbon needed. Formulas were provided for
calculating both.

## Objectives
* Part 1: Find the total square feet of wrapping paper needed.
* Part 2: Find the required number of feet of ribbon.

## Answers For My Input
* Part 1: 1588178
* Part 2: 3783758

## My Approach

I chose to create a type `Box` that implements the calculations for the
required paper and ribbons. I also created a constructor `NewBoxFromDimensions`
that handles the work of turning `2x3x4` into a `Box` with the appropriate
dimensions.

On top of that I created a type `BoxScanner` that wraps a `bufio.Scanner` for
reading the input file. It can be seen in action in `main.go`.

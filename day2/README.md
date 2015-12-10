# AOC Day 2

I chose to create a type `Box` that implements the calculations for the
required paper and ribbons. I also created a constructor `NewBoxFromDimensions`
that handles the work of turning `2x3x4` into a `Box` with the appropriate
dimensions.

On top of that I created a type `BoxScanner` that wraps a `bufio.Scanner` for
reading the input file. It can be seen in action in `main.go`.

Throughout development I tried to write tests first for everything I was
making.

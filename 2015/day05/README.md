# AoC Day 5
## Doesn't He Have Intern-Elves For This

Santa has a set of weirdly arbitrary rules for designating a string as
`naughty` or `nice` such as if it contains at least three vowels, a pair of
letters that appears twice etc.

## Objectives
* Part 1: Using one set of rules count the number of nice strings.
* Part 2: Using a different set of rules count the nice strings.

## Answers For My Input
* Part 1: 255
* Part 2: 55

## My Approach

I wrote a func for each rule set `isNiceOld` and `isNiceNew`.
These funcs were easy examples for writing table driven tests as
they both have the signature `func (string) bool`. My `func main`
just scanned over the input using `bufio.Scanner` and counted the
results from each rule set.

For `isNiceOld` I used the rune generating behavior of
`range`ing over a string. Within the loop it was easy to check for
the rules and return a result.

For `isNiceNew` I wanted more precise control when looping through
the runes so I cast the string to a `[]rune` and used the classic
form of `for`. To look for a repeated pair used slice operations
to look for the current pair in a new slice representing the rest
of the string.

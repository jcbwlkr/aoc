# AoC Day 13
## Knights of the Dinner Table
You are given a list of guests for a party and their opinions of all other
guests measured in happiness points.

## Objectives
* Part 1: Find the optimal seating arrangement that generates the most
  happiness.
* Part 2: Add yourself to the list assuming you have 0 opinions of everyone and
  they have 0 opinion of you.

## Answers For My Input
* Part 1: 664
* Part 2: 640

## My Approach
My input only has 8 unique guests so I decided to just brute force it. I
created a type `Guest` and a type `Table` which is just a `[]Guest`. I then
added the methods `TotalHappiness` and `Optimize` to `Table`. The `Optimize`
method simply generates all combinations of `Table` and tests the happiness of
each.

The hardest part was actually in my test code which compares if two tables are
equal since I don't care about the actual position of a guest in the slice I
only care about their relative position (who is sitting on their right and
left).

Part 2 was trivial to implement because I was able to leverage the zero value
behavior of maps. I just appended myself to the table and recalled `Optimize`.

```
	tbl = append(tbl, Guest{Name: "Host"})
	tbl.Optimize()
```

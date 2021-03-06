defmodule Freq do
  # Walks the whole list of inputs and returns the final value.
  def part1([head | tail], freq) do
    part1(tail, freq + atoi(head))
  end

  # Recursive base case, returns current number when the list is empty []
  def part1([], freq), do: freq

  # Walks the list as many times as necessary until it sees a repeated frequency
  def part2(list, freq, seen) do
    case p2Walk(list, freq, seen) do
      # If I get 2 values then walk the whole list again
      {val, seen} -> part2(list, val, seen)
      # If I get 1 value then it's the answer so return it
      val -> val
    end
  end

  # Walks the list of inputs. Returns either
  # {final_frequency} -> if it finds a repeated frequency
  # {current_frequency, seen_frequencies} -> if it didn't find a repeat
  def p2Walk([head | tail], freq, seen) do
    freq = freq + atoi(head)

    case MapSet.member?(seen, freq) do
      true -> freq
      false -> p2Walk(tail, freq, MapSet.put(seen, freq))
    end
  end

  # Recursive base case, returns {current number, seen_frequencies} when the list is empty []
  def p2Walk([], freq, seen), do: {freq, seen}

  # atoi returns the int value of s. Invalid values are just treated as 0.
  def atoi(s) do
    case Integer.atoi(s) do
      {x, _} -> x
      :error -> 0
    end
  end
end

################################################################################
# Execution starts
################################################################################
input = File.read!("input.txt")
lines = String.split(String.trim(input), "\n")

num = Freq.part1(lines, 0)

IO.puts("Part #1")
IO.puts("Final frequency: #{num}")

# Try again using "reduce" just for fun
num = Enum.reduce(lines, 0, fn val, freq -> Freq.atoi(val) + freq end)

IO.puts("Part #1 alternate:")
IO.puts("Final frequency: #{num}")

# Create MapSet of seen frequencies. Put "0" in the set to begin with.
seen = MapSet.new()
seen = MapSet.put(seen, 0)

num = Freq.part2(lines, 0, seen)

IO.puts("Part #2")
IO.puts("First repeated frequency: #{num}")

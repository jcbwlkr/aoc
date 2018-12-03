defmodule Freq do
  # Walks the whole list of inputs and returns the final value.
  def part1([head | tail], freq) do
    part1(tail, freq + parse(head))
  end

  # Recursive base case, returns current number when the list is empty []
  def part1([], freq), do: freq

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
  # {current_frequency, map_of_seen_frequencies} -> if it didn't find a repeat
  def p2Walk([head | tail], freq, seen) do
    freq = freq + parse(head)

    case seen[freq] do
      true -> freq
      nil -> p2Walk(tail, freq, Map.put(seen, freq, true))
    end
  end

  # Recursive base case, returns {current number, map} when the list is empty []
  def p2Walk([], freq, seen), do: {freq, seen}

  # parse returns the int value of s. Invalid values are just treated as 0.
  def parse(s) do
    case Integer.parse(s) do
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
IO.puts("Final value: #{num}")

num = Freq.part2(lines, 0, %{0 => true})

IO.puts("Part #2")
IO.puts("Final value: #{num}")

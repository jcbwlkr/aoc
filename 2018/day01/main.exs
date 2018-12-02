defmodule Freq do

  # Walks the whole list of inputs and returns the final value.
  def part1([head | tail], n) do

    # Parse string to int. Doesn't handle errors?
    val = String.to_integer head

    # Process rest of list
    part1(tail, n + val)
  end

  # Recursive base case, returns current number when the list is empty []
  def part1([], n), do: n
end

################################################################################
# Execution starts
################################################################################
input = File.read! "input.txt"
lines = String.split String.trim(input), "\n"

num = Freq.part1 lines,  0

IO.puts "Part #1"
IO.puts "Final value: #{num}"

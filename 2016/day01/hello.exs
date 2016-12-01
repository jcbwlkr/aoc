defmodule Elf do

  # Walks the whole list of directions and returns the final position.
  def walk([head | tail], facing, x, y) do
    # Get new direction and distance to travel
    {direction, distance} = parse head

    # Turn the new direction and move that distance
    facing = turn direction, facing
    {x, y} = move facing, x, y, distance

    # Process rest of list
    walk(tail, facing, x, y)
  end

  # Recursive base case, returns current info when the list is empty []
  def walk([], facing, x, y), do: {facing, x, y}

  # Parses an input string like "R23" into the direction and distance to travel
  def parse(input) do
    {distance, _} = Integer.parse(String.slice input, 1..-1)
    direction = case String.first input do
      "R" -> :right
      "L" -> :left
    end

    {direction, distance}
  end

  # Turns right or left returning the new direction
  def turn(:right, :north), do: :east
  def turn(:right, :south), do: :west
  def turn(:right, :east), do: :south
  def turn(:right, :west), do: :north
  def turn(:left, :north), do: :west
  def turn(:left, :south), do: :east
  def turn(:left, :east), do: :north
  def turn(:left, :west), do: :south

  # Moves the elf len spaces. Returns new x and y
  def move(:north, x, y, len), do: {x, y+len}
  def move(:south, x, y, len), do: {x, y-len}
  def move(:east, x, y, len), do: {x+len, y}
  def move(:west, x, y, len), do: {x-len, y}
end

################################################################################
# Execution starts
################################################################################
input = File.read! "input.txt"
directions = String.split input, ", "

{facing, x, y} = Elf.walk directions, :north, 0, 0
distance = Kernel.abs(x) + Kernel.abs(y)

IO.puts "Part #1"
IO.puts "Final position: facing #{facing} at x #{x}, y #{y}"
IO.puts "That is #{distance} blocks from the starting position"

IO.puts ""
IO.puts "Part #2"

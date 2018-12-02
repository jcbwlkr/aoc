#!/usr/bin/env bash

p1Valid=0;
p2Valid=0;
while read ln; do
  # Look for dupes.
  echo "${ln}" | tr " " "\n" | sort | uniq -c | grep -qv "\b1 " || p1Valid=$((p1Valid+1))

  # Look for anagrams. Sort each word's letters alphabetically then join them
  # back together and look for dupes.
  sorted=$(for word in $(echo "${ln}" | tr " " "\n"); do echo $word | fold -w1 | sort | paste -s -d \\0 -; done;)
  echo "${sorted}" | tr " " "\n" | sort | uniq -c | grep -qv "\b1 " || p2Valid=$((p2Valid+1))
done < input.txt

echo "There are ${p1Valid} passwords valid in part 1"
echo "There are ${p2Valid} passwords valid in part 2"

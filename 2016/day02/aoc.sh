#!/bin/bash

# Part 1: 19636
# Part 2: 3CC43

set -e

if [ "$#" -ne 4 ]; then
    echo "This program requires 4 arguments"
    echo " 1. The board file"
    echo " 2. The starting line number (0 based)"
    echo " 3. The starting column number (0 based)"
    echo " 4. The directions input file"
    exit 1
fi

# Read the board file into an array, one entry per line
board=()
i=0
for l in $(cat $1); do
    board[$i]=$l
    i=$((i+1))
done

# Define initial position on the board
pos=($2 $3)
combo=""

# Read the directions one line at a time and move the pos around the board.
for instructions in $(cat $4); do
    while [[ "${instructions}" != "" ]]; do
        # Chop the first char off $instructions
        dir=${instructions:0:1}
        instructions=${instructions:1}

        # Get current position
        ln=${pos[0]}
        col=${pos[1]}

        # Determine new position
        case "${dir}" in
            "U")
                ((ln--)) || :
                ;;
            "D")
                ((ln++)) || :
                ;;
            "L")
                ((col--)) || :
                ;;
            "R")
                ((col++)) || :
                ;;
        esac

        if [ $ln -ge 0 -a $col -ge 0 ]; then
            row=${board[$ln]}
            num=${row:$col:1}
            # Validate the number isn't blank. Remove literal dots since
            # they're placeholders for blanks.
            if [ "${num//./}" != "" ]; then
                pos[0]=$ln
                pos[1]=$col
            fi
        fi
    done

    # Update combo with result of this line
    row=${board[${pos[0]}]}
    num=${row:${pos[1]}:1}
    combo=${combo}${num}
done;

echo "Combination: ${combo}"

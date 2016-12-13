# Get all the lines. We do this so we can preallocate a long enough space
lns <- readLines("input.txt")

# Figure out how many columns there are in the input. This is how many rows
# we'll need in our list.
height <- nchar(lns[1])
# The number of lines in the input is the width of each row in our list.
width <- length(lns)

# Allocate the space
chars <- vector("list", height)
for (i in seq(1, height)) {
    chars[[i]] <- vector("character", width)
}

# Read the input into our new structure
for (i in seq_along(lns)) {
    ln <- lns[[i]]
    for (j in seq(1, nchar(ln))) {
        chars[[j]][[i]] <- substr(ln, j, j)
    }
}

word <- ""
for (i in seq_along(chars)) {
    most <- names(which.max(table(chars[[i]])))
    word <- paste(word, most, sep="")
}

print(word)

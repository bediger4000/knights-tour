# Knight's Tour

[Wikipedia article](https://en.wikipedia.org/wiki/Knight%27s_tour)

An implementation of the "knight's tour" puzzle.

## Build and Run

This is a simple Go program with no dependencies outside the
standard Go packages.

```sh
$ go build tour1.go
$ ./tour1 5 | sort | uniq | wc -l
1728
$
```

## Program Design

The program pre-calculates the `x,y` positions of every legal
move for each square on a board.
Boards can be arbitrarily-sized.

The program iteratively puts a knight on each square of the board,
the recursively finds all the tours that start on that square.
The recursive function recieves the board so far,
and the `x,y` position of the knight on that board.
Squares of the board have a boolean annotation for whether or not
that square has been visited.

## Analysis

# Knight's Tour

[Wikipedia article](https://en.wikipedia.org/wiki/Knight%27s_tour)

An implementation of the "knight's tour" puzzle.

## Build and Run

This is a simple Go program with no dependencies outside the
standard Go packages.

You invoke it with a single number,
the size of a side of the chessboard.
It prints tours to stdout.

```sh
$ go build tour1.go
$ ./tour1 5 | sort | uniq | wc -l
1728
$
```

A 5-sided board is the only one that can be exhaustively
searched for tours in any reasonable amount of time.

The program prints out every tour it discovers in Go-fmt-package "%v" mode.
If you want the unique tours, you have to filter the program's output.

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

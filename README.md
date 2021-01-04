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

This is a brute-force backtracking algorithm that potentially
finds all knight's tours of an N-by-N chess board.
It does nothing fancy to avoid duplicate,
rotationally symmetric, or mirror image tours.

The program pre-calculates the `x,y` positions of every legal
move for each square on a board.
Boards can be arbitrarily-sized.

The program iteratively puts the knight on each square of the board,
then recursively finds all the tours that start on that square.
The recursive function receives the board so far,
and the `x,y` position of the knight on that board.
Squares of the board have a boolean annotation for whether or not
that square has been visited.
If the iteration reaches a depth equal to the number of squares
on the board, the program found a tour.

## Analysis

I did this because my son was learning Python,
and a "Knight's Tour" came up in his class.
It seems pretty hard for a beginner-level problem,
particularly if you do a variable size-of-board.

Knight's tours aren't really enumerable for any worthwhile size of board.
I mean, there are 1728 unique tours on a 5-by-5 board.

This might be a worthwhile problem to do to educate yourself
on backtracking solutions.

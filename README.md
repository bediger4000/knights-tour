# Knight's Tour

[Wikipedia article](https://en.wikipedia.org/wiki/Knight%27s_tour)

An implementation of the "knight's tour" puzzle.

### Daily Coding Problem: Problem #962 [Hard]

This problem was asked by Google.

A knight's tour is a sequence of moves by a knight on a chessboard such
that all squares are visited once.

Given N,
write a function to return the number of knight's tours on an N by N chessboard.

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

## Interview Analysis

Since knight's tour came up as a Daily Coding Problem,
this seems like a take home type of question,
probably for a senior developer type job.
It's definitely too hard to do as a whiteboard problem in front of interviewers.

There's probably a lot of ways to solve this,
so the interviewer could look at nuance if the candidate gets a solution.
If it's a take-home-problem, the interviewer should be prepared to see
google-able solutions. Knight's tour is a really old problem.
Perhaps the interviewer could use the candidate's solution as a jumping-of-point
for discussion of program design, or pecularities of the programming language used,
or design decisions.

The [Daily Coding Problem book](https://www.amazon.com/Daily-Coding-Problem-exceptionally-interviews/dp/1793296634)
could have used this as an example in Chapter 14, *Backtracking*.
You can actually get solutions for smaller boards,
where their example of "solving sudoku efficiently" isn't efficient,
and probably doesn't find a solution in anything like reasonable time.

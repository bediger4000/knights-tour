package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

// for a N-square-per-side board, with <0,0> in one corner,
// and <N,N> in the diagonally opposite corner, a knight can
// make moves to squares with these <delta-x,delta-y> increments
// to its coordinates. Some <x,y> positions might have incremented
// coordinates that are "off the board" one side or another.
var moveIncrements = [][2]int{
	{-1, 2},
	{1, 2},
	{-2, -1},
	{-2, 1},
	{2, -1},
	{2, 1},
	{-1, -2},
	{1, -2},
}

// create a slice-of-slices of allowable moves for whatever
// side-size of board the user chooses. This global variable
// gets set so that allowableMoves[x][y] is a list of the
// legal (on-the-board) board coordinates of a knight's possible
// next positions.
var allowableMoves [][][][2]int

func main() {
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	allowableMoves = generateMoves(n)

	board := make([][]bool, n)
	for i := 0; i < n; i++ {
		board[i] = make([]bool, n)
	}

	sofar := make([][2]int, n*n)

	// Have to place the first knight on the board, since tryMoves uses
	// the current position to determine which potential moves to look
	// at next.  Find all tours from that starting <x,y> coordinate.

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			sofar[0][0], sofar[0][1] = i, j
			board[i][j] = true
			tryMoves(1, i, j, board, n*n, sofar)
			board[i][j] = false
		}
	}
}

func tryMoves(ply int, a, b int, board [][]bool, boardSize int, sofar [][2]int) {

	// Found a tour, time to stop recursing.
	if ply == boardSize {
		fmt.Printf("%v\n", sofar)
		return
	}

	for _, move := range allowableMoves[a][b] {
		x, y := move[0], move[1]
		if !board[x][y] {
			board[x][y] = true
			sofar[ply][0], sofar[ply][1] = x, y
			tryMoves(ply+1, x, y, board, boardSize, sofar)
			board[x][y] = false
		}
	}
}

// generateMoves creates a slice-of-slices the same side-size
// as the board. Each <x,y> position in the slice-of-slices
// has a slice of <i,j> board positions a knight can move to
// from <x,y>
func generateMoves(sideSize int) [][][][2]int {
	moves := make([][][][2]int, sideSize)
	for i := 0; i < sideSize; i++ {
		moves[i] = make([][][2]int, sideSize)
		for j := 0; j < sideSize; j++ {
			var localMoves [][2]int
			for _, increment := range moveIncrements {
				x, y := i+increment[0], j+increment[1]
				if x >= 0 && y >= 0 && x < sideSize && y < sideSize {
					localMoves = append(localMoves, [2]int{x, y})
				}
			}
			moves[i][j] = localMoves
		}
	}
	return moves
}

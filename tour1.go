package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

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
		for j := 0; j < n; j++ {
			fmt.Printf("<%d,%d>  %v\n", i, j, allowableMoves[i][j])
		}
	}

	sofar := make([][2]int, n*n+1)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			sofar[0][0], sofar[0][1] = i, j
			board[i][j] = true
			tryMoves(1, i, j, board, n, sofar)
			board[i][j] = false
		}
	}
}

func tryMoves(ply int, a, b int, board [][]bool, sideSize int, sofar [][2]int) {
	sz := sideSize * sideSize

	if ply == sz {
		for i := 0; i < sz; i++ {
			fmt.Printf("%d    %d %d\n", i, sofar[i][0], sofar[i][1])
		}
		os.Exit(0)
	}

	for _, move := range allowableMoves[a][b] {
		x, y := move[0], move[1]
		if !board[x][y] {
			board[x][y] = true
			sofar[ply][0], sofar[ply][1] = x, y
			tryMoves(ply+1, x, y, board, sideSize, sofar)
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

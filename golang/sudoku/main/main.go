package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	flag.Parse()
	puzzle := flag.Arg(0)
	if len(puzzle) != 81 {
		fmt.Println("puzzle must be 81 characters long")
	}

	board := BoardFromString(puzzle)
	if Solve(board) {
		fmt.Println(board.String())
	} else {
		fmt.Println("No solution found")
	}
}

func BoardFromString(input string) *Board {
	// Check if the input has exactly 81 characters
	if len(input) != 81 {
		log.Fatal("Input must be exactly 81 characters long")
	}

	var board Board

	// Convert string into a 9x9 board
	for i := 0; i < 81; i++ {
		// Convert character to integer
		num, err := strconv.Atoi(string(input[i]))
		if err != nil {
			num = -2
		}
		board[i/9][i%9] = num
	}

	return &board
}

type Board [9][9]int

func (b *Board) String() string {
	buf := strings.Builder{}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			_, _ = fmt.Fprintf(&buf, "%d", b[i][j])
		}
	}
	return buf.String()
}

func Solve(board *Board) bool {
	return solve(board, 0, 0)
}

func solve(board *Board, startR, startC int) bool {
	row, col, ok := findNextEmpty(board, startR, startC)
	if !ok {
		return true
	}

	for guess := 1; guess <= 9; guess++ {
		if isValid(board, guess, row, col) {
			board[row][col] = guess
			r, c := increment(row, col)
			if solve(board, r, c) {
				return true
			}
		} else {
			continue
		}
		board[row][col] = -1
	}

	return false
}

func increment(row, col int) (int, int) {
	col++
	if col >= 9 {
		row++
		col = 0
	}
	return row, col
}

func findNextEmpty(board *Board, startR, startC int) (int, int, bool) {
	for i := startR; i < 9; i++ {
		var calcStartC int
		if i == startR {
			calcStartC = startC
		}

		for j := calcStartC; j < 9; j++ {
			if board[i][j] <= 0 {
				return i, j, true
			}
		}
	}

	return 0, 0, false
}

func isValid(board *Board, guess, r, c int) bool {
	if guess < 1 || guess > 9 {
		panic("guess must be in range 1-9")
	}

	for _, val := range board[r] {
		if val == guess {
			return false
		}
	}

	for _, row := range board {
		if row[c] == guess {
			return false
		}
	}

	blockR := (r / 3) * 3
	blockC := (c / 3) * 3
	for i := blockR; i < blockR+3; i++ {
		for j := blockC; j < blockC+3; j++ {
			if board[i][j] == guess {
				return false
			}
		}
	}

	return true
}

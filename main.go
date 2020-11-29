package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	PC   = "PC"
	USER = "USER"
	DRAW = "DRAW"
	NONE = "NONE"
)

func main() {
	fmt.Println("*** Welcome to XO game ***")

ResetGame:
	rand.Seed(time.Now().Unix())

	board := make([]string, 9)

	for i := 0; i < len(board); i++ {
		board[i] = " "
	}

	turn := turn()

	for {
		displayBoard(board)

		finished, winner := finished(board)
		if finished {
			switch winner {
			case USER:
				fmt.Println("You are the winner!")
			case PC:
				fmt.Println("PC is the winner!")
			case DRAW:
				fmt.Println("DRAW")
			}
		PLAYAgainQuestion:
			fmt.Println("Do you want to play agian? Y/n")
			var command string
			fmt.Scanf("%s\n", &command)
			switch strings.ToUpper(command) {
			case "Y":
				goto ResetGame
			case "N":
				fmt.Println("Good Bye!")
				os.Exit(0)
			default:
				fmt.Println("Invalid command")
				goto PLAYAgainQuestion
			}
		}

		if turn == PC {
			board[pcMove(board)] = "X"
			turn = USER
		} else {
		UserInput:
			fmt.Println("Now It's your turn, Enter your number:")
			var choose int
			fmt.Scanf("%d\n", &choose)

			choose--
			if choose < 0 || choose > 9 || board[choose] != " "{
				fmt.Println("Invalid input! (You have to enter a number between 0~9)")
				goto UserInput
			}

			board[choose] = "O"
			turn = PC
		}
	}

}

func finished(b []string) (bool, string) {
	finished, winner := false, NONE

	// Check linear wins
	finished, winner = checkLinearFinish(b, 0)
	if finished {
		return finished, winner
	}
	finished, winner = checkLinearFinish(b, 1)
	if finished {
		return finished, winner
	}
	finished, winner = checkLinearFinish(b, 2)
	if finished {
		return finished, winner
	}

	// Check columnar wins
	finished, winner = checkColumnarFinish(b, 0)
	if finished {
		return finished, winner
	}

	finished, winner = checkColumnarFinish(b, 1)
	if finished {
		return finished, winner
	}

	finished, winner = checkColumnarFinish(b, 2)
	if finished {
		return finished, winner
	}

	// Check diagonal win
	finished, winner = checkDiagonalFinish(b)
	if finished {
		return finished, winner
	}

	// find out how many cells are empty
	emptyCells := 0

	for i := 0; i < len(b); i++ {
		if b[i] == " " {
			emptyCells++
		}
	}

	if emptyCells == 0{
		return true, DRAW
	}

	return finished, winner
}

func checkLinearFinish(b []string, lineIdx int) (bool, string){
	firstInLine := lineIdx * 3
	if b[firstInLine] != " " && b[firstInLine] == b[firstInLine+ 1] && b[firstInLine+ 1] == b[firstInLine+ 2] {
		if b[firstInLine] == "X" {
			return true, PC
		}else {
			return true, USER
		}
	}
	return false, NONE
}

func checkColumnarFinish(b []string, colIdx int) (bool, string){
	if b[0 + colIdx] != " " && b[0 + colIdx] == b[3 + colIdx] && b[3 + colIdx] == b[6 + colIdx]{
		if b[0 + colIdx] == "X" {
			return true, PC
		}else {
			return true, USER
		}
	}
	return false, NONE
}

func checkDiagonalFinish(b []string) (bool, string){
	if b[0] != " " && b[0] == b[4] && b[4] == b[8] {
		if b[0] == "X" {
			return true, PC
		} else {
			return true, USER
		}
	}
	if b[2] != " " && b[2] == b[4] && b[4] == b[6] {
		if b[2] == "X" {
			return true, PC
		} else {
			return true, USER
		}

	}
	return false, NONE
}

func rnd(min, max int) int{
	return rand.Intn(max - min + 1) + min
}

func turn() string{
	turn := rnd(0, 1)

	if turn % 2 == 0 {
		return "PC"
	}
	return "USER"
}

func pcMove(board []string) int{
	emptyCells := make([]int, 0)
	// find index of empty cells
	for i, val := range board {
		if val == " " {
			emptyCells = append(emptyCells, i)
		}
	}

	// Chose
	selectedIdx := rnd(0, len(emptyCells) - 1)

	return emptyCells[selectedIdx]
}

func displayBoard(b []string) {
	print("\033[H\033[2J")
	fmt.Printf(`
+-----------------------+
|1      |2      |3      |
|   ` + b[0] + `   |   ` + b[1] + `   |   ` + b[2] + `   |
|       |       |       |
+-------+-------+-------+
|4      |5      |6      |
|   ` + b[3] + `   |   ` + b[4] + `   |   ` + b[5] + `   |
|       |       |       |
+-------+-------+-------+
|7      |8      |9      |
|   ` + b[6] + `   |   ` + b[7] + `   |   ` + b[8] + `   |
|       |       |       |
+-----------------------+
    `)
}

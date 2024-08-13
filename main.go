package main

import (
	"fmt"
	"math/rand/v2"
)

const N = 4
const OTWO = 3 //3

func main() {
	var board [N][N]int
	var wasd string
	var tempboard [N][N]int
	choose(&board)
	choose(&board)
	read(&board)
	for {
		tempboard = board
		fmt.Scan(&wasd)
		switch wasd {
		case "w":
			compreseANDmove(&board)
		case "s":
			reverse(&board)
			compreseANDmove(&board)
			reverse(&board)
		case "a":
			transpose(&board)
			compreseANDmove(&board)
			transpose(&board)
		case "d":
			transpose(&board)
			reverse(&board)
			compreseANDmove(&board)
			reverse(&board)
			transpose(&board)
		default:
			fmt.Println("Wrong command")
			continue
		}
		if tempboard == board { //Проверка на возможность мува
			continue
		}
		choose(&board)
		read(&board)
	}
}

func compreseANDmove(board *[N][N]int) { //Где то тут багуля на двойную сборку 2+2+4 = 8 сразу((
	for i := 0; i < N; i++ { //Еще <2+2+2<  = 2+4 а норм 4+2
		for j := 0; j < N; j++ {
			if board[i][j] != 0 {
				for k := 0; k < i; k++ {
					if board[k][j] == 0 {
						board[k][j], board[i][j] = board[i][j], board[k][j]
					}
				}
				for k := 1; k < N; k++ {
					if board[k-1][j] == board[k][j] {
						board[k-1][j] += board[k][j]
						board[k][j] = 0
					}
				}
			}
			if board[i][j] == 2048 {
				panic("You won")
			}
		}
	}
}

func transpose(board *[N][N]int) {
	tempboard := *board
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			board[i][j] = tempboard[j][i]
		}
	}
}

func reverse(board *[N][N]int) {
	for i := 0; i < N/2; i++ {
		board[i], board[(N-1)-i] = board[(N-1)-i], board[i]
	}
}

func read(board *[N][N]int) {
	for i, j := range board {
		fmt.Println(i, j)
	}
	fmt.Println()
}

func choose(board *[N][N]int) { //спавн цыферки + проверка на луз
	var lose int
	for {
		p1 := rand.IntN(N)
		p2 := rand.IntN(N)
		if board[p1][p2] == 0 {
			if rand.IntN(4) == OTWO {
				board[p1][p2] = 4
			} else {
				board[p1][p2] = 2
			}
			break
		}
		if lose > (N*N)+2 {
			panic("You lose")
		}
		lose++
	}
}

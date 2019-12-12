package main

import (
	"fmt"
	"math/rand"
	"time"
)

type tictacboard [3][3]rune

func main() {

	rand.Seed(time.Now().UnixNano())

	var playerMove bool
	var whoWon string
	var win bool

	var board tictacboard

	fmt.Println("Starting Game:Board Empty")
	board.displayBoard()

	if rand.Intn(2) ==0 {
		playerMove = true
	} else {
		playerMove = false
	}

	for i:=0 ; i<9 ; i++ {
		if playerMove {
			fmt.Println("Player Move: ", i+1)
			board.player()
			playerMove =false
		} else {
			fmt.Println("Computer Move: ", i+1)
			time.Sleep(time.Second)
			board.computer()
			playerMove = true

		}

		if whoWon , win = board.check(); win {
			break
		}
		board.displayBoard()

	}


	fmt.Printf("*****%v won*****\nFinal Board View:\n", whoWon)
	board.displayBoard()
}

func (t *tictacboard) displayBoard() {
	fmt.Print("-------------")
	for i := 0; i < 3; i++ {
		fmt.Print("\n|")
		for j := 0; j < 3; j++ {
			fmt.Printf(" %c |", t[i][j])
		}
		fmt.Print("\n-------------")
	}
	fmt.Print("\n")
}

func (t *tictacboard) player() {

	var x,y int 

	fmt.Println("Enter the Row(1-3) and the Column(1-3) positions: ")
	if _,err := fmt.Scan(&x,&y); err == nil {
		x--
		y--
		if(x>=0 && x<=2) && (y>=0 && y<=2) && (t[x][y] == 0) {
			t[x][y] ='X'
		} else {
			fmt.Println("Invalid input or position not empty. Try Again")
			t.player()
		}
	} else {
		fmt.Println("Invalid input or position not empty. Try Again")
		t.player()
	}
}

func (t *tictacboard) computer() {

	var x,y int 
	for {
		x = rand.Intn(3)
		y = rand.Intn(3)
		if t[x][y] == 0 {
			t[x][y] = 'O'
			break
		}
	}
}

func (t *tictacboard) check() (string,bool) {
	for i:=0;i<3;i++ {
		if(rune(t[i][0]) == 'X') && (t[i][0] == t[i][1]) && (t[i][0] == t[i][2]) {
			return "Player" , true
		} else if (rune(t[i][0]) == 'O') && (t[i][0] == t[i][1]) && (t[i][0] == t[i][2]) {
			return "Computer", true
		}
	}

	for i := 0; i < 3; i++ {
		if (rune(t[0][i]) == 'X') && (t[0][i] == t[1][i]) && (t[0][i] == t[2][i]) {
			return "Player", true
		} else if (rune(t[0][i]) == 'O') && (t[0][i] == t[1][i]) && (t[0][i] == t[2][i]) {
			return "Computer", true
		}
	}

	if ((rune(t[0][0]) == 'X') && (t[0][0] == t[1][1]) && (t[1][1] == t[2][2])) || ((rune(t[0][2]) == 'X') && (t[0][2] == t[1][1]) && (t[1][1] == t[2][0])) {
		return "Player", true
	} else if ((rune(t[0][0]) == 'O') && (t[0][0] == t[1][1]) && (t[1][1] == t[2][2])) || ((rune(t[0][2]) == 'O') && (t[0][2] == t[1][1]) && (t[1][1] == t[2][0])) {
		return "Computer", true
	}

	return "No one" , false
}

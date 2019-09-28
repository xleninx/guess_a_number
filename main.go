package main

import (
	"fmt"
	"github.com/gookit/color"
	"math/rand"
	"time"
)

func main() {
	color.Cyan.Println("Guess the number i'm thinking between 1 and 100")
	rand.Seed(time.Now().UnixNano())
	guess_a_number(rand.Intn(100))
}

func guess_a_number(number int) {
	var userNumber int = 0
	for number != userNumber {
		fmt.Scanf("%d Try", &userNumber)
		if number < userNumber {
			smaller_number()
		} else if number > userNumber {
			bigger_number()
		} else {
			user_win()
		}
	}
}

func smaller_number() {
	color.Yellow.Println("The number is smaller")

}

func bigger_number() {
	color.Green.Println("The number is bigger")
}

func user_win() {
	color.Green.Println("Great you guessed")
}

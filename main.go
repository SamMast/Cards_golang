package main

import "fmt"

func main() {
	// These 2 lines of code are the same:
	// var card string = "Ace of Spades"
	// only use := when assigning a new variable
	card := "Ace of Spades"
	card = "Five of Diamonds"

	fmt.Println(card)
}

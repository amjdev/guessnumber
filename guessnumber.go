package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

func youGuess() {
	rand.Seed(time.Now().Unix())
	number := rand.Intn(99) + 1
	found := false
	var guess int

	for !found {
		fmt.Print("What's your try? ")
		fmt.Scanf("%d", &guess)
		if guess == number {
			fmt.Println("You found it!")
			found = true
		} else if guess > number {
			fmt.Println("My number is lower than", guess)
		} else {
			fmt.Println("My number is greater than", guess)
		}
	}
}

func iGuess() {
	min := 1
	max := 100
	found := false
	var myTry int
	var answer string

	for !found {
		myTry = int(math.Trunc(float64(min+max) / 2.0))
		fmt.Print("Is it ", myTry, "? (G)reater, (L)ower or (R)ight  ")
		fmt.Scanf("%s", &answer)
		answer = strings.ToUpper(answer)
		if answer == "R" {
			fmt.Println("Yes! I did it! Thank you for playing")
			found = true
		} else if answer == "G" {
			min = myTry + 1
		} else if answer == "L" {
			max = myTry - 1
		} else {
			fmt.Println("Sorry, I don't understand you")
		}
		if !found && max < min {
			fmt.Println("You cheated me. I don't want to play anymore :(")
			found = true
		}
	}
}

func main() {
	fmt.Println("I'm thinking in a number between 1 and 100")
	youGuess()
	fmt.Println("Congratulations!!!!\n")
	fmt.Println("Now it's my turn. Think in a number between 1 and 100")
	iGuess()
}

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Guess the number!")

	secretNumber := rand.Intn(100) + 1

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Please input your guess.")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Please type a number!")
			continue
		}

		switch {
		case guess < secretNumber:
			fmt.Println("Too small!")
		case guess > secretNumber:
			fmt.Println("Too big!")
		default:
			fmt.Println("You win!")
			return
		}
	}
}

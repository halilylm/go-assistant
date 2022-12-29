package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// askSecretKey reads the secret key in the user input
func askSecretKey() (string, error) {
	fmt.Print("Enter your api key: ")
	return readUserInput(), nil
}

// askQuestion reads the question in the user input
func askQuestion() (string, error) {
	fmt.Printf("You: ")
	return readUserInput(), nil
}

// readUserInput read user's input
func readUserInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	return scanner.Text()
}

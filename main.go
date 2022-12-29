package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("welcome to conversation, type exit to stop the conversation")

	// get secret key from keyring
	apiKey, err := getApiKey()
	if err != nil {
		openBrowser()
		apiKey, err = askSecretKey()
		if err != nil {
			log.Fatalln(err)
		}
		setApiKey(apiKey)
	}

	// start the conversation
	for {
		question, err := askQuestion()
		if err != nil {
			log.Fatalln(err)
		}
		// stop the conversation if exit typed
		if question == "exit" {
			break
		}
		getAnswer(apiKey, question)
	}

	fmt.Println("see you then!")
}

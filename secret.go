package main

import (
	"log"

	"github.com/zalando/go-keyring"
)

const (
	service = "gpt_secret_key"
	user    = "go"
)

// setApiKey sets api key
func setApiKey(value string) {
	err := keyring.Set(service, user, value)
	if err != nil {
		log.Fatal(err)
	}
}

// getApiKey gets api key
func getApiKey() (string, error) {
	return keyring.Get(service, user)
}

// deleteApiKey deletes api key
func deleteApiKey() error {
	return keyring.Delete(service, user)
}

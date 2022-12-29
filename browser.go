package main

import (
	"errors"
	"log"
	"os/exec"
	"runtime"
)

const (
	url = "https://beta.openai.com/account/api-keys"
)

// openBrowser opens a browser tab to see api keys in openai
func openBrowser() {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = errors.New("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"context"
	"fmt"
	"log"

	"github.com/PullRequestInc/go-gpt3"
)

const (
	maxTokens   = 511
	temperature = 0.5
	engine      = gpt3.TextDavinci003Engine
)

// getAnswer gets answer from gpt
func getAnswer(apiKey, question string) {
	i := 0
	ctx := context.Background()
	client := gpt3.NewClient(apiKey)
	if err := client.CompletionStreamWithEngine(ctx, engine, gpt3.CompletionRequest{
		Prompt: []string{
			question,
		},
		MaxTokens:   gpt3.IntPtr(maxTokens),
		Temperature: gpt3.Float32Ptr(temperature),
	}, func(resp *gpt3.CompletionResponse) {
		if i > 1 {
			fmt.Print(resp.Choices[0].Text)
		}
		i++
	}); err != nil {
		if err := deleteApiKey(); err != nil {
			log.Fatalln(err)
		}
		log.Fatalln(err)
	}
	fmt.Println()
}

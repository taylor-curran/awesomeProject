package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

func main() {
	apiKey := os.Getenv("OPENAI_API_KEY")
	
	if apiKey == "" {
		fmt.Println("Error: OPENAI_API_KEY environment variable not set")
		os.Exit(1)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	client := openai.NewClient(option.WithAPIKey(apiKey))

	prompt := "Write a haiku about programming in Go:"
	
	chatCompletion, err := client.Chat.Completions.New(ctx, openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(prompt),
		},
		Model:       openai.ChatModelGPT3_5Turbo,
		MaxTokens:   openai.Int(150),
		Temperature: openai.Float(0.7),
	})

	if err != nil {
		fmt.Printf("Error with OpenAI API: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("\nPrompt: " + prompt)
	fmt.Println("\nResponse from OpenAI:")
	
	if len(chatCompletion.Choices) > 0 {
		fmt.Println(chatCompletion.Choices[0].Message.Content)
		
		fmt.Printf("\nModel: %s\n", chatCompletion.Model)
		fmt.Printf("Tokens used: %d\n", chatCompletion.Usage.TotalTokens)
	} else {
		fmt.Println("No completion choices returned from the AI")
	}
}

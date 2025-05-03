# OpenAI Completion Request Demo

A simple Go application that demonstrates making a completion request to OpenAI's API.

## Prerequisites

- Go installed on your system
- An OpenAI API key

## Setup

1. Clone this repository
2. Set your OpenAI API key as an environment variable:

```bash
export OPENAI_API_KEY="your-api-key-here"
```

## Running the Demo

To run the demo, execute:

```bash
go run main.go
```

The program will send a request to OpenAI's API asking for a haiku about Go programming and display the response.

## Code Structure

- `main.go` - Contains the entire demo application
- `agent_plan.md` - Contains the planning document for this demo

## Features

- Makes a request to OpenAI's completion API
- Parses and displays the response
- Handles errors gracefully
- Uses environment variables for API keys (security best practice)

## Sample Output

```
Sending request to OpenAI API...

Prompt: Write a haiku about programming in Go:

Response from OpenAI:
[The haiku from the API will be displayed here]

Model: gpt-3.5-turbo-instruct
Tokens used: 42
```

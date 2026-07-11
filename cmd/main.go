package main

import (
	"context"
	"deepn/internal"
	"deepn/internal/providers"
	"fmt"
)

func main() {

	container, err := internal.Bootstrap()
	if err != nil {
		fmt.Printf("Error occurred: %v\n", err)
		return
	}

	provider, err := container.GetProvider()
	if err != nil {
		fmt.Printf("Error occurred: %v\n", err)
		return
	}

	response, err := provider.Chat(context.Background(), providers.Request{Message: "Hello, is your name steven ?"})
	if err != nil {
		fmt.Printf("Error occurred: %v\n", err)
		return
	}
	fmt.Printf("Response: %s\n", response.Message)

}

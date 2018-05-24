package main

import (
	"context"
	"fmt"
	"log"

	language "cloud.google.com/go/language/apiv1"
	"github.com/alice/googleApis"
)

func main() {
	ctx := context.Background()

	client, err := language.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// text := "Olá, meu nome é Igor! Tudo bem contigo?"
	text := "Eu odeio java, é realmente muito ruim"

	googleApis.AnalizeSentiment(ctx, client, text)

	fmt.Printf("Text: %v\n", text)
}

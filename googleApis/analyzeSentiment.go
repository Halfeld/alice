package googleApis

import (
	"context"
	"fmt"

	language "cloud.google.com/go/language/apiv1"
	"github.com/alice/helpers"
	languagepb "google.golang.org/genproto/googleapis/cloud/language/v1"
)

// AnalizeSentiment returns analyzes from the text
func AnalizeSentiment(ctx context.Context, client *language.Client, text string) *languagepb.AnalyzeSentimentResponse {
	sentiment, err := client.AnalyzeSentiment(ctx, &languagepb.AnalyzeSentimentRequest{
		Document: &languagepb.Document{
			Source: &languagepb.Document_Content{
				Content: text,
			},
			Type: languagepb.Document_PLAIN_TEXT,
		},
		EncodingType: languagepb.EncodingType_UTF8,
	})
	helpers.ErrorHandler(err, "Cannot analyze this text")

	if sentiment.DocumentSentiment.Score >= 0 {
		fmt.Println("Sentiment: positive")
	} else {
		fmt.Println("Sentiment: negative")
	}

	return sentiment
}

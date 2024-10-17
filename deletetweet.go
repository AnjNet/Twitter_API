package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/dghubble/oauth1"
)

// Function to delete a tweet
func DeleteTweet(apiKey, apiSecret, accessToken, accessSecret, tweetID string) {
	// OAuth1 Config
	config := oauth1.NewConfig(apiKey, apiSecret)
	token := oauth1.NewToken(accessToken, accessSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter API Endpoint for deleting a tweet
	url := "https://api.twitter.com/1.1/statuses/destroy/" + tweetID + ".json"

	// Create POST request to delete the tweet
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	// Send the request
	resp, err := httpClient.Do(req)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response: %v", err)
	}

	// Output the response from Twitter
	fmt.Println("Response:", string(body))
}

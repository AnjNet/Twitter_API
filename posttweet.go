package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/dghubble/oauth1"
)

// Function to post a tweet
func PostTweet(apiKey, apiSecret, accessToken, accessSecret, tweetText string) {
	// OAuth1 Config
	config := oauth1.NewConfig(apiKey, apiSecret)
	token := oauth1.NewToken(accessToken, accessSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter API Endpoint for posting a tweet
	url := "https://api.twitter.com/1.1/statuses/update.json"

	// Create POST request
	req, err := http.NewRequest("POST", url, strings.NewReader("status="+tweetText))
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Send request
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

	fmt.Println("Response:", string(body))
}

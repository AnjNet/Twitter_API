package main

import (
	"flag"
	"log"
)

// Twitter API Credentials
const (
	apiKey       = "****L0GK4K"
	apiSecret    = "YOUR_API_SECRET" // Replace with actual API secret
	accessToken  = "1846198139268517888-5RGhJ3ntBhG8ipLCSk0p4dDiYh7jFl"
	accessSecret = "OG84J1ytynDSqCAKWpm4DvUzYdVoXi3BoCpWp1iO9CqFJ"
)

func main() {
	// Command-line flags to determine action
	action := flag.String("action", "post", "Action to perform: post or delete")
	tweetText := flag.String("tweet", "Hello from Twitter API!", "Text of the tweet to post")
	tweetID := flag.String("id", "", "ID of the tweet to delete")
	flag.Parse()

	// Determine which action to perform based on the user input
	if *action == "post" {
		PostTweet(apiKey, apiSecret, accessToken, accessSecret, *tweetText)
	} else if *action == "delete" {
		if *tweetID == "" {
			log.Fatalf("Tweet ID must be provided for delete action")
		}
		DeleteTweet(apiKey, apiSecret, accessToken, accessSecret, *tweetID)
	} else {
		log.Fatalf("Unknown action: %s. Use 'post' or 'delete'", *action)
	}
}

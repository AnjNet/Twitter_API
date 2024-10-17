# **Twitter API Project**
## **Introduction**
The project shows how to use the Twitter API with the Go programming language. It allows us to post a tweet and delete a tweet. We use OAuth1.0a for authentication to make sure we can connect to Twitter safely. This assignment will help to learn how to work with APIs, handle authentication, and send requests and get responses in Go.
## **Pre-Requisites**
Setup a Twitter Developer Account from **https://developer.x.com/**
Then go to the Project and get the below details:
-API Key
-API Secret Key
-Bearer Token
-Access Token
-Access Token Secret
-Set up User Authentication Settings and configure OAuth callback URL as http://localhost:3000
## **Define System Variables**
Replace the following details in the main.go file with the actual API Key, API Secret, Access Token, and Access Secret which we got from Twitter Developer Portal.
## **Install the packages**
`go get github.com/dghubble/oauth1`
`go get github.com/dghubble/go-twitter/twitter`
## **Posting Tweet**
To run posttweet, run the following command `go run main.go -action post -tweet "Hello Twitter API"`
## **Deleting Tweet**
To run deletetweet, run the following command `go run main.go -action delete -tweetID "TweetID_to_delete"`
## **Code**
1. `main.go`:
The entry point of the program.
2. `posttweet.go`:
Contains the function `PostTweet` that handles posting a new tweet to Twitter.
3. `deletetweet.go`:
Contains the function `DeleteTweet` that handles deleting an existing tweet from Twitter.
4. `go.mod`:
Go module file that defines the dependencies required for the project.


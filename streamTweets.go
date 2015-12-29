package main

import (
	"os"
	"fmt"
	"net/url"
	"log"
	
	"github.com/ChimeraCoder/anaconda"
)

func init() {
	consumer_key := os.Getenv("TWITTER_CONSUMER_KEY")
	consumer_secret := os.Getenv("TWITTER_CONSUMER_SECRET")
	access_token := os.Getenv("TWITTER_ACCESS_TOKEN")
	access_token_secret := os.Getenv("TWITTER_ACCESS_TOKEN_SECRET")
	
	if consumer_key == "" || consumer_secret == "" {
		log.Fatalln("consumer tokens left blank")
	}
	
	if access_token == "" || access_token_secret == "" {
		log.Fatalln("access tokens left blank")
	}
	fmt.Println("Consumer tokens and keys available in environment \n Initiating connection with twitter api")
}

func main() {
	consumer_key := os.Getenv("TWITTER_CONSUMER_KEY")
	consumer_secret := os.Getenv("TWITTER_CONSUMER_SECRET")
	access_token := os.Getenv("TWITTER_ACCESS_TOKEN")
	access_token_secret := os.Getenv("TWITTER_ACCESS_TOKEN_SECRET")
	
	anaconda.SetConsumerKey(consumer_key)
	anaconda.SetConsumerSecret(consumer_secret)
	api := anaconda.NewTwitterApi(access_token, access_token_secret)
	v := url.Values{}
	//Get stream
	
	twitterStream := api.PublicStreamSample(v)
	fmt.Println("Starting streaming tweets ...")
	for t := range twitterStream.C {
		switch v := t.(type) {
			case anaconda.Tweet:
				if v.Lang == "en"{ //Capture only Tweets in English
					fmt.Printf("%-15s: %s\n", v.User.ScreenName, v.Text)
				}
		}
	}
}

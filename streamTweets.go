package main

import (
	"os"
	"fmt"
	"net/url"
	"log"
	"time"
	"bufio"
	
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
//streamListener will listen to tweets
func streamListener(stream chan anaconda.Tweet) {
	consumer_key := os.Getenv("TWITTER_CONSUMER_KEY")
	consumer_secret := os.Getenv("TWITTER_CONSUMER_SECRET")
	access_token := os.Getenv("TWITTER_ACCESS_TOKEN")
	access_token_secret := os.Getenv("TWITTER_ACCESS_TOKEN_SECRET")
	
	anaconda.SetConsumerKey(consumer_key)
	anaconda.SetConsumerSecret(consumer_secret)
	api := anaconda.NewTwitterApi(access_token, access_token_secret)
	v := url.Values{}
	
	twitterStream := api.PublicStreamSample(v)
	fmt.Println("Starting streaming tweets ...")
	for t := range twitterStream.C {
		switch v := t.(type) {
			case anaconda.Tweet:
				if v.Lang == "en"{ //Capture only Tweets in English
					stream <- v
				}
		}
	}
}
//streamWriter will listen to channel and process/extract
//necessary data from tweets
//func streamWriter(stream chan anaconda.Tweet) {
	
//}
func main() {
	
	//Get stream
	filename := "tweetStream-" + time.Now().Format(time.RFC3339) + ".csv"
	fmt.Printf("Writting tweets to %s\n",filename)
	fout, err := os.Create(filename)
	if err != nil {
		log.Panic(err)
	}
	defer func() {
		if err := fout.Close(); err != nil {
			log.Panic(err)
		}
		log.Printf("Closed file out")
	}()
	writer := bufio.NewWriter(fout)
	defer func() {
		if err := writer.Flush(); err != nil {
			log.Panic(err)
		}
		log.Printf("flushed contents of writer")
	}()
	fmt.Fprintln(writer, "Text, User.TimeZone, CreatedAt, longitude, latitude")
	stream := make(chan anaconda.Tweet, 10) //buffered channel size 10
	go streamListener(stream)
	//go streamWriter(stream)
	for v := range(stream){
		if v.HasCoordinates() { //Assumption that coordinates are type point
			longitude, err1 := v.Longitude()
			latitude, err2 := v.Latitude()
			if err1 != nil || err2 != nil { // ideally this should never happen
				fmt.Fprintf(writer, "%s, %s, %s, , \n",v.Text, v.User.TimeZone, v.CreatedAt)
			} else {
				fmt.Fprintf(writer, "%s, %s, %s, %v, %v \n",v.Text, v.User.TimeZone, v.CreatedAt, longitude, latitude)
			}
		} else {
			fmt.Fprintf(writer, "%s, %s, %s, , \n",v.Text, v.User.TimeZone, v.CreatedAt)
		}
		
	}
}

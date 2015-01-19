package main

import (
	"flag"
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"net/url"
	"os"
)

func main() {

	api := twitterApi()
	query, sm, pm := flags()

	streamingFilter := url.Values{}
	streamingFilter.Set("track", *query)

	searchResult, err := api.PublicStreamFilter(streamingFilter)

	if err == nil {
		tm := TweetManager{SendMessage: *sm, PersistToPostgres: *pm, InC: make(chan interface{})}
		tm.Init()
		for tweet := range searchResult.C {
			theTweet := tweet.(anaconda.Tweet)
			tm.Manage(theTweet)
		}
	} else {
		fmt.Printf("error: %s", err)
	}

}

func twitterApi() *anaconda.TwitterApi {
	anaconda.SetConsumerKey(os.Getenv("TWITTER_API_CONSUMER_KEY"))
	anaconda.SetConsumerSecret(os.Getenv("TWITTER_API_CONSUMER_SECRET"))
	return anaconda.NewTwitterApi(os.Getenv("TWITTER_API_TOKEN"), os.Getenv("TWITTER_API_SECRET"))
}

func flags() (*string, *bool, *bool) {
	query := flag.String("query", "golang", "defaults to golang")
	sm := flag.Bool("sm", true, "sends a message to the message queue, defaults to true")
	pm := flag.Bool("cm", false, "persist messages to postgres, defaults to false")
	flag.Parse()

	return query, sm, pm
}

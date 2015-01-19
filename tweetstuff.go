package main

import (
	"flag"
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"net/url"
	"os"
	"os/signal"
)

func main() {
	notificationChannel := make(chan os.Signal)
	cleanupChannel := make(chan bool)
	signal.Notify(notificationChannel, os.Interrupt)

	api := twitterApi()
	query, sm, pm := flags()

	streamingFilter := url.Values{}
	streamingFilter.Set("track", *query)

	searchResult, err := api.PublicStreamFilter(streamingFilter)

	if err == nil {
		tm := TweetManager{
			SendMessage:       *sm,
			PersistToPostgres: *pm,
			InC:               notificationChannel,
			TwitterApi:        api,
			CleanupChannel:    cleanupChannel}
		tm.Init()
		go func() {
			for tweet := range searchResult.C {
				theTweet := tweet.(anaconda.Tweet)
				tm.Manage(theTweet)
			}
			close(tm.InC)
		}()
	} else {
		fmt.Printf("error: %s", err)
	}

	<-cleanupChannel

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

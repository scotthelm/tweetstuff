package main

import "github.com/ChimeraCoder/anaconda"
import "fmt"
import "net/url"
import "flag"
import "os"

func main() {
	anaconda.SetConsumerKey(os.Getenv("TWITTER_API_CONSUMER_KEY"))
	anaconda.SetConsumerSecret(os.Getenv("TWITTER_API_CONSUMER_SECRET"))
	api := anaconda.NewTwitterApi(os.Getenv("TWITTER_API_TOKEN"), os.Getenv("TWITTER_API_SECRET"))

	query := flag.String("query", "golang", "defaults to golang")
	flag.Parse()

	streamingFilter := url.Values{}
	streamingFilter.Set("track", *query)
	searchResult, err := api.PublicStreamFilter(streamingFilter)

	if err == nil {
		for tweet := range searchResult.C {
			theTweet := tweet.(anaconda.Tweet)
			fmt.Printf("%s --(%s) \n", theTweet.Text, theTweet.User.Name)
		}
	} else {
		fmt.Printf("error: %s", err)
	}

}

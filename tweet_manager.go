package main

import (
	"encoding/json"
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"gopkg.in/stomp.v1"
)

type TweetManager struct {
	SendMessage       bool
	PersistToPostgres bool
	InC               chan interface{}
	MessageQueueConn  *stomp.Conn
	MQConnectionError error
	MessageQueueUrl   string
	Db                *sqlx.DB
	DbError           error
	TweetSub          *stomp.Subscription
	TweetSubError     error
}

func (tm *TweetManager) Manage(tweet anaconda.Tweet) {
	if tm.SendMessage {
		go tm.SendToMessageQueue(tweet)
	}
	//fmt.Printf("%s --(%s) \n", tweet.Text, tweet.User.Name)
}

func (tm *TweetManager) Init() {
	tm.MessageQueueConn, tm.MQConnectionError = stomp.Dial("tcp", "127.0.0.1:61613", stomp.Options{})

	if tm.PersistToPostgres {
		tm.Db, tm.DbError = sqlx.Connect(
			"postgres",
			"postgres://tweetstuff_login:tweetstuff_login_pass@localhost/tweetstuff")
		tm.TweetSub, tm.TweetSubError = tm.MessageQueueConn.Subscribe("test.queue", stomp.AckAuto)

		go tm.GetMessageTweets()
	}

	go tm.channelComm()

}

func (tm *TweetManager) SendToMessageQueue(t anaconda.Tweet) {
	js, _ := json.Marshal(t)
	tm.MessageQueueConn.Send("test.queue", "application/json", js, nil)
}

func (tm *TweetManager) channelComm() {
	for message := range tm.InC {
		fmt.Printf("manager message received: %s\n", message)
	}
}

func (tm *TweetManager) GetMessageTweets() {
	for message := range tm.TweetSub.C {
		var tweet = new(anaconda.Tweet)
		json.Unmarshal(message.Body, tweet)
		var latitude, longitude float64
		if tweet.Coordinates != nil {
			coords := tweet.Coordinates.(map[string]interface{})["coordinates"].([]interface{})
			latitude = coords[1].(float64)
			longitude = coords[0].(float64)
		}

		myTweet := Tweet{Tweet: tweet.Text, Author: tweet.User.Name, Latitude: latitude, Longitude: longitude}
		myTweet.Insert(tm.Db)
		fmt.Printf("dequeued message: %s\n", tweet.Text)
	}
}

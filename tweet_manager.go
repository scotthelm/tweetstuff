package main

import (
	"encoding/json"
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"gopkg.in/stomp.v1"
)

type TweetManager struct {
	SendMessage       bool
	PersistToPostgres bool
	InC               chan interface{}
	MessageQueueConn  *stomp.Conn
	ConnectionError   error
}

func (tm *TweetManager) Manage(tweet anaconda.Tweet) {
	if tm.SendMessage {
		go tm.SendToMessageQueue(tweet)
	}
	fmt.Printf("%s --(%s) \n", tweet.Text, tweet.User.Name)
}

func (tm *TweetManager) Init() {
	if tm.SendMessage {
		tm.MessageQueueConn, tm.ConnectionError = stomp.Dial("tcp", "127.0.0.1:61613", stomp.Options{})
	}
}

func (tm *TweetManager) SendToMessageQueue(t anaconda.Tweet) {
	js, _ := json.Marshal(t)
	tm.MessageQueueConn.Send("test.queue", "application/json", js, nil)
}

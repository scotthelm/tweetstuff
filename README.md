# Tweetstuff

This is a project to help me learn how to write [Go](http://golang.org) programs.

It reads from the Twitter streaming API using the fantastic
[Anaconda](https://github.com/ChimeraCoder/anaconda) library, and print the
tweet text and author to stdout.

Depending on the command line flags given to it, it will forward the JSON of the
message to a [Stomp protocol](https://stomp.github.io/) message queue broker.
In this project, I use [ActiveMQ](http://activemq.apache.org/). and optionally
dequeue the messages and put selected fields into a database.

## Prerequisites

You must set up ActiveMQ, have a Postgresql database and an active Twitter
account (with an application that you have set up).

To set up active mq on my development machine, I followed the instructions
[found here](http://activemq.apache.org/getting-started.html)

To set up Postgres, I use [Puppet](http://puppetlabs.com/) to provision my
[Vagrant VM](https://www.vagrantup.com/).

Postgresql installation instructions can be found [here](https://www.vagrantup.com/)

Instructions on setting up a Twitter application can be found [here](https://dev.twitter.com/)

## Lessons Learned

* How to [set up a Go development environment](http://golang.org/).
* How to `go get` libraries to build on
* How to read command line flags
* How to read environment variables to protect sensitive information and keep
them out of your source code.
* How to use the Go concurrency primitives to achieve parallel operations
** goroutines to asynchronously perform actions and channels to communicate between them
* How to handle OS signals
* How to connect to services and wire them together.

To be fair, this is a Frankenstein's monster of a project, but It allowed me to
figure out how the pieces fit together.

## Usage

````
./tweetstuff -help
Usage of ./tweetstuff:
  -cm=false: persist messages to postgres, defaults to false
  -query="golang": defaults to golang
  -sm=true: sends a message to the message queue, defaults to true

````


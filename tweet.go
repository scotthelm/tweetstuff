package main

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"time"
)

type Tweet struct {
	Id         int32
	Tweet      string
	Author     string
	Latitude   float64
	Longitude  float64
	created_at time.Time
}

func (t *Tweet) Insert(db *sqlx.DB) {
	db.MustExec(`
	INSERT INTO
	tweets
	(tweet, author, latitude, longitude)
	values
	($1, $2, $3, $4) `,
		t.Tweet,
		t.Author,
		t.Latitude,
		t.Longitude)
}

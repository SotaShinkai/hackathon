package model

import "github.com/oklog/ulid"

type Tweet struct {
	Id       ulid.ULID `json:"id"`
	UserName string    `json:"username"`
	UserId   string    `json:"userId"`
	Content  string    `json:"content"`
}

type TweetNoId struct {
	UserName string `json:"username"`
	UserId   string `json:"userId"`
	Content  string `json:"content"`
}

type Id struct {
	Id string `json:"id"`
}

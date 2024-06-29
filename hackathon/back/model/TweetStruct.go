package model

type Tweet struct {
	Id       string `json:"id"`
	UserName string `json:"username"`
	UserId   string `json:"userId"`
	Content  string `json:"content"`
	Fav      int64  `json:"fav"`
	ReplyId  string `json:"replyId"`
}

type TweetNoId struct {
	UserName string `json:"username"`
	UserId   string `json:"userId"`
	Content  string `json:"content"`
	ReplyId  string `json:"replyId"`
}

type Id struct {
	Id string `json:"id"`
}

type IsFavedTweet struct {
	Id      string `json:"id"`
	IsFaved bool   `json:"isFaved"`
}

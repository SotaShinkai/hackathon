package service

import (
	"back/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/oklog/ulid"
	"log"
)

type TweetService struct{}

func (TweetService) PostTweet(id ulid.ULID, tweet model.TweetNoId) error {
	_, err := db.Exec("INSERT INTO tweet VALUES (?, ?, ?, ?)", id.String(), tweet.UserName, tweet.UserId, tweet.Content)
	if err != nil {
		log.Println("Error inserting tweet", err)
		return err
	}
	return nil
}

func (TweetService) GetTweet() []model.Tweet {
	rows, err := db.Query("SELECT * FROM tweet")
	if err != nil {
		log.Println("Error getting tweet", err)
	}

	tweets := make([]model.Tweet, 0)
	for rows.Next() {
		var tweet model.Tweet
		if err := rows.Scan(&tweet.Id, &tweet.UserName, &tweet.UserId, &tweet.Content); err != nil {
			log.Println("Error getting tweet", err)
			if err := rows.Close(); err != nil {
				log.Println("Error closing tweet", err)
			}
			return nil
		}
		tweets = append(tweets, tweet)
	}
	return tweets
}

func (TweetService) DeleteTweet(id model.Id) error {
	_, err := db.Exec("DELETE FROM tweet WHERE id = ?", id)
	if err != nil {
		log.Println("Error deleting tweet", err)
		return err
	}
	return nil
}

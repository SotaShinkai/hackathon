package controller

import (
	"back/model"
	"back/service"
	"crypto/rand"
	"encoding/json"
	"github.com/oklog/ulid"
	"net/http"
)

func TweetAdd(w http.ResponseWriter, r *http.Request) {

	tweetNoId := model.TweetNoId{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&tweetNoId); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id := ulid.MustNew(ulid.Now(), rand.Reader)

	tweetService := service.TweetService{}
	err := tweetService.PostTweet(id, tweetNoId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func TweetGet(w http.ResponseWriter, r *http.Request) {
	tweetService := service.TweetService{}
	tweets := tweetService.GetTweet()
	bytes, err := json.Marshal(tweets)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}

func TweetDelete(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	id := model.Id{}
	if err := decoder.Decode(&id); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	tweetService := service.TweetService{}
	err := tweetService.DeleteTweet(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	switch r.Method {
	case http.MethodGet:
		TweetGet(w, r)
	case http.MethodPost:
		TweetAdd(w, r)
	case http.MethodDelete:
		TweetDelete(w, r)
	default:
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.WriteHeader(http.StatusOK)
		return
	}
}

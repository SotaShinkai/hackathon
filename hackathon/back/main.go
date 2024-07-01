package main

import (
	"back/controller"
	"back/service"
	"log"
	"net/http"
)

func main() {
	service.InitDatabase()
	http.HandleFunc("/tweets", controller.Handler)

	log.Println("Listen...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

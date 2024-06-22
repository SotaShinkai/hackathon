package main

import (
	"back/controller"
	"back/service"
	"net/http"
)

func main() {
	service.InitDatabase()
	http.HandleFunc("/tweet", controller.handler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

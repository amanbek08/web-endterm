package main

import (
	"log"
	"web-endterm/internal/app"
	"web-endterm/internal/model"
	"web-endterm/internal/service"
	"web-endterm/pkg/store"
)

func main() {
	var posts []*model.Post

	store := store.NewStore(posts)

	service := service.NewService(store)

	server := app.NewServer(service)

	err := server.Start("8080")
	if err != nil {
		log.Fatal(err.Error())
	}
}

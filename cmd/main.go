package main

import (
	"apiProject/configs"
	"apiProject/internal/auth"
	"apiProject/internal/link"
	"apiProject/pkg/db"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	db := db.NewDb(conf)
	router := http.NewServeMux()
	linkRepository := link.NewLinkRepository(db)
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: conf,
	})
	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepository: linkRepository,
	})
	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}
	server.ListenAndServe()
}

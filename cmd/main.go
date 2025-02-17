package main

import (
	"apiProject/configs"
	"apiProject/internal/auth"
	"apiProject/internal/link"
	"apiProject/internal/user"
	"apiProject/pkg/db"
	"apiProject/pkg/middleware"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	db := db.NewDb(conf)
	router := http.NewServeMux()
	linkRepository := link.NewLinkRepository(db)
	userRepository := user.NewUserRepository(db)
	authService := auth.NewUserService(userRepository)

	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config:      conf,
		AuthService: authService,
	})
	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepository: linkRepository,
	})
	stack := middleware.Chain(
		middleware.CORS,
		middleware.Logging,
		middleware.IsAuth)
	server := http.Server{
		Addr:    ":8081",
		Handler: stack(router),
	}
	server.ListenAndServe()
}

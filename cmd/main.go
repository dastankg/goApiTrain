package main

import (
	"apiProject/configs"
	"apiProject/internal/auth"
	"apiProject/internal/link"
	"apiProject/internal/stat"
	"apiProject/internal/user"
	"apiProject/pkg/db"
	"apiProject/pkg/event"
	"apiProject/pkg/middleware"
	"net/http"
)

func App() http.Handler {
	conf := configs.LoadConfig()
	db := db.NewDb(conf)
	router := http.NewServeMux()
	eventBus := event.NewEventBus()

	linkRepository := link.NewLinkRepository(db)
	userRepository := user.NewUserRepository(db)
	statRepository := stat.NewStatRepository(db)

	authService := auth.NewUserService(userRepository)
	deps := stat.StatServiceDeps{
		EventBus:       eventBus,
		StatRepository: statRepository,
	}
	statService := stat.NewStatService(&deps)

	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config:      conf,
		AuthService: authService,
	})
	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepository: linkRepository,
		Config:         conf,
		EventBus:       eventBus,
	})

	stat.NewStatHandler(router, stat.StatHandlerDeps{
		StatRepository: statRepository,
		Config:         conf,
	})
	go statService.AddClick()
	stack := middleware.Chain(
		middleware.CORS,
		middleware.Logging)
	return stack(router)

}

func main() {
	app := App()
	server := http.Server{
		Addr:    ":8081",
		Handler: app,
	}

	server.ListenAndServe()
}

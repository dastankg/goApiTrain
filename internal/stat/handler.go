package stat

import (
	"apiProject/configs"
	"apiProject/pkg/response"
	"net/http"
	"time"
)

type StatHandler struct {
	StatRepository *StatRepository
}

type StatHandlerDeps struct {
	StatRepository *StatRepository
	Config         *configs.Config
}

func NewStatHandler(router *http.ServeMux, deps StatHandlerDeps) {
	handler := &StatHandler{
		StatRepository: deps.StatRepository,
	}
	router.HandleFunc("GET /stat", handler.GetStat())
}

func (handler *StatHandler) GetStat() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		from, _ := time.Parse(time.DateOnly, req.URL.Query().Get("from"))
		to, _ := time.Parse(time.DateOnly, req.URL.Query().Get("to"))
		by := req.URL.Query().Get("by")

		stats := handler.StatRepository.GetStats(by, from, to)
		response.NewResponse(w, stats, 200)
	}
}

package link

import (
	"apiProject/pkg/req"
	"apiProject/pkg/response"
	"fmt"
	"net/http"
)

type LinkHandler struct {
	LinkRepository *LinkRepository
}
type LinkHandlerDeps struct {
	LinkRepository *LinkRepository
}

func NewLinkHandler(router *http.ServeMux, deps LinkHandlerDeps) {
	handler := &LinkHandler{
		LinkRepository: deps.LinkRepository,
	}
	router.HandleFunc("POST /link", handler.Create())
	router.HandleFunc("PATCH /link/{id}", handler.Update())
	router.HandleFunc("DELETE /link/{id}", handler.Delete())
	router.HandleFunc("GET /{hash}", handler.Get())
}

func (handler *LinkHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		body, err := req.HandleBody[LinkCreateRequest](&w, r)
		if err != nil {
			return
		}
		link := NewLink(body.Url)
		createdLink, err := handler.LinkRepository.Create(link)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		response.NewResponse(w, createdLink, 201)

	}
}

func (handler *LinkHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

	}
}
func (handler *LinkHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		id := req.PathValue("id")
		fmt.Println(id)
	}
}
func (handler *LinkHandler) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		hash := req.PathValue("hash")
		link, err := handler.LinkRepository.Get(hash)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		fmt.Println(link.Url)
		http.Redirect(w, req, link.Url, http.StatusTemporaryRedirect)
	}
}

package auth

import (
	"apiProject/configs"
	request_ "apiProject/pkg/req"
	"apiProject/pkg/response"
	"fmt"
	"net/http"
)

type AuthHandler struct {
	*configs.Config
}
type AuthHandlerDeps struct {
	*configs.Config
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config: deps.Config,
	}
	router.HandleFunc("POST /auth", handler.Auth())
	router.HandleFunc("POST /register", handler.Register())
}

func (handler *AuthHandler) Auth() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		body, err := request_.HandleBody[LoginRequest](&w, req)
		if err != nil {
			response.NewResponse(w, err.Error(), 402)
			return
		}
		fmt.Println(body)
		res := LoginResponse{
			Token: "123",
		}
		response.NewResponse(w, res, 201)

	}
}

func (register *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		body, err := request_.HandleBody[RegisterRequest](&w, req)
		if err != nil {
			response.NewResponse(w, err.Error(), 402)
			return
		}
		fmt.Println(body)
		res := RegisterResponse{
			Token: "1234",
		}
		response.NewResponse(w, res, 201)

	}
}

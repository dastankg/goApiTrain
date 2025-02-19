package auth

import (
	"apiProject/configs"
	"apiProject/pkg/jwt"
	request_ "apiProject/pkg/req"
	"apiProject/pkg/response"
	"net/http"
)

type AuthHandler struct {
	*configs.Config
	*AuthService
}
type AuthHandlerDeps struct {
	*configs.Config
	*AuthService
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config:      deps.Config,
		AuthService: deps.AuthService,
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
		email, err := handler.AuthService.Login(body.Email, body.Password)
		if err != nil {
			response.NewResponse(w, err.Error(), 402)
			return
		}
		token, err := jwt.NewJWT(handler.Config.Auth.Secret).Create(jwt.JWTData{
			Email: email,
		})
		if err != nil {
			response.NewResponse(w, err.Error(), 402)
			return
		}
		res := LoginResponse{
			Token: token,
		}
		response.NewResponse(w, res, 201)

	}
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		body, err := request_.HandleBody[RegisterRequest](&w, req)
		if err != nil {
			response.NewResponse(w, err.Error(), 402)
			return
		}

		email, err := handler.AuthService.Register(body.Name, body.Email, body.Password)
		if err != nil {
			response.NewResponse(w, err.Error(), 402)
			return
		}
		token, err := jwt.NewJWT(handler.Config.Auth.Secret).Create(jwt.JWTData{
			Email: email,
		})
		if err != nil {
			response.NewResponse(w, err.Error(), 402)
			return
		}
		res := RegisterResponse{
			Token: token,
		}
		response.NewResponse(w, res, 201)

	}
}

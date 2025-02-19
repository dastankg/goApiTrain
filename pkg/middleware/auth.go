package middleware

import (
	"apiProject/configs"
	"apiProject/pkg/jwt"
	"context"
	"fmt"
	"net/http"
	"strings"
)

const (
	ContextEMailKey = "email"
)

func IsAuth(next http.Handler, config *configs.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bearerToken := r.Header.Get("Authorization")
		fmt.Println(bearerToken)
		if !strings.HasPrefix(bearerToken, "Bearer ") {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(http.StatusText(http.StatusUnauthorized)))
			return
		}
		bearerToken = strings.TrimPrefix(bearerToken, "Bearer ")
		isValid, data := jwt.NewJWT(config.Auth.Secret).Parse(bearerToken)
		fmt.Println(isValid)
		if !isValid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), ContextEMailKey, data.Email)
		req := r.WithContext(ctx)
		next.ServeHTTP(w, req)
	})
}

package middleware

import (
	"fmt"
	"net/http"
	"strings"
)

func IsAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bearerToken := r.Header.Get("Authorization")
		bearerToken = strings.TrimPrefix(bearerToken, "Bearer")
		fmt.Println("Bearer", bearerToken)
		next.ServeHTTP(w, r)
	})
}

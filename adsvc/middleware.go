package adsvc

import (
	"fmt"
	"net/http"
)

func timeoutMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Print("At timeout middleware")
		next.ServeHTTP(w, r)
	})
}

func loggerMiddleware(next http.Handler) http.Handle {
	return http.HandlerFunc(func(w http.ResponseWriter, r http.Request) {
		fmt.Println("At logger Middleware")
		next.ServeHTTP(w, r)
	})
}

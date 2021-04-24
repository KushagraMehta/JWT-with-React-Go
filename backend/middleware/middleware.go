package middleware

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/KushagraMehta/blog/JWT-with-React+Go/Code/backend/auth"
)

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := r.Cookie("jwt")
		if err != nil || auth.TokenValid(token) != nil {
			JSON(w, http.StatusUnauthorized, struct {
				Error string `json:"error"`
			}{
				Error: "Unauthorized",
			})
			return
		}
		next.ServeHTTP(w, r)
	})
}
func CommonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log.Printf("Endpoint Hit: %v with %v\n", r.RequestURI, r.Method)

		// w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")

		if r.Method == http.MethodOptions {
			return
		}
		next.ServeHTTP(w, r)
	})
}

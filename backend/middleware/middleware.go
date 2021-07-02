package middleware

import (
	"encoding/json"
	"fmt"
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

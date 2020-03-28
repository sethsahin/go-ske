package middlewares

import (
	"net/http"
)

func setMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}

package middleware

import (
	"net/http"
)

func WithCorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Vary", "Origin")
		// Add the "Vary: Access-Control-Request-Method" header.
		w.Header().Add("Vary", "Access-Control-Request-Method")

		w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "PUT,DELETE,PATCH")
		// Write the headers along with a 200 OK status and return from
		// the middleware with no further action.
		w.WriteHeader(http.StatusOK)
		next.ServeHTTP(w, r)
	})
}

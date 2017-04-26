package middleware

import "net/http"

const cacheControlValue = "no-cache, no-store, private"

// NewNoCacheHandler creates a http.Handler which will add the cache-control
// headers before calling h.
func NewNoCacheHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", cacheControlValue)
		h.ServeHTTP(w, r)
	})
}

package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type statusCodeHandler int

func (code statusCodeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(int(code))
}

func Test_NoCacheHandler(t *testing.T) {
	h := NewNoCacheHandler(statusCodeHandler(http.StatusTeapot))

	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()

	h.ServeHTTP(w, req)

	cacheControl := w.Header().Get("Cache-Control")
	require.Equal(t, http.StatusTeapot, w.Code)

	assert.Contains(t, cacheControl, "no-cache")
	assert.Contains(t, cacheControl, "no-store")
	assert.Contains(t, cacheControl, "private")
}
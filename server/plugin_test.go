package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServeHTTP(t *testing.T) {
	p := Plugin{}
	p.router = p.initRouter()

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/nonexistent", nil)

	p.ServeHTTP(nil, w, r)

	result := w.Result()
	defer func() { _ = result.Body.Close() }()

	assert.Equal(t, http.StatusNotFound, result.StatusCode)
}

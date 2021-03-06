package middleware_test

import (
	"net/http"
	"testing"

	"github.com/belldata/bdx/middleware"
)

func TestLogger(t *testing.T) {
	router := setRouter()
	router.Use(middleware.Logger)
	request(router, http.MethodGet, "/", "")
}

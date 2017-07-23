package router

import (
	"net/http"
)

type Entry interface {
	Path() string
	RouterRule(req *http.Request) bool
	ServeHTTP(http.ResponseWriter, *http.Request)
}

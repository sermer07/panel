package router

import (
	"net/http"
)

func NewRouter() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/", HomeHandler)
	router.HandleFunc("/about", AboutHandler)

	return router
}

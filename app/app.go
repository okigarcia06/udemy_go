package app

import (
	"net/http"

	"some-time-api/handler"
)

func Run() error {
	http.HandleFunc("/api/time", handler.TimeHandler)
	return http.ListenAndServe(":8080", nil)
}

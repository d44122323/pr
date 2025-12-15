package main

import (
	"net/http"
)

func SetupRouter(cfg Config) http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/submit", FormHandler(cfg))

	fs := http.FileServer(http.Dir("../front"))
	mux.Handle("/", fs)

	return mux
}

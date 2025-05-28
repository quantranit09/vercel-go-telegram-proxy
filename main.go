package main

import (
	"net/http"

	"vercel-go-telegram-proxy/api"
)

func main() {
	http.HandleFunc("/", api.Listen)
	http.ListenAndServe(":8080", nil)
}

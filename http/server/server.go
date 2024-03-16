package server

import (
	"net/http"
	"simpleblog/http/server/handler"
)

// Start http server
func Run() {
	// routing
	http.HandleFunc("/", handler.Index)
	http.HandleFunc("/send", handler.Send)

	http.ListenAndServe(":8080", nil)
}
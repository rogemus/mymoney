package web

import (
	"fmt"
	"net/http"
)

func Hello() {
	print("Hello")
}

var serverPort = ":3333"

func AppServer() {
	mux := http.NewServeMux()

	// API
	mux.HandleFunc("GET /hello", GetHello)

	publicFiles := http.FileServer(http.Dir("../../ui/public"))
	mux.Handle("/", publicFiles)

	err := http.ListenAndServe(serverPort, mux)

	if err != nil {
		fmt.Println("Error while oppening the server", err)
	}
}

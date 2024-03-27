package web

import (
	"log"
	"net/http"
)

func (a *App) RunServer() {
	srv := &http.Server{
		Addr:    a.Addr,
		Handler: a.Routes(),
	}

  log.Printf("Listening on port: %v ...", a.Addr)
	err := srv.ListenAndServe()

	if err != nil {
    log.Fatal("Error while oppening the server :()")
	}
}

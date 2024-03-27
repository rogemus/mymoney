package web

import (
	"fmt"
	"net/http"
)

func (a *App) RunServer() {
	srv := &http.Server{
		Addr:    a.Addr,
		Handler: a.Routes(),
	}

	err := srv.ListenAndServe()

	if err != nil {
		fmt.Println("Error while oppening the server", err)
	}
}

package web

import (
	"fmt"
	"net/http"
	"tracker/pkg/utils"
)

func (a *App) RunServer() {
	srv := &http.Server{
		Addr:    a.Addr,
		Handler: a.Routes(),
	}

	utils.LogInfo(fmt.Sprintf("Listening on port: %v ...", a.Addr))
	err := srv.ListenAndServe()

	if err != nil {
		utils.LogFatal("Error while oppening the server :()")
	}
}

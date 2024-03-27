package main

import (
	"tracker/cmd/web"
	"tracker/pkg/models"
)

func main() {
  db := models.Database{}
  db.Connect()
  defer db.Close()

  app := &web.App{
    Addr: ":3333",
    DB: &db,
    PublicDir: "./ui/public/",
  }

  app.RunServer()
}

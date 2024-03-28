package main

import (
	"github.com/go-sql-driver/mysql"
	"os"
	"tracker/cmd/web"
	"tracker/pkg/database"
	"tracker/pkg/utils"
)

func main() {
	cfg := mysql.Config{
		User:      os.Getenv("DBUSER"),
		Passwd:    os.Getenv("DBPASS"),
		Addr:      "localhost:3309",
		DBName:    "tracker",
		ParseTime: true,
	}

	err := database.InitDB(cfg.FormatDSN())

	if err != nil {
		utils.LogFatal(err.Error())
	}

	app := &web.App{
		Addr:      ":3333",
		PublicDir: "./ui/public/",
	}

	app.RunServer()
}

package main

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"log"
	"os"
	"tracker/cmd/web"
	"tracker/pkg/models"
)

func connect(cfg string) *sql.DB {
	db, err := sql.Open("mysql", cfg)
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()

	if pingErr != nil {
		log.Fatal(pingErr)
	}

	return db
}

func main() {
	cfg := mysql.Config{
		User:      os.Getenv("DBUSER"),
		Passwd:    os.Getenv("DBPASS"),
		Addr:      "localhost:3309",
		DBName:    "tracker",
		ParseTime: true,
	}

	db := connect(cfg.FormatDSN())
	log.Println("Connected do DB!")

	app := &web.App{
		Addr:      ":3333",
		Database:  &models.Database{DB: db},
		PublicDir: "./ui/public/",
	}

	app.RunServer()
}

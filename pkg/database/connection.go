package database

import (
	"database/sql"
	"tracker/pkg/utils"
)

var db *sql.DB

func InitDB(config string) error {
	var err error
	db, err = sql.Open("mysql", config)

	if err != nil {
		return err
	}

	if err = db.Ping(); err != nil {
		return err
	}

	utils.LogInfo("Connected do Database 🎉 !!!")
	return nil
}

func GetDB() *sql.DB {
	return db
}

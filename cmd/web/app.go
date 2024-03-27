package web

import "tracker/pkg/models"

type App struct {
	Addr      string
	PublicDir string
  DB *models.Database
}

package handler

import (
	"html/template"
	"net/http"
	"tracker/pkg/utils"
)

type DashboardHandler struct {
}

func NewDashboardHandler() DashboardHandler {
	return DashboardHandler{}
}

func (h *DashboardHandler) MainView(w http.ResponseWriter, r *http.Request) {
	templ, err := template.
		New("dashboard").
		ParseFiles("ui/views/dashboard.html", "ui/views/_base.html")

	if err != nil {
		utils.LogError(err.Error())
		return
	}

	err = templ.ExecuteTemplate(w, "base", nil)

	if err != nil {
		utils.LogError(err.Error())
		return
	}
}

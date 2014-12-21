package controllers

import (
	"net/http"

	"github.com/NeilGarb/www/models"
	"github.com/NeilGarb/www/views"
)

type HomeController struct{}

func NewHomeController() *HomeController {
	controller := &HomeController{}
	return controller
}

func (self *HomeController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if r.Method == "GET" {
		self.Get(w, r)
	} else {
		http.Error(w, "Method not allowed", 405)
	}
}

func (self *HomeController) Get(w http.ResponseWriter, r *http.Request) {
	page := &struct {
		Title string
		Works []*models.Work
		Techs map[models.Tech]float64
	}{
		Title: "Home",
		Works: models.AllWorks(),
		Techs: models.WorkTechs(),
	}
	err := views.Get("home").ExecuteTemplate(w, "layout", page)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}

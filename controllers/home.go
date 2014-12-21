package controllers

import (
	"net/http"

	"github.com/NeilGarb/www/models"
	"github.com/NeilGarb/www/views"
)

var (
	allWorks  = models.AllWorks()
	workTechs = models.WorkTechs()
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
		Title       string
		Description string
		Works       []*models.Work
		Techs       map[models.Tech]float64
	}{
		Title:       "Neil Garb - Web and Software Development",
		Description: "I'm a freelance software developer based in Cape Town, South Africa, specializing in PHP, Python and Go.",
		Works:       allWorks,
		Techs:       workTechs,
	}
	err := views.Get("home").ExecuteTemplate(w, "layout", page)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}

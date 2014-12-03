package controllers

import (
	"github.com/NeilGarb/www/views"
	"net/http"
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
	}{
		Title: "Home",
	}
	err := views.Get("home").ExecuteTemplate(w, "layout", page)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}

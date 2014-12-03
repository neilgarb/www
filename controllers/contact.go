package controllers

import (
	"github.com/NeilGarb/www/views"
	"net/http"
)

type ContactController struct{}

func NewContactController() *ContactController {
	controller := &ContactController{}
	return controller
}

func (self *ContactController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if r.Method == "GET" {
		self.Get(w, r)
	} else if r.Method == "POST" {
		self.Post(w, r)
	} else {
		http.Error(w, "Method not allowed", 405)
	}
}

func (self *ContactController) Get(w http.ResponseWriter, r *http.Request) {
	page := &struct {
		Title string
	}{
		Title: "Contact",
	}
	err := views.Get("contact").ExecuteTemplate(w, "layout", page)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}

func (self *ContactController) Post(w http.ResponseWriter, r *http.Request) {
}

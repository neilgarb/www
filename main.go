package main

import (
	"flag"
	"github.com/NeilGarb/www/controllers"
	"github.com/NeilGarb/www/views"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var err error

func main() {

	// Config
	listen := flag.String("listen", ":8081", "Listen")
	viewDir := flag.String("view_dir", "./views", "View dir")
	staticDir := flag.String("static_dir", "./static", "Static dir")
	flag.Parse()

	// Views
	views.SetFolder(*viewDir)

	// Router
	router := mux.NewRouter()
	fs := http.FileServer(http.Dir(*staticDir))
	router.Handle("/", controllers.NewHomeController())
	router.Handle("/contact", controllers.NewContactController())
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	// Listen
	log.Printf("Listening on %v", *listen)
	log.Fatal(http.ListenAndServe(*listen, router))
}

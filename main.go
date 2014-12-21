package main

import (
	"flag"
	"log"
	"net/http"
	"runtime"

	"github.com/NeilGarb/www/controllers"
	"github.com/NeilGarb/www/views"
	"github.com/gorilla/mux"
)

type Config struct {
	maxprocs *int
	listen   *string
	views    *string
	static   *string
}

var err error

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// Config
	config := Config{}
	config.maxprocs = flag.Int("maxprocs", 0, "Maximum number of goprocs to use (0 for number of cores).")
	config.listen = flag.String("listen", ":8081", "Address:port to listen on")
	config.views = flag.String("views", "./views", "Directory containing tpl files")
	config.static = flag.String("static", "./static", "Directory containing static files")
	flag.Parse()

	maxprocs := runtime.NumCPU()
	if *config.maxprocs > 0 {
		maxprocs = *config.maxprocs
	}
	log.Printf("Running on %v core(s).", maxprocs)
	runtime.GOMAXPROCS(maxprocs)

	// Views
	views.SetFolder(*config.views)

	// Router
	router := mux.NewRouter()
	fs := http.FileServer(http.Dir(*config.static))
	router.Handle("/", controllers.NewHomeController())
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	// Listen
	log.Printf("Listening on %v", *config.listen)
	log.Fatal(http.ListenAndServe(*config.listen, router))
}

package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/presentation-assets/", http.StripPrefix("/presentation-assets/", http.FileServer(http.Dir("presentation/assets"))))
	mux.Handle("/app-assets/", http.StripPrefix("/app-assets/", http.FileServer(http.Dir("app/assets"))))
	mux.HandleFunc("/app", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/app" {
			http.NotFound(w, r)
			return
		}
		http.ServeFile(w, r, "app/index.html")
	})
	mux.HandleFunc("/slides", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/slides" {
			http.NotFound(w, r)
			return
		}
		http.ServeFile(w, r, "presentation/index.html")
	})
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		http.ServeFile(w, r, "presentation/index.html")
	})
	log.Println("UOB Library slides are running at http://localhost:8081/")
	log.Fatal(http.ListenAndServe(":8081", mux))
}


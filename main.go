package main

import (
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()

	str, _ := os.Getwd()
	files := http.FileServer(http.Dir(str + "/public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	bindView(mux)
	bindApi(mux)

	server := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}

	server.ListenAndServe()
}

func bindView(mux *http.ServeMux) {
	mux.HandleFunc("/", HomeView)
	mux.HandleFunc("/list", ListView)
	mux.HandleFunc("/create", CreateView)

}

func bindApi(mux *http.ServeMux) {
	mux.HandleFunc("/add", Add)
	mux.HandleFunc("/delete", Delete)
}

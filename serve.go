package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()

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

}

func bindApi(mux *http.ServeMux) {
	mux.HandleFunc("/add", Add)
	mux.HandleFunc("/delete", Delete)
}

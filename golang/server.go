package main

import (
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Hello world!")
	// })
	//
	// http.Handle("/", func())
	http.ListenAndServe(":8080", buildRouter())
}

func buildRouter() http.Handler {
	router := chi.NewRouter()
	router.Get("/test", func(wrt http.ResponseWriter, rd *http.Request) {
		wrt.Write([]byte("welcome"))
	})
	router.Get("/", func(wrt http.ResponseWriter, rd *http.Request) {
		wrt.Write([]byte("first page"))
	})
	return router
}

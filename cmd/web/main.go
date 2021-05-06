package main

import (
	"log"
	"net/http"
)

func main()  {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/ad", showAd)
	mux.HandleFunc("/ad/create", createAd)

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from ejobs clone"))
}

func showAd(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific ad..."))
}

func createAd(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new ad..."))
}

package main

import (
	"golang-server/controllers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", welcome.Index)
	http.HandleFunc("/public/video/", Video)
	http.HandleFunc("/public/static/", Static)
	err := http.ListenAndServe(":9080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func Static(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
}

func Video(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
}

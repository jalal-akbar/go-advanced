package main

import (
	"log"
	"net/http"
)

func StartNonTls() {
	mux := new(http.ServeMux)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("redirect to https://localhost/")
		http.Redirect(w, r, "https://localhost/", http.StatusTemporaryRedirect)
	})

	http.ListenAndServe(":80", mux)
}

func main() {
	go StartNonTls()

	mux := new(http.ServeMux)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	log.Println("server started at :443")
	err := http.ListenAndServeTLS(":443", "server.crt", "server.key", mux)
	if err != nil {
		panic(err)
	}
}

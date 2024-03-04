package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/cors", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "https://www.google.com")
		w.Header().Set("Access-Control-Allow-Methods", "OPTIONS,GET,POST,PUT")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-CSRF-Token")
		if r.Method == "OPTIONS" {
			w.Write([]byte("allowed"))
			return
		}
		w.Write([]byte("hello"))
	})
	// multiple origin
	http.HandleFunc("/multiplecors", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "https://www.google.com, https://novalagung.com")
		w.Header().Set("Access-Control-Allow-Methods", "OPTIONS,GET,POST,PUT")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-CSRF-Token")
		if r.Method == "OPTIONS" {
			w.Write([]byte("allowed"))
			return
		}
		w.Write([]byte("hello"))
	})
	// allow all
	http.HandleFunc("/allowallcors", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		if r.Method == "OPTIONS" {
			w.Write([]byte("allowed"))
			return
		}
		w.Write([]byte("hello"))
	})

	log.Println("Starting app at :9000")
	http.ListenAndServe(":9000", nil)
}

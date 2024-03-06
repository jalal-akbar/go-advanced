package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/rs/cors"
)

func main() {
	e := echo.New()
	// cors middleware
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{"https://www.google.com, https://novalagung.com"},
		AllowedMethods: []string{http.MethodOptions, http.MethodGet, http.MethodPost, http.MethodPut},
		AllowedHeaders: []string{"Content-Type, X-CRSF-Token"},
		Debug:          true,
	})
	e.Use(echo.WrapMiddleware(corsMiddleware.Handler))

	e.GET("/corsmiddleware", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "Hello")
	})

	e.Logger.Fatal(e.Start(":9000"))

	// http.HandleFunc("/cors", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Header().Set("Access-Control-Allow-Origin", "https://www.google.com")
	// 	w.Header().Set("Access-Control-Allow-Methods", "OPTIONS,GET,POST,PUT")
	// 	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-CSRF-Token")
	// 	if r.Method == "OPTIONS" {
	// 		w.Write([]byte("allowed"))
	// 		return
	// 	}
	// 	w.Write([]byte("hello"))
	// })
	// // multiple origin
	// http.HandleFunc("/multiplecors", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Header().Set("Access-Control-Allow-Origin", "https://www.google.com, https://novalagung.com")
	// 	w.Header().Set("Access-Control-Allow-Methods", "OPTIONS,GET,POST,PUT")
	// 	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-CSRF-Token")
	// 	if r.Method == "OPTIONS" {
	// 		w.Write([]byte("allowed"))
	// 		return
	// 	}
	// 	w.Write([]byte("hello"))
	// })
	// // allow all
	// http.HandleFunc("/allowallcors", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Header().Set("Access-Control-Allow-Origin", "*")
	// 	w.Header().Set("Access-Control-Allow-Methods", "*")
	// 	w.Header().Set("Access-Control-Allow-Headers", "*")
	// 	if r.Method == "OPTIONS" {
	// 		w.Write([]byte("allowed"))
	// 		return
	// 	}
	// 	w.Write([]byte("hello"))
	// })
	// // preflight request
	// http.HandleFunc("/preflight", func(w http.ResponseWriter, r *http.Request) {
	// 	if r.Method == "OPTIONS" {
	// 		w.Write([]byte("allowed"))
	// 		return
	// 	}
	// })

	// log.Println("Starting app at :9000")
	// http.ListenAndServe(":9000", nil)
}

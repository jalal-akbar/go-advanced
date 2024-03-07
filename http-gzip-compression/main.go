package main

import (
	"io"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Gzip())

	e.GET("/image", func(ctx echo.Context) error {
		f, err := os.Open("sample.png")
		if err != nil {
			return err
		}
		_, err = io.Copy(ctx.Response(), f)
		if err != nil {
			return err
		}
		return nil
	})

	e.Logger.Fatal(e.Start(":9000"))
	// mux := new(http.ServeMux)

	// mux.HandleFunc("/image", func(w http.ResponseWriter, r *http.Request) {
	// 	f, err := os.Open("sample.png")
	// 	if f != nil {
	// 		defer f.Close()
	// 	}
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}

	// 	_, err = io.Copy(w, f)
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}
	// })

	// server := new(http.Server)
	// server.Addr = ":9000"
	// server.Handler = gziphandler.GzipHandler(mux)

	// server.ListenAndServe()
}

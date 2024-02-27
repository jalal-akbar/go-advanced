package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	r := echo.New()
	r.GET("/", func(ctx echo.Context) error {
		data := "Welcome"
		return ctx.String(http.StatusOK, data)
	})
	if err := r.Start(":9000"); err != nil {
		fmt.Println(err.Error())
	}
}

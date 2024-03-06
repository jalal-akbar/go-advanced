package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type M map[string]interface{}

func main() {
	tmpl := template.Must(template.ParseGlob("./*.html"))

	e := echo.New()

	const CSRFTokeHeader = "X-CSRF-Token"
	const CSRFKey = "csrf"

	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "header:" + CSRFTokeHeader,
		ContextKey:  CSRFKey,
	}))
	e.GET("/index", func(ctx echo.Context) error {
		data := make(M)
		data[CSRFKey] = ctx.Get(CSRFKey)
		return tmpl.Execute(ctx.Response(), data)
	})
	e.POST("/sayhello", func(ctx echo.Context) error {
		data := make(M)
		if err := ctx.Bind(&data); err != nil {
			return err
		}
		message := fmt.Sprintf("hello %s", data["name"])
		return ctx.JSON(http.StatusOK, message)
	})

	e.Logger.Fatal(e.Start(":9000"))
}

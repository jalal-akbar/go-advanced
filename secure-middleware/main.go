package main

import (
	"net/http"

	"github.com/unrolled/secure"

	"github.com/labstack/echo"
)

func main() {
	secureMiddleware := secure.New(secure.Options{
		BrowserXssFilter:        true,
		ContentTypeNosniff:      true,
		FrameDeny:               true,
		CustomFrameOptionsValue: "SAMEORIGIN",
		AllowedHosts:            []string{"localhost:9000", "www.google.com"},
	})
	e := echo.New()
	e.Use(echo.WrapMiddleware(secureMiddleware.Handler))

	e.GET("/index", func(ctx echo.Context) error {
		ctx.Response().Header().Set("Access-Control-Allow-Origin", "*")
		return ctx.String(http.StatusOK, "Hello")
	})

	e.Logger.Fatal(e.StartTLS(":9000", "server.crt", "server.key"))
}

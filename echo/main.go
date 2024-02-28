package main

import (
	payload "echo/4-parsing-http-request-payload"
	validation "echo/5-http-request-payload-validation"
	rendering "echo/7-template-rendering"
	ml "echo/8-middleware-and-logging"

	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	// 7 template rendering
	e.Renderer = rendering.NewRender("./*.html", true)
	//validate
	e.Validator = &validation.CustomValidator{Validator: validator.New()}
	// 6.HTTP Error Handling
	e.HTTPErrorHandler = func(err error, ctx echo.Context) {
		report, ok := err.(*echo.HTTPError)
		if !ok {
			report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		ctx.Logger().Error(report)
		ctx.JSON(report.Code, report)
	}
	// middleware

	e.Use(ml.MiddlewareOne)
	e.Use(ml.MiddlewareTwo)
	e.Use(echo.WrapMiddleware(ml.MiddlewareSomething)) // non-echo middleware
	// logging middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{ // logging middleware
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(ml.MiddlewareLogging)
	e.HTTPErrorHandler = ml.ErrorHandler
	//routing
	e.Any("/user", func(ctx echo.Context) error { // Payload
		u := new(payload.UserForPayload)
		if err := ctx.Bind(u); err != nil {
			return ctx.JSON(http.StatusNotFound, u)
		}
		return ctx.JSON(http.StatusOK, u)
	})
	e.POST("/users", func(ctx echo.Context) error { // Validate
		u := new(validation.UserForValidate)
		if err := ctx.Bind(u); err != nil {
			return err
		}
		if err := ctx.Validate(u); err != nil {
			return err
		}
		return ctx.JSON(http.StatusOK, true)
	})
	e.GET("/index", func(ctx echo.Context) error { // Template Renderer
		data := rendering.M{"pesan": "hallo jalal"}
		return ctx.Render(http.StatusOK, "index.html", data)
	})
	e.GET("/middleware", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, true)
	})
	e.GET("/logging", func(ctx echo.Context) error { // logging middleware
		return ctx.JSON(http.StatusOK, true)
	})
	lock := make(chan error)
	go func(lock chan error) { lock <- e.Start(":9000") }(lock)

	time.Sleep(1 * time.Millisecond)
	ml.MakeLogEntry(nil).Warning("application started withour ssl/tls enabled")

	if err := <-lock; err != nil {
		ml.MakeLogEntry(nil).Panic("failed to start application")
	}

	// fmt.Println("server started at :9000")

	// e.Logger.Fatal(e.Start(":9000"))
}

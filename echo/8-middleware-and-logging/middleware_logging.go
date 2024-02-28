package middlewareandlogging

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

// middleware
func MiddlewareOne(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		fmt.Println("from middleware one")
		return next(ctx)
	}
}

func MiddlewareTwo(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		fmt.Println("from middleware two")
		return next(ctx)
	}
}

func MiddlewareSomething(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("from middleware something")
		next.ServeHTTP(w, r)
	})
}

// logger
func MakeLogEntry(ctx echo.Context) *log.Entry {
	if ctx == nil {
		return log.WithFields(log.Fields{
			"at": time.Now().Format("2006-01-02 15:04:05"),
		})
	}
	return log.WithFields(log.Fields{
		"at":     time.Now().Format("2006-01-02 15:04:05"),
		"method": ctx.Request(),
		"uri":    ctx.Request().URL.String(),
		"ip":     ctx.Request().RemoteAddr,
	})
}

func MiddlewareLogging(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		MakeLogEntry(ctx).Info("incoming request")
		return next(ctx)
	}
}

func ErrorHandler(err error, ctx echo.Context) {
	report, ok := err.(*echo.HTTPError)

	if ok {
		report.Message = fmt.Sprintf("http error %d - %v", report.Code, report.Message)
	} else {
		report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	MakeLogEntry(ctx).Error(report.Message)
	ctx.HTML(report.Code, report.Message.(string))
}

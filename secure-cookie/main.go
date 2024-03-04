package main

import (
	"net/http"
	"time"

	"github.com/gorilla/securecookie"
	"github.com/labstack/echo"
	gubrak "github.com/novalagung/gubrak/v2"
)

type M map[string]interface{}

var sc = securecookie.New([]byte("very-secret"), []byte("a-lot-secret-yay"))

func setCookie(ctx echo.Context, name string, data M) error {
	encoded, err := sc.Encode(name, data)
	if err != nil {
		panic(err)
	}

	cookie := &http.Cookie{
		Name:     name,
		Value:    encoded,
		Secure:   true,
		HttpOnly: false,
		Path:     "/",
		Expires:  time.Now().Add(1 * time.Hour),
	}

	http.SetCookie(ctx.Response(), cookie)

	return nil
}

func getCookie(ctx echo.Context, name string) (M, error) {
	cookie, err := ctx.Request().Cookie(name)
	data := M{}

	if err == nil {
		if err = sc.Decode(name, cookie.Value, &data); err == nil {
			return data, nil
		}
	}

	return nil, err
}

func removeCookie(ctx echo.Context, name string) {
	cookie := &http.Cookie{}
	cookie.Name = name
	cookie.Path = "/"
	cookie.MaxAge = -1
	cookie.Expires = time.Unix(0, 0)
	http.SetCookie(ctx.Response(), cookie)
}

func main() {
	const cookieName = "data"
	e := echo.New()

	e.GET("/securecookie", func(ctx echo.Context) error {
		data, err := getCookie(ctx, cookieName)
		if err != nil && err != http.ErrNoCookie && err != securecookie.ErrMacInvalid {
			return err
		}
		if data == nil {
			data = M{"Message": "Hello", "ID": gubrak.RandomString(32)}
			err = setCookie(ctx, cookieName, data)
			if err != nil {
				return err
			}
		}
		return ctx.JSON(http.StatusOK, data)
	})
	e.GET("/deletecookie", func(ctx echo.Context) error {
		removeCookie(ctx, cookieName)
		return ctx.Redirect(http.StatusTemporaryRedirect, "/securecookie")
	})

	e.Logger.Fatal(e.Start(":9000"))

}

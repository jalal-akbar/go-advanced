package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/globalsign/mgo"
	"github.com/gorilla/context"
	"github.com/kidstuff/mongostore"
	"github.com/labstack/echo"
)

const SESSION_ID = "id"

// MongoDB Store
func newMongoStore() *mongostore.MongoStore {
	mgoSession, err := mgo.Dial("localhost:27017")
	if err != nil {
		log.Println("ERROR", err)
		os.Exit(0)
	}

	dbCollection := mgoSession.DB("learnwebgolang").C("session")
	maxAge := 86400 * 7
	ensureTTL := true
	authKey := []byte("my-auth-key-very-secret")
	encryptionKey := []byte("my-encryption-key-very-secret123")

	store := mongostore.NewMongoStore(dbCollection, maxAge, ensureTTL, authKey, encryptionKey)

	return store
}

// Postgres SQL Store
// func newPostgresStore() *pgstore.PGStore {
// 	url := "postgres://jalalakbar:@127.0.0.1:5432/jalalakbar?sslmode=disable"
// 	authKey := []byte("my-auth-key-very-secret")
// 	encryptionKey := []byte("my-encryption-key-very-secret123")
// 	store, err := pgstore.NewPGStore(url, authKey, encryptionKey)
// 	if err != nil {
// 		log.Println("ERROR", err)
// 		os.Exit(0)
// 	}
// 	return store
// }

// // Secure Cookie Store
// func newCookieStore() *sessions.CookieStore {
// 	authKey := []byte("my-auth-key-very-secret")
// 	encryptionKey := []byte("my-encryption-key-very-secret123")

// 	store := sessions.NewCookieStore(authKey, encryptionKey)
// 	store.Options.Path = "/"
// 	store.Options.MaxAge = 86400 * 7
// 	store.Options.HttpOnly = true

// 	return store
// }

func main() {
	store := newMongoStore()
	e := echo.New()
	// session store and context clear handler
	e.Use(echo.WrapMiddleware(context.ClearHandler))
	// routing
	e.GET("setsession", func(ctx echo.Context) error {
		// create SessionData
		session, _ := store.Get(ctx.Request(), SESSION_ID)
		session.Values["message1"] = "Hello"
		session.Values["message2"] = "World"
		session.Save(ctx.Request(), ctx.Response())

		return ctx.Redirect(http.StatusTemporaryRedirect, "/getsession")
	})
	e.GET("getsession", func(ctx echo.Context) error {
		// access SessionData from object
		session, _ := store.Get(ctx.Request(), SESSION_ID)

		if len(session.Values) == 0 {
			return ctx.String(http.StatusOK, "Empty Result")
		}

		return ctx.String(http.StatusOK, fmt.Sprintf(
			"%s %s",
			session.Values["message1"],
			session.Values["message2"],
		))
	})
	e.GET("/deletesession", func(ctx echo.Context) error {
		// delete SessionData from object
		session, _ := store.Get(ctx.Request(), SESSION_ID)
		session.Options.MaxAge = -1 // -1 = expired
		session.Save(ctx.Request(), ctx.Response())

		return ctx.Redirect(http.StatusTemporaryRedirect, "/getsession")
	})
	e.Logger.Fatal(e.Start(":9000"))
}

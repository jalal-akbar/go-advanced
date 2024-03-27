package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	gubrak "github.com/novalagung/gubrak/v2"
)

type CustomMux struct {
	http.ServeMux
	middlewares []func(next http.Handler) http.Handler
}

func (c *CustomMux) RegisterMiddleware(next func(next http.Handler) http.Handler) {
	c.middlewares = append(c.middlewares, next)
}

func (c *CustomMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var current http.Handler = &c.ServeMux

	for _, next := range c.middlewares {
		current = next(current)
	}
	current.ServeHTTP(w, r)
}

type key any

var userInfoKey key = "userInfo"

type M map[string]interface{}

var (
	APPLICATION_NAME          = "My Simple JWT App"
	LOGIN_EXPIRATION_DURATION = time.Duration(1) * time.Hour
	JWT_SIGNING_METHOD        = jwt.SigningMethodHS256
	JWT_SIGNATURE_KEY         = []byte("the secret")
)

type MyClaims struct {
	jwt.RegisteredClaims
	Username string `json:"username"`
	Email    string `json:"email"`
	Group    string `json:"group"`
}

func main() {
	mux := &CustomMux{}
	mux.RegisterMiddleware(MiddlewareJWTAuthorization)

	mux.HandleFunc("/index", HandlerIndex)
	mux.HandleFunc("/login", HandlerLogin)

	server := new(http.Server)
	server.Addr = ":8080"
	server.Handler = mux

	fmt.Println("starting server at ", server.Addr)
	server.ListenAndServe()
}

func HandlerLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "unsupported method", http.StatusBadRequest)
		return
	}

	username, password, ok := r.BasicAuth()
	if !ok {
		http.Error(w, "invalid username or password", http.StatusBadRequest)
		return
	}

	ok, userInfo := authenticateUser(username, password)
	if !ok {
		http.Error(w, "invalid username or password", http.StatusBadRequest)
		return
	}

	claims := MyClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    APPLICATION_NAME,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(LOGIN_EXPIRATION_DURATION)),
		},
		Username: userInfo["username"].(string),
		Email:    userInfo["email"].(string),
		Group:    userInfo["group"].(string),
	}
	token := jwt.NewWithClaims(JWT_SIGNING_METHOD, claims)

	signedToken, err := token.SignedString(JWT_SIGNATURE_KEY)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tokenStrig, _ := json.Marshal(M{"token": signedToken})
	w.Write([]byte(tokenStrig))

}

func authenticateUser(username, password string) (bool, M) {
	basePath, _ := os.Getwd()
	dbPath := filepath.Join(basePath + "./users.json")
	buf, _ := os.ReadFile(dbPath)

	data := make([]M, 0)
	if err := json.Unmarshal(buf, &data); err != nil {
		return false, nil
	}

	res := gubrak.From(data).Find(func(each M) bool {
		return each["username"] == username && each["password"] == password
	}).Result()
	if res != nil {
		resM := res.(M)
		delete(resM, "password")
		return true, resM
	}
	return false, nil
}

func MiddlewareJWTAuthorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/login" {
			next.ServeHTTP(w, r)
			return
		}
		authorizationHeader := r.Header.Get("Authorization")
		if !strings.Contains(authorizationHeader, "Bearer") {
			http.Error(w, "invalid token", http.StatusBadRequest)
			return
		}

		tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if method, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("signing method invalid")
			} else if method != JWT_SIGNING_METHOD {
				return nil, fmt.Errorf("signing method invalid")
			}
			return JWT_SIGNATURE_KEY, nil
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(context.Background(), userInfoKey, claims)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

func HandlerIndex(w http.ResponseWriter, r *http.Request) {
	userInfo := r.Context().Value(userInfoKey).(jwt.MapClaims)
	message := fmt.Sprintf("hello %s (%s)", userInfo["username"], userInfo["group"])
	w.Write([]byte(message))
}

// curl -X POST --user noval:kaliparejaya123 http://localhost:8080/login
// {"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTEzOTg1MDgsImlzcyI6Ik15IFNpbXBsZSBKV1QgQXBwIiwidXNlcm5hbWUiOiJqYWxhbCIsImVtYWlsIjoiamFsYWxAZ21haWwuY29tIiwiZ3JvdXAiOiJhZG1pbiJ9.ahjZkoX6acvRxmV0NaPic76EA47L8ktKFAb-5Wm7_5o"}

// curl -X GET \
//  --header "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTEzOTg1MDgsImlzcyI6Ik15IFNpbXBsZSBKV1QgQXBwIiwidXNlcm5hbWUiOiJqYWxhbCIsImVtYWlsIjoiamFsYWxAZ21haWwuY29tIiwiZ3JvdXAiOiJhZG1pbiJ9.ahjZkoX6acvRxmV0NaPic76EA47L8ktKFAb-5Wm7_5o" \
// http://localhost:8080/index

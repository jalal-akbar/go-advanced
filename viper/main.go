package main

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/labstack/echo"
	"github.com/spf13/viper"
)

func main() {
	e := echo.New()
	// // json
	// viper.SetConfigType("json")
	// viper.AddConfigPath(".")
	// viper.SetConfigName("app.config")
	// if err := viper.ReadInConfig(); err != nil {
	// 	e.Logger.Fatal(err)
	// }
	// // ymal
	// viper.SetConfigType("ymal")
	// viper.AddConfigPath(".")
	// viper.SetConfigName("app.config")
	// if err := viper.ReadInConfig(); err != nil {
	// 	e.Logger.Fatal(err)
	// }

	// viper.WatchConfig()
	// viper.OnConfigChange(func(in fsnotify.Event) {
	// 	fmt.Println("Config file changed", in.Name)
	// })
	// env
	configAppName := os.Getenv("APP_NAME")
	if configAppName == "" {
		e.Logger.Fatal("APP_NAME config is required")
	}
	configServerPort := os.Getenv("SERVER_PORT")
	if configServerPort == "" {
		e.Logger.Fatal("SERVER_PORT config is required")
	}

	e.GET("/viper", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, true)
	})

	server := new(http.Server)
	server.Addr = ":" + configServerPort
	if configServerReadTimeout := os.Getenv("SERVER_READ_TIMEOUT_IN_MINUTE"); configServerReadTimeout != "" {
		duration, _ := strconv.Atoi(configServerReadTimeout)
		server.ReadTimeout = time.Duration(duration) * time.Minute
	}
	if configServerWriteTimeout := os.Getenv("SERVER_WRITE_TIMEOUT_IN_MINUTE"); configServerWriteTimeout != "" {
		duration, _ := strconv.Atoi(configServerWriteTimeout)
		server.ReadTimeout = time.Duration(duration) * time.Minute
	}

	e.Logger.Print("Starting ", viper.GetString("appName"))
	e.Logger.Fatal(e.Start(":" + viper.GetString("server.port")))
}

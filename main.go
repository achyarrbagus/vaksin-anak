package main

import (
	"fmt"
	"kartu-vaksin/pkg/mysql"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("failed to load env")
	}
	var port = os.Getenv("PORT")

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PATCH, echo.DELETE},
		AllowHeaders: []string{"X-Requested-With", "Content-Type", "Authorization"},
	}))

	fmt.Println(port)

	mysql.DatabaseInit()

	// e.Static("/uploads", "./uploads")
	fmt.Println("server running localhost: " + port)
	e.Logger.Fatal(e.Start("localhost:" + port)) // delete localhost
}

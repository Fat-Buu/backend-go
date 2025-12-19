package main

import (
	"log"

	_ "github.com/backend-go/docs"
	"github.com/backend-go/internal/app"
)

// @title Fiber User API
// @version 1.0
// @description Fiber CRUD API example
// @host localhost:3000
// @BasePath /backend-go/v1
func main() {
	app.LoadEnv()

	port := app.GetPort()
	app := app.SetupFiberApp()

	log.Fatal(app.Listen(":" + port))
}

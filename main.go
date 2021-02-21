package main

import (
	"fiber-gorm/database"
	"fiber-gorm/handler"
	"fiber-gorm/model"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/user", handler.GetUsers)
	app.Get("/api/user/:id", handler.GetUser)
	app.Post("/api/user", handler.CreateUser)
	app.Patch("/api/user/:id", handler.UpdateUser)
	app.Delete("/api/user/:id", handler.DeleteUser)
}

func main() {
	db, err := database.GetDBConn()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to Database")
	db.AutoMigrate(&model.User{})
	fmt.Println("Database Migrated")

	app := fiber.New()
	app.Use(cors.New())

	setupRoutes(app)

	log.Fatal(app.Listen(":5000"))
}

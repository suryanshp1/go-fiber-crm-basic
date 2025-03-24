package main

import (
	"github.com/suryanshp1/go-fiber-crm-basic/database"
	"github.com/gofiber/fiber"
	"github.com/suryanshp1/go-fiber-crm-basic/lead"
	"github.com/jinzhu/gorm"
	"fmt"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/leads", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead", lead.NewLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "leads.db")
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Database connection successfully opened")
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database Migrated")
}

func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	app.Listen(3000)

	defer database.DBConn.Close()
}
package main

import (
	"fmt"

	"github.com/go-fiber-crm-basic/database"
	"github.com/go-fiber-crm-basic/lead"
	"github.com/gofiber/fiber"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead", lead.NewLead)
	app.Delete("api/v1/lead/id", lead.DeleteLead)

}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "leads.db")
	if err != nil {
		panic("failed to connect with database")
	}
	fmt.Println("Connection opend to database")
	database.DBConn.autoMigrate(&lead.Lead{})
	fmt.Println("database Migrated")
}

func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	app.Listen(3000)
	defer database.DBConn.Close()
}

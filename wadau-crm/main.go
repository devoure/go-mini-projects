package main

import (
	"fmt"

	"github.com/devoure/wadau-crm/data"
	"github.com/devoure/wadau-crm/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func initDatabase() {
	var err error
	database.DB, err = gorm.Open("sqlite3", "customersData.db")
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println(">>> Connected to SQLite Database")
	database.DB.AutoMigrate(&data.Data{})
	fmt.Println(">>> Database Migrated")
}

func setupRoutes(app *fiber.App) {
	app.Get("/wadau/v1/customers/", data.GetAllCustomersData)
	app.Get("/wadau/v1/customers/:id", data.GetCustomerData)
	app.Post("/wadau/v1/customers/add", data.AddCustomerRecord)
	app.Delete("/wadau/v1/customers/delete/:id", data.DeleteCustomerRecord)
	app.Put("/wadau/v1/customers/update/:id", data.UpdateCustomerRecord)
}

func main() {
	app := fiber.New()
	fmt.Println(">>>> Initializing Database")
	initDatabase()
	setupRoutes(app)
	fmt.Println(">>>> Starting WADAU CRM server at port 3000: ")
	app.Listen(3000)
	defer database.DB.Close()
}

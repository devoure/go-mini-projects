package data

import (
	"github.com/devoure/wadau-crm/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

type Data struct {
	gorm.Model
	Name       string `json:"name"`
	CustomerID string `json:"customer_id"`
	Location   string `json:"location"`
	Purchases  int    `json:"purchases"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
}

func GetAllCustomersData(c *fiber.Ctx) {
	db := database.DB
	var data []Data
	db.Find(&data)
	c.JSON(data)
}

func GetCustomerData(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DB
	var data Data
	db.Find(&data, id)
	c.JSON(data)
}

func AddCustomerRecord(c *fiber.Ctx) {
	db := database.DB
	var data Data
	err := c.BodyParser(&data)
	if err != nil {
		c.Status(503).Send(err)
		return
	}
	db.Create(&data)
	c.JSON(data)
}

func DeleteCustomerRecord(c *fiber.Ctx) {
	db := database.DB
	id := c.Params("id")
	var data Data
	db.First(&data, id)
	if data.CustomerID == "" {
		c.Status(500).Send("Customer Record to be Deleted NOT FOUND")
		return
	}
	db.Delete(&data)
	c.Send("Customer Record Successfully DELETED")
}

func UpdateCustomerRecord(c *fiber.Ctx) {
	db := database.DB
	id := c.Params("id")
	var data Data
	db.First(&data, id)
	if data.CustomerID == "" {
		c.Status(500).Send("Customer Record to be Updated NOT FOUND")
		return
	}
	db.Delete(&data)
	var newData Data
	err := c.BodyParser(&newData)
	if err != nil {
		c.Status(503).Send(err)
		return
	}
	db.Create(&newData)
	c.Send("Customer Record Successfully UPDATED")
}

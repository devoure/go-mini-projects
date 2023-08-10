package main

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

var mg MongoInstance

const dbName = "wadau-hrms"
const mongoURI = "mongodb://devoure:password4devoure@localhost:27017/" + dbName

type Employee struct {
	// when field is empty then the field will be ignored
	ID     string  `json:"id,omitempty" bson:"_id,omitempty"`
	Name   string  `json:"name"`
	Salary float64 `json:"salary"`
	Age    float64 `json:"age"`
}

func ConnectDB() error {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	db := client.Database(dbName)

	if err != nil {
		return err
	}
	mg = MongoInstance{
		Client: client,
		Db:     db,
	}
	return nil
}

func getEmployeeData(c *fiber.Ctx) error {

	query := bson.D{{}}

	cursor, err := mg.Db.Collection("employees").Find(c.Context(), query)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	var employees []Employee
	err = cursor.All(c.Context(), &employees)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(employees)
}

func addNewEmployee(c *fiber.Ctx) error {
	collection := mg.Db.Collection("employees")

	employee := new(Employee)

	if err := c.BodyParser(employee); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	employee.ID = ""

	insertionResult, err := collection.InsertOne(c.Context(), employee)

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	filter := bson.D{{Key: "_id", Value: insertionResult.InsertedID}}
	createdRecord := collection.FindOne(c.Context(), filter)

	createdEmployee := &Employee{}
	createdRecord.Decode(createdEmployee)

	return c.Status(201).JSON(createdEmployee)

}

func main() {
	err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	app := fiber.New()

}

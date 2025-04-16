package handlers

import (
	"context"
	"fmt"

	"fiber-app/database"
	"fiber-app/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func GetEmployees(c *fiber.Ctx) error {
	
	collection := database.GetCollection("employees")

	
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		fmt.Println("Error fetching employees from MongoDB:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch employees",
		})
	}
	defer cursor.Close(context.TODO())

	
	var employees []models.Employees
	if err := cursor.All(context.TODO(), &employees); err != nil {
		fmt.Println("Error decoding employees:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to decode employees",
		})
	}

	for _, emp := range employees {
		if emp.Department == "" {
			fmt.Printf("Employee missing department: %+v\n", emp)
		}
	}


	fmt.Printf("Fetched Employees: %+v\n", employees)


	return c.JSON(fiber.Map{
		"message": "Employees retrieved successfully",
		"data":    employees,
	})
}


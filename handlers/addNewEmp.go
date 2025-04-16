

package handlers

import (
	"context"
	"fmt"

	"fiber-app/database"
	"fiber-app/models"

	"github.com/gofiber/fiber/v2"
)

func SaveEmployee(c *fiber.Ctx) error {
	var employee models.Employees

	// Parse request body
	if err := c.BodyParser(&employee); err != nil {
		fmt.Println("Error parsing request body:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request format",
		})
	}

	// Debug: Print received data before inserting
	fmt.Printf("Received Employee Data: %+v\n", employee)

	// Ensure MongoDB connection is active
	collection := database.GetCollection("employees")

	// Insert employee into MongoDB
	result, err := collection.InsertOne(context.TODO(), employee)
	if err != nil {
		fmt.Println("Error inserting into MongoDB:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to save employee",
		})
	}

	// Debug: Print inserted ID
	fmt.Println("Inserted Employee ID:", result.InsertedID)

	// Success response
	return c.JSON(fiber.Map{
		"message": "Employee saved successfully",
		"id":      result.InsertedID,
	})
}



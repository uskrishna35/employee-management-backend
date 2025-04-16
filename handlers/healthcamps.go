package handlers

import (
	"github.com/gofiber/fiber/v2"
)

// HealthCamp struct
type HealthCamp struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// GetHealthCamps handles POST requests to return health camps
func GetHealthCamps(c *fiber.Ctx) error {
	// Sample data (Replace this with a database query)
	healthCamps := []HealthCamp{
		{ID: 1, Name: "Magnolia Health Camp"},
		{ID: 2, Name: "Kriyatec Health Camp"},
	}

	return c.JSON(healthCamps) // Return health camps list
}

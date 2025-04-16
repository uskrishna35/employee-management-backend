package handlers

import (
	"github.com/gofiber/fiber/v2"
)

// Organization struct
type Organization struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// GetOrganizations returns a list of organizations
func PostOrganizations(c *fiber.Ctx) error {
	organizations := []Organization{
		{ID: 1, Name: "Magnolia Community Health"},
		{ID: 2, Name: "Kriyatec"},
	}

	return c.JSON(organizations)
}

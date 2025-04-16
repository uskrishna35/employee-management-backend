package handlers

import (
	"github.com/gofiber/fiber/v2"
)

// Facility struct
type Facility struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	OrganizationID int    `json:"organization_id"` // Foreign key reference
}

// GetFacilities returns facilities based on organization
func PostFacilities(c *fiber.Ctx) error {
	orgID := c.Query("organization_id") // Get organization ID from query params

	// Sample data (replace with database query later)
	allFacilities := []Facility{
		{ID: 1, Name: "Magnolia Facility 1", OrganizationID: 1},
		{ID: 2, Name: "Magnolia Facility 2", OrganizationID: 1},
		{ID: 3, Name: "Kriyatec Facility 1", OrganizationID: 2},
	}

	// Filter facilities based on the selected organization
	var filteredFacilities []Facility
	for _, facility := range allFacilities {
		if orgID == "" || string(rune(facility.OrganizationID)) == orgID {
			filteredFacilities = append(filteredFacilities, facility)
		}
	}

	return c.JSON(filteredFacilities)
}

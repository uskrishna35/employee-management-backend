package handlers

import (
	"context"
	"fmt"
	"time"

	"fiber-app/database"
	"fiber-app/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	
)

// SaveSelection godoc
// @Summary Save a selection
// @Description Store selection data (organization, facility, camp, employee, department)
// @Tags Selections
// @Accept  json
// @Produce  json
// @Param selection body models.Selection true "Selection Data"
// @Success 200 {object} fiber.Map
// @Failure 400 {object} fiber.Map
// @Router /selections [post]




func SaveSelection(c *fiber.Ctx) error {
	// Step 1: Parse raw request body into a generic map
	var data map[string]interface{}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	fmt.Printf("Raw Request Data: %+v\n", data)

    // Step 6: Extract Organization
    var organizationID, organizationName string
    if orgData, ok := data["org_id"].(map[string]interface{}); ok {
        if id, ok := orgData["id"]; ok {
            organizationID = fmt.Sprintf("%v", id)
        }
        if name, ok := orgData["name"]; ok {
            organizationName = fmt.Sprintf("%v", name)
        }
    }
    
	// Step 3: Extract Facility
	var facilityID, facilityName string
	if facility, ok := data["facility_id"].(map[string]interface{}); ok {
		if id, ok := facility["id"]; ok {
			facilityID = fmt.Sprintf("%v", id)
		}
		if name, ok := facility["name"]; ok {
			facilityName = fmt.Sprintf("%v", name)
		}
	}

	// Step 4: Extract Camp
	var campID, campName string
	if camp, ok := data["camp_id"].(map[string]interface{}); ok {
		if id, ok := camp["id"]; ok {
			campID = fmt.Sprintf("%v", id)
		}
		if name, ok := camp["name"]; ok {
			campName = fmt.Sprintf("%v", name)
		}
	}

	// Step 5: Extract Department
	var departmentID, departmentName string
	if data["department"] != nil {
		departmentID = fmt.Sprintf("%v", data["department"])
		departmentName = models.GetDepartmentNameByID(departmentID)
	}

	// Step 6: Extract Employee
	empID := fmt.Sprintf("%v", data["emp_id"])
	firstName := fmt.Sprintf("%v", data["first_name"])
	lastName := fmt.Sprintf("%v", data["last_name"])
	employeeName := firstName + " " + lastName

	// Step 7: Build Selection struct
	selection := models.Selection{
		OrganizationID:   organizationID,
		OrganizationName: organizationName,
		FacilityID:       facilityID,
		FacilityName:     facilityName,
		CampID:           campID,
		CampName:         campName,
		EmployeeID:       empID,
		EmployeeName:     employeeName,
		DepartmentID:     departmentID,
		DepartmentName:   departmentName,
		CreatedAt:        primitive.NewDateTimeFromTime(time.Now()),
	}

	// Step 8: Insert into MongoDB
	collection := database.DB.Collection("selections")
	res, err := collection.InsertOne(context.TODO(), selection)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to insert document",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Selection saved successfully",
		"id":      res.InsertedID,
	})
}

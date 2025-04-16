package handlers

import (
	"context"
	"fmt"

	"fiber-app/database" // Import the database package

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func DeleteInvalidSelections(c *fiber.Ctx) error {
	collection := database.DB.Collection("selections")

	filter := bson.M{
		"$or": []bson.M{
			{"employee_id": bson.M{"$exists": false}},
			{"employee_id": ""},
		},
	}

	res, err := collection.DeleteMany(context.Background(), filter)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete invalid selections",
		})
	}

	return c.JSON(fiber.Map{
		"message": fmt.Sprintf("Deleted %d invalid selection records", res.DeletedCount),
	})
}

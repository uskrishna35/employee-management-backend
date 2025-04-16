package handlers

import (
	"context"
	"fmt"

	"fiber-app/database"
	"fiber-app/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func GetSelections(c *fiber.Ctx) error {
	collection := database.DB.Collection("selections")

	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch selections",
		})
	}
	defer cursor.Close(context.TODO())

	var selections []models.Selection
	if err := cursor.All(context.TODO(), &selections); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to parse selections",
		})
	}

	fmt.Printf("Fetched selections: %+v\n", selections) // for debugging

	return c.JSON(selections)
}

func GetALLSelectionsWithQuery(c *fiber.Ctx) error {
	collection := database.DB.Collection("selections")


	pipeline := bson.A{
		bson.D{{"$match", bson.D{{"employee_id", bson.D{{"$exists", true}}}}}},
		bson.D{
			{"$lookup",
				bson.D{
					{"from", "employees"},
					{"localField", "employee_id"},
					{"foreignField", "emp_id"},
					{"as", "result"},
				},
			},
		},
		bson.D{
			{"$unwind",
				bson.D{
					{"path", "$result"},
					{"preserveNullAndEmptyArrays", true},
				},
			},
		},
		bson.D{
			{"$set",
				bson.D{
					{"first_name", "$result.first_name"},
					{"last_name", "$result.last_name"},
					{"employeeName",
						bson.D{
							{"$concat",
								bson.A{
									"$result.first_name",
									"$result.last_name",
								},
							},
						},
					},
				},
			},
		},
	}

	cursor, err := collection.Aggregate(context.Background(), pipeline)
	if err != nil {
		return fiber.NewError(500, err.Error())
	}
	defer cursor.Close(context.Background())

	var results []bson.M
	if err := cursor.All(context.Background(), &results); err != nil {
		return fiber.NewError(500, err.Error())
	}

	// Optional: Print the results for debugging
	fmt.Printf("Fetched selections: %+v\n", results)

	return c.JSON(results)
}

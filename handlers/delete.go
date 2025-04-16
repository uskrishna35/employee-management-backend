package handlers

import (
    "context"
    "time"

    "github.com/gofiber/fiber/v2"
    "go.mongodb.org/mongo-driver/bson"
    "fiber-app/database"
)

func DeleteSelection(c *fiber.Ctx) error {
    empID := c.Params("employee_id")
    if empID == "" {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "message": "emp_id is required",
        })
    }

    collection := database.GetCollection("selections") // use your DB helper
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    filter := bson.M{"employee_id": empID}

    result, err := collection.DeleteOne(ctx, filter)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "message": "Failed to delete record",
            "error":   err.Error(),
        })
    }

    if result.DeletedCount == 0 {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "message": "No record found with that emp_id",
        })
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "message": "Record deleted successfully",
    })
}

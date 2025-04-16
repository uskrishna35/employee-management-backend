package handlers

import (

	"github.com/gofiber/fiber/v2"
)

type department struct {
	ID   int `json:"id"`
	Name string `json:"name"`
	}
	func GetDepartment( c *fiber.Ctx) error{
		department := []department{
			{ID: 1, Name: "QA"},
			{ID: 2, Name: "Kriyatec"},
		}

		

		return c.JSON( department )

	}
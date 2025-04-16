package routes

import (
	"fiber-app/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/login", handlers.Login)
	api.Post("/register", handlers.Register)
	api.Get("/agg", handlers.GetALLSelectionsWithQuery)
	api.Post("/organizations", handlers.PostOrganizations)
	api.Post("/facilities", handlers.PostFacilities)
	app.Post("/api/healthcamps", handlers.GetHealthCamps)
	app.Post("/api/department", handlers.GetDepartment)
	app.Post("/api/employees", handlers.SaveEmployee)
	app.Get("/api/employees", handlers.GetEmployees)
	app.Post("/api/selection", handlers.SaveSelection)
	app.Get("/api/selections", handlers.GetSelections)
	app.Delete("/api/selections_delete/:employee_id", handlers.DeleteSelection)
	app.Delete("/api/selections/delete-invalid", handlers.DeleteInvalidSelections)
}

package routes

import (
	"world-api/controllers"
	"world-api/services"
	"world-api/utils"

	"github.com/gofiber/fiber/v2"
)

func ConfigRoutes(router *fiber.App) *fiber.App {

	main := router.Group("api/v1")

	countryController := controllers.CountryControllerImpl{
		CountriesAPI: &services.CountriesAPIImpl{
			MakeRequest: utils.CountryAPIMakeRequest,
		},
	}

	main.Get("/country/:id", countryController.HandleSearch)

	return router
}

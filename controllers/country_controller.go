package controllers

import (
	"world-api/services"

	"github.com/gofiber/fiber/v2"
)

type CountryController interface {
	HandleSearch(c *fiber.Ctx)
}

type CountryControllerImpl struct {
	CountriesAPI services.CountriesAPI
}

func (ctl *CountryControllerImpl) HandleSearch(c *fiber.Ctx) error {

	country, err := ctl.CountriesAPI.GetCountry(c.Context(), c.Params("id"))

	if err != nil {
		c.Status(400)
		c.JSON(err)

		return nil
	}

	c.Status(200)
	c.JSON(country)

	return nil

}

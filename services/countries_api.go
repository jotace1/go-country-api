package services

import (
	"context"
	"encoding/json"
	"world-api/entities"
	"world-api/utils"
)

type CountriesAPI interface {
	GetCountry(context context.Context, countryName string) (entities.Country, error)
}

type CountriesAPIImpl struct {
	MakeRequest utils.MakeAPIRequest
}

func (c *CountriesAPIImpl) GetCountry(context context.Context, countryName string) (entities.Country, error) {
	var countries []entities.Country

	bufferResponse, err := c.MakeRequest(context, countryName)

	if err != nil {
		return entities.Country{}, err
	}

	err = json.Unmarshal([]byte(bufferResponse), &countries)

	if err != nil {
		return entities.Country{}, err
	}

	return countries[0], nil
}

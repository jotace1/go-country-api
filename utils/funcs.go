package utils

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
)

type MakeAPIRequest func(context.Context, string) ([]byte, error)

func CountryAPIMakeRequest(context context.Context, countryName string) ([]byte, error) {

	url := fmt.Sprintf("https://restcountries.com/v3.1/name/%s", countryName)
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	bufferResponse, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return bufferResponse, nil
}

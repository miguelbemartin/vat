// Package vat provides an API Client for working with the Open Service by European Commission.
//
// See https://ec.europa.eu/info/index_en for full documentation of the API which this
// package makes requests to.
package vat

import (
	"github.com/miguelbemartin/vat/models"
	"github.com/miguelbemartin/vat/services"
)

// Client holds a connection to the service
type Client struct {
	// Services used for communicating with the external service.
	validator *services.ValidatorService
	rates     *services.RatesService
}

// NewClient will create http client to create http request
func NewClient() *Client {
	// Create a new instance
	c := &Client{}

	// Init services
	c.validator = services.NewValidatorService()
	c.rates = services.NewRatesService()

	return c
}

// Validate will validate the vat number
func (c Client) Validate(vat string) (bool, error) {
	return c.validator.Validate(vat)
}

// GetRate will return the rate by the country given
func (c Client) GetRate(countryCode string) (*models.Rate, error) {
	return c.rates.Get(countryCode)
}

package vat

import (
	"github.com/miguelbemartin/vat/models"
	"github.com/miguelbemartin/vat/services"
)

// client holds a connection to the service
type client struct {
	// Services used for communicating with the external service.
	validator *services.ValidatorService
	rates     *services.RatesService
}

// NewClient will create http client to create http request
func NewClient() *client {
	// Create a new instance
	c := &client{}

	// Init services
	c.validator = services.NewValidatorService()
	c.rates = services.NewRatesService()

	return c
}

// Validate will validate the vat number
func (c client) Validate(vat string) (bool, error) {
	return c.validator.Validate(vat)
}

// GetRate will return the rate by the country given
func (c client) GetRate(countryCode string) (*models.Rate, error) {
	return c.rates.Get(countryCode)
}

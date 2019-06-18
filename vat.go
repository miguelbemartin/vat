package vat

import (
	"github.com/miguelbemartin/vat/models"
	"github.com/miguelbemartin/vat/services"
)

// Client holds a connection to the service
type client struct {
	// Services used for communicating with the external service.
	validator *services.ValidatorService
	rates     *services.RatesService
}

// newClient will create http client to create http request
func newClient() *client {
	// Create a new instance
	c := &client{}

	// Init services
	c.validator = services.NewValidatorService()
	c.rates = services.NewRatesService()

	return c
}

// Validate will validate the vat number
func Validate(vat string) (bool, error) {
	client := newClient()
	return client.validator.Validate(vat)
}

// GetRate will validate the vat number
func GetRate(countryCode string) (*models.Rate, error) {
	client := newClient()
	return client.rates.Get(countryCode)
}

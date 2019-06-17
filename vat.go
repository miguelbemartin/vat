package vat

// Client holds a connection to the service
type client struct {
	// Services used for communicating with the external service.
	validator *validatorService
	rates     *ratesService
}

// newClient will create http client to create http request
func newClient() *client {


	c := &client{}

	// Init services
	c.validator = newValidatorService(c)
	c.rates = newRatesService(c)

	return c
}

// Validate will validate the vat number
func Validate(vat string) (bool, error) {
	client := newClient()
	return client.validator.Validate(vat)
}

// GetRate will validate the vat number
func GetRate(countryCode string) (*Rate, error) {
	client := newClient()
	return client.rates.Get(countryCode)
}

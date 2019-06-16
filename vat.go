package vat

import (
	"net/http"
	"time"

	cache "github.com/patrickmn/go-cache"
)

const (
	serviceURL = ""
)

// Client holds a connection to the service
type Client struct {
	// client is the HTTP client the package will use for requests.
	client *http.Client

	// serviceURL is the url to fetch the info from the service.
	serviceURL string

	// expiry is the time of duration of the cache
	expiry time.Duration
}

// NewClient will create http client to create http request
func NewClient() *Client {
	httpClient := &http.Client{
		Timeout: time.Second * 10,
		Transport: &http.Transport{
			MaxIdleConns:       10,
			IdleConnTimeout:    30 * time.Second,
			DisableCompression: true,
		},
	}

	expiry := 10 * time.Minute

	c := &Client{
		client:     httpClient,
		serviceURL: serviceURL,
		expiry:     expiry,
	}

	// Init a new store.
	store := cache.New(expiry, 10*time.Minute)

	return c
}

// NewRequest create a new request to the external service
func (c *Client) NewRequest() (*http.Request, error) {

	// Create the request.
	req, err := http.NewRequest("GET", serviceURL, nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// Do sends a service request and returns the response
func (c *Client) Do(req *http.Request) (*http.Response, error) {
	response, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	return response, err
}

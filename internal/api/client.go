package api

import (
	"encoding/json"
	"errors"
	"net/http"
)

const (
	// DefaultEndpoint is the default API endpoint for RUVDS.
	DefaultEndpoint = "https://api.ruvds.com/v2"
)

// Currencies.
const (
	CurrencyRUB = 1
	CurrencyUSD = 3
	CurrencyEUR = 4
)

var (
	// Currencies is a map of currency codes to their string representations.
	Currencies = map[int]string{
		CurrencyRUB: "RUB",
		CurrencyUSD: "USD",
		CurrencyEUR: "EUR",
	}
)

// Counties ISO codes.
const (
	CountryRussia       = "RU"
	CountryKazakhstan   = "KZ"
	CountrySwitzerland  = "CH"
	CountryNetherlands  = "NL"
	CountryGermany      = "DE"
	CountryGreatBritain = "GB"
	CountryTurkiye      = "TR"
)

var (
	// Countries is a map of country ISO codes to their names.
	Countries = map[string]string{
		CountryRussia:       "Russia",
		CountryKazakhstan:   "Kazakhstan",
		CountrySwitzerland:  "Switzerland",
		CountryNetherlands:  "Netherlands",
		CountryGermany:      "Germany",
		CountryGreatBritain: "Great Britain",
		CountryTurkiye:      "Turkiye",
	}

	RuCountryToCode = map[string]string{
		"Россия":         CountryRussia,
		"Казахстан":      CountryKazakhstan,
		"Швейцария":      CountrySwitzerland,
		"Нидерланды":     CountryNetherlands,
		"Германия":       CountryGermany,
		"Великобритания": CountryGreatBritain,
		"Турция":         CountryTurkiye,
	}
)

// API client for interacting with the RUVDS API v2.
type Client struct {
	token      string
	endpoint   string
	httpClient *http.Client
}

// NewClient creates a new API client with the provided token and endpoint.
// The token is used for authentication, and the endpoint is the base URL for the API.
// The httpClient is used for making requests to the API.
func NewClient(token, endpoint string) *Client {
	if endpoint == "" {
		endpoint = DefaultEndpoint
	}
	return &Client{
		token:      token,
		endpoint:   endpoint,
		httpClient: &http.Client{},
	}
}

// doRequest performs an HTTP request to the RUVDS API.
// It takes the HTTP method, path, and an optional body as parameters.
// It returns the HTTP response or an error if the request fails.
// The method should be one of "GET", "POST", "PUT", "DELETE", etc.
// The path is the API endpoint to which the request is made staring from "/".
// The body is optional and can be used for POST or PUT requests.
// The response is expected to be in JSON format.
// If the response status code is not in the 2xx range, an error is returned.
// The method returns the HTTP response or an error if the request fails.
func (c *Client) doRequest(method, path string, _ /*body*/ any) (*http.Response, error) {
	req, err := http.NewRequest(method, c.endpoint+path, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+c.token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, errors.New("API request failed with status: " + resp.Status)
	}

	return resp, nil
}

// doGet performs a GET request to the RUVDS API.
// It takes the path as a parameter and returns the HTTP response or an error.
// The path is the API endpoint to which the request is made starting from "/".
// The response is expected to be in JSON format.
// If the response status code is not in the 2xx range, an error is returned.
// The method returns the HTTP response or an error if the request fails.
func (c *Client) doGet(path string, params ...string) (*http.Response, error) {
	if len(params) > 0 {
		path += "?"
		for i, param := range params {
			if i > 0 {
				path += "&"
			}
			path += param
		}
	}
	return c.doRequest("GET", path, nil)
}

/*

// doPost performs a POST request to the RUVDS API.
// It takes the path and an optional body as parameters and returns the HTTP response or an error.
// The path is the API endpoint to which the request is made starting from "/".
// The body is optional and can be used for POST requests.
func (c *Client) doPost(path string, body any) (*http.Response, error) {
	return c.doRequest("POST", path, body)
}

// doPut performs a PUT request to the RUVDS API.
// It takes the path and an optional body as parameters and returns the HTTP response or an error.
// The path is the API endpoint to which the request is made starting from "/".
// The body is optional and can be used for PUT requests.
func (c *Client) doPut(path string, body any) (*http.Response, error) {
	return c.doRequest("PUT", path, body)
}

// doDelete performs a DELETE request to the RUVDS API.
// It takes the path as a parameter and returns the HTTP response or an error.
// The path is the API endpoint to which the request is made starting from "/".
// The response is expected to be in JSON format.
// If the response status code is not in the 2xx range, an error is returned.
// The method returns the HTTP response or an error if the request fails.
func (c *Client) doDelete(path string) (*http.Response, error) {
	return c.doRequest("DELETE", path, nil)
}

*/

// getEntity retrieves an entity of type T from the RUVDS API.
// It takes the path and optional parameters as arguments.
// The path is the API endpoint to which the request is made starting from "/".
func getEntity[T any](c *Client, path string, params ...string) (*T, error) {
	resp, err := c.doGet(path, params...)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result T
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

/*

// createEntity creates a new entity of type T in the RUVDS API.
// It takes the path and the body of type T as arguments.
// The path is the API endpoint to which the request is made starting from "/".
func createEntity[T any](c *Client, path string, body T) (*T, error) {
	resp, err := c.doPost(path, body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result T
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

// updateEntity updates an existing entity of type T in the RUVDS API.
// It takes the path and the body of type T as arguments.
// The path is the API endpoint to which the request is made starting from "/".
func updateEntity[T any](c *Client, path string, body T) (*T, error) {
	resp, err := c.doPut(path, body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result T
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

// deleteEntity deletes an entity from the RUVDS API.
// It takes the path as an argument.
// The path is the API endpoint to which the request is made starting from "/".
func deleteEntity(c *Client, path string) error {
	resp, err := c.doDelete(path)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		return errors.New("failed to delete entity, status: " + resp.Status)
	}

	return nil
}

*/

// GetDataCenters retrieves a list of data centers from the RUVDS API.
func (c *Client) GetDataCenters() (*DataCentersResponse, error) {
	resp, err := getEntity[DataCentersResponse](c, "/datacenters")
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetOS retrieves a list of operating systems from the RUVDS API.
func (c *Client) GetOS() (*OSResponse, error) {
	resp, err := getEntity[OSResponse](c, "/os")
	if err != nil {
		return nil, err
	}
	return resp, nil
}

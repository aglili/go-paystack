package paystack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/aglili/gopaystack/config"
)

// CreateCustomer creates a new customer
// It returns the customer details if successful
// It returns an error if the request fails
// It takes a CreateCustomerRequest object as an argument
// It returns a CustomerResponse object and an error
// CreateCustomer creates a new customer in the Paystack system.
// It takes a CreateCustomerRequest object as input and returns a CustomerResponse object or an error.
//
// The function performs the following steps:
// 1. Constructs the URL for the Paystack customer creation endpoint.
// 2. Marshals the request object into JSON format.
// 3. Creates a new HTTP POST request with the JSON payload.
// 4. Sets the necessary headers, including the authorization token and content type.
// 5. Sends the HTTP request using the client's HTTP client.
// 6. Reads and processes the response body.
// 7. Checks for any API errors based on the response status code.
// 8. Unmarshals the response body into a CustomerResponse object.
//
// Parameters:
// - req: A pointer to a CreateCustomerRequest object containing the customer details.
//
// Returns:
// - A pointer to a CustomerResponse object containing the created customer details.
// - An error if any step in the process fails.
func (c *Client) CreateCustomer(req *CreateCustomerRequest) (*CustomerResponse, error) {
	url := config.BaseURL + "/customer"

	payload, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request: %v", err)
	}

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	request.Header.Set("Authorization", "Bearer "+c.secretKey)
	request.Header.Set("Content-Type", "application/json")

	response, err := c.httpClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API error: %s", body)
	}

	var customerResponse CustomerResponse
	err = json.Unmarshal(body, &customerResponse)
	if err != nil {
		return nil, fmt.Errorf("error parsing response: %v", err)
	}

	return &customerResponse, nil
}



// ListCustomers retrieves a list of customers from the Paystack API.
// It sends a GET request to the /customer endpoint with the provided request payload.
//
// Parameters:
//   - req: A pointer to a ListCustomersRequest struct containing the request parameters.
//
// Returns:
//   - A pointer to a ListCustomersResponse struct containing the response data.
//   - An error if the request fails or the response cannot be parsed.
//
// Possible errors:
//   - If the request payload cannot be marshalled to JSON.
//   - If the HTTP request cannot be created.
//   - If the HTTP request fails.
//   - If the response body cannot be read.
//   - If the API returns a non-200 status code.
//   - If the response body cannot be unmarshalled to ListCustomersResponse.
func (c *Client) ListCustomers(req *ListCustomersRequest ) (*ListCustomersResponse, error) {
	url := config.BaseURL + "/customer"

	payload, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request: %v", err)
	}

	request, err := http.NewRequest("GET", url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	request.Header.Set("Authorization", "Bearer "+c.secretKey)
	request.Header.Set("Content-Type", "application/json")

	response, err := c.httpClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API error: %s", body)
	}

	var listCustomersResponse ListCustomersResponse
	err = json.Unmarshal(body, &listCustomersResponse)
	if err != nil {
		return nil, fmt.Errorf("error parsing response: %v", err)
	}

	return &listCustomersResponse, nil
}



// GetCustomer retrieves a customer by email or customer code from the Paystack API.
// It sends a GET request to the /customer/:email_or_code endpoint.
//
// Parameters:
//   - customerCodeOrEmail: A string containing the customer code or email to retrieve.
//
// Returns:
//   - A pointer to a CustomerResponse struct containing the customer details.
//   - An error if the request fails or the response cannot be parsed.
//
// Possible errors:
//   - If the HTTP request cannot be created.
//   - If the HTTP request fails.
//   - If the response body cannot be read.
//   - If the API returns a non-200 status code.
//   - If the response body cannot be unmarshalled to a CustomerResponse struct.
func (c *Client) GetCustomer(customerCodeOrEmail string) (*GetCustomerResponse, error) {
	url := config.BaseURL + "/customer/" + customerCodeOrEmail

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	request.Header.Set("Authorization", "Bearer "+c.secretKey)
	request.Header.Set("Content-Type", "application/json")

	response, err := c.httpClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API error: %s", body)
	}

	var customerResponse GetCustomerResponse
	err = json.Unmarshal(body, &customerResponse)
	if err != nil {
		return nil, fmt.Errorf("error parsing response: %v", err)
	}

	return &customerResponse, nil
}


func (c *Client) UpdateCustomer(customerCode string, req *UpdateCustomerRequest) (*CustomerResponse, error) {
	url := config.BaseURL + "/customer" + customerCode

	payload, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request: %v", err)
	}

	request, err := http.NewRequest("PUT", url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	request.Header.Set("Authorization", "Bearer"+c.secretKey)
	request.Header.Set("Content-Type", "application/json")

	response, err := c.httpClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API error: %s", body)
	}

	var customerResponse CustomerResponse
	err = json.Unmarshal(body, &customerResponse)
	if err != nil {
		return nil, fmt.Errorf("error parsing response: %v", err)
	}

	return &customerResponse, nil
}
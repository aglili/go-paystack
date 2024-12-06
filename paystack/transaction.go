package paystack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/aglili/gopaystack/config"
)

// InitializeTransaction initializes a new transaction with the provided request data.
// It sends a POST request to the Paystack API to create the transaction.
//
// Parameters:
//   - req: A pointer to an InitializeTransactionRequest struct containing the transaction details.
//
// Returns:
//   - A pointer to a TransactionResponse struct containing the response from the Paystack API.
//   - An error if the request fails or the response cannot be parsed.
//
// Possible errors:
//   - If the request payload cannot be marshalled to JSON.
//   - If the HTTP request cannot be created.
//   - If the HTTP request fails.
//   - If the response body cannot be read.
//   - If the API returns a non-200 status code.
//   - If the response body cannot be unmarshalled to a TransactionResponse struct.
func (c *Client) InitializeTransaction(req *InitializeTransactionRequest) (*TransactionResponse, error) {
	url := config.BaseURL + "/transaction/initialize"

	payload, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request: %v", err)
	}

	//create a new request
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	//set the request headers
	request.Header.Set("Authorization", "Bearer "+c.secretKey)
	request.Header.Set("Content-Type", "application/json")

	// make the request
	response, err := c.httpClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer response.Body.Close()

	//decode the response
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API error: %s", body)
	}

	// parse the response
	var transactionResponse TransactionResponse
	err = json.Unmarshal(body, &transactionResponse)
	if err != nil {
		return nil, fmt.Errorf("error parsing response: %v", err)
	}

	return &transactionResponse, nil
}

// VerifyTransaction verifies a transaction on Paystack using the provided reference.
// It sends a GET request to the Paystack API and returns the transaction details.
//
// Parameters:
//   - reference: The reference string of the transaction to be verified.
//
// Returns:
//   - *schema.VerifyTransactionResponse: The response containing the transaction details.
//   - error: An error object if an error occurred during the request or response parsing.
func (c *Client) VerifyTransaction(reference string) (*VerifyTransactionResponse, error) {
	url := fmt.Sprintf("%s/transaction/verify/%s", config.BaseURL, reference)

	//create a new request
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	//set the request headers
	request.Header.Set("Authorization", "Bearer "+c.secretKey)
	request.Header.Set("Content-Type", "application/json")

	// make the request
	response, err := c.httpClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer response.Body.Close()

	//decode the response
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API error: %s", body)
	}

	// parse the response
	var transactionResponse VerifyTransactionResponse
	err = json.Unmarshal(body, &transactionResponse)
	if err != nil {
		return nil, fmt.Errorf("error parsing response: %v", err)
	}

	return &transactionResponse, nil
}

// ListTransactions retrieves a list of transactions from the Paystack API.
// It sends a GET request to the /transaction endpoint with the provided request payload.
//
// Parameters:
//   - req: A pointer to a ListTransactionsRequest struct containing the request parameters.
//
// Returns:
//   - A pointer to a ListTransactionsResponse struct containing the response data.
//   - An error if the request fails or the response cannot be decoded.
func (c *Client) ListTransactions(req *ListTransactionsRequest) (*ListTransactionsResponse, error) {
	url := config.BaseURL + "/transaction"

	payload, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request: %v", err)
	}

	//create a new request
	request, err := http.NewRequest("GET", url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	//set the request headers
	request.Header.Set("Authorization", "Bearer "+c.secretKey)
	request.Header.Set("Content-Type", "application/json")

	// make the request
	response, err := c.httpClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}

	defer response.Body.Close()

	//decode the response

	var listTransactionsResponse ListTransactionsResponse
	err = json.NewDecoder(response.Body).Decode(&listTransactionsResponse)
	if err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	return &listTransactionsResponse, nil
}

// FetchTransaction retrieves the details of a transaction from Paystack using the provided reference.
// It sends a GET request to the Paystack API and returns the transaction details.
//
// Parameters:
//   - reference: A string representing the transaction reference.
//
// Returns:
//   - *schema.VerifyTransactionResponse: A pointer to the response structure containing transaction details.
//   - error: An error object if an error occurred during the request or response parsing.
//
// Example:
//
//	transaction, err := client.FetchTransaction("transaction_reference")
//	if err != nil {
//	    log.Fatalf("Error fetching transaction: %v", err)
//	}
//	fmt.Printf("Transaction details: %+v\n", transaction)
func (c *Client) FetchTransaction(reference string) (*VerifyTransactionResponse, error) {
	url := fmt.Sprintf("%s/transaction/%s", config.BaseURL, reference)

	//create a new request
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	//set the request headers
	request.Header.Set("Authorization", "Bearer "+c.secretKey)
	request.Header.Set("Content-Type", "application/json")

	// make the request
	response, err := c.httpClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer response.Body.Close()

	//decode the response
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API error: %s", body)
	}

	// parse the response
	var transactionResponse VerifyTransactionResponse
	err = json.Unmarshal(body, &transactionResponse)
	if err != nil {
		return nil, fmt.Errorf("error parsing response: %v", err)
	}

	return &transactionResponse, nil
}

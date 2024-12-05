package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aglili/gopaystack/config"
	"github.com/aglili/gopaystack/paystack"

	"github.com/stretchr/testify/assert"
)

func TestInitialisePayment(t *testing.T) {
	// mock the response
	Response := paystack.TransactionResponse{
		Status:  true,
		Message: "Transaction initialized",
		Data: struct {
			AuthorizationURL string `json:"authorization_url"`
			AccessCode       string `json:"access_code"`
			Reference        string `json:"reference"`
		}{
			AuthorizationURL: "https://checkout.paystack.com/9k2f3k4",
			AccessCode:       "9k2f3k4",
			Reference:        "9k2f3k4",
		},
	}

	// create a mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "POST")
		assert.Equal(t, r.URL.Path, "/transaction/initialize")

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Response)
	}))

	defer server.Close()

	config.BaseURL = server.URL

	client := paystack.NewClient("sk_test_1234567890")

	req := &paystack.InitializeTransactionRequest{
		Reference:   "9k2f3k4",
		Amount:      10000,
		Email:       "test@test.com",
		CallbackURL: "https://example.com/callback",
	}

	res, err := client.InitializeTransaction(req)
	assert.Nil(t, err)
	assert.NoError(t, err)
	assert.Equal(t, res.Status, true)
	assert.Equal(t, res.Message, "Transaction initialized")
	assert.Equal(t, res.Data.AccessCode, "9k2f3k4")
	assert.Equal(t, res.Data.Reference, "9k2f3k4")
	assert.Equal(t, res.Data.AuthorizationURL, "https://checkout.paystack.com/9k2f3k4")

}


func TestVerifyTransaction(t *testing.T){
	// mock the response
	Response := paystack.VerifyTransactionResponse{
		Status:  true,
		Message: "Transaction fetched",
		Data: struct {
			Amount           int    `json:"amount"`
			TransactionDate  string `json:"transaction_date"`
			TransactionStatus string `json:"status"`
			Reference        string `json:"reference"`
		}{
			Amount:           10000,
			TransactionDate:  "2020-12-12T12:12:12",
			TransactionStatus: "success",
			Reference:        "9k2f3k4",
		},
	}

	// create a mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "GET")
		assert.Equal(t, r.URL.Path, "/transaction/verify/9k2f3k4")

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Response)
	}))

	defer server.Close()

	config.BaseURL = server.URL

	client := paystack.NewClient("sk_test_1234567890")

	res, err := client.VerifyTransaction("9k2f3k4")
	assert.Nil(t, err)
	assert.NoError(t, err)
	assert.Equal(t, res.Status, true)
	assert.Equal(t, res.Message, "Transaction fetched")
	assert.Equal(t, res.Data.Amount, 10000)
	assert.Equal(t, res.Data.Reference, "9k2f3k4")
	assert.Equal(t, res.Data.TransactionStatus, "success")
	assert.Equal(t, res.Data.TransactionDate, "2020-12-12T12:12:12")
}



func TestListTransactions(t *testing.T)  {
	// mock the response
	Response := paystack.ListTransactionsResponse{
		Status:  true,
		Message: "Transactions fetched",
		Data: []struct {
			ID              int    `json:"id"`
			TransactionDate string `json:"transaction_date"`
			Amount          int    `json:"amount"`
			Currency        string `json:"currency"`
			Channel         string `json:"channel"`
			Reference       string `json:"reference"`
			Status          string `json:"status"`
		}{
			{
				ID:              1,
				TransactionDate: "2020-12-12T12:12:12",
				Amount:          10000,
				Currency:        "NGN",
				Channel:         "card",
				Reference:       "9k2f3k4",
				Status:          "success",
			},
		},
	}

	// create a mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "GET")
		assert.Equal(t, r.URL.Path, "/transaction")

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Response)
	}))

	defer server.Close()

	config.BaseURL = server.URL

	client := paystack.NewClient("sk_test_1234567890")

	req := &paystack.ListTransactionsRequest{
		PerPage: 10,
		Page: 1,
	}

	res, err := client.ListTransactions(req)
	assert.Nil(t, err)
	assert.NoError(t, err)
	assert.Equal(t, res.Status, true)
	assert.Equal(t, res.Message, "Transactions fetched")
	assert.Equal(t, len(res.Data), 1)
	assert.Equal(t, res.Data[0].ID, 1)
	assert.Equal(t, res.Data[0].TransactionDate, "2020-12-12T12:12:12")
	assert.Equal(t, res.Data[0].Amount, 10000)
	assert.Equal(t, res.Data[0].Currency, "NGN")
	assert.Equal(t, res.Data[0].Channel, "card")	
}



func TestFetchTransaction(t *testing.T)  {

	// mock the response
	Response := paystack.VerifyTransactionResponse{
		Status:  true,
		Message: "Transaction fetched",
		Data: struct {
			Amount           int    `json:"amount"`
			TransactionDate  string `json:"transaction_date"`
			TransactionStatus string `json:"status"`
			Reference        string `json:"reference"`
		}{
			Amount:           10000,
			TransactionDate:  "2020-12-12T12:12:12",
			TransactionStatus: "success",
			Reference:        "9k2f3k4",
		},
	}

	// create a mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "GET")
		assert.Equal(t, r.URL.Path, "/transaction/9k2f3k4")

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Response)
	}))

	defer server.Close()

	config.BaseURL = server.URL

	client := paystack.NewClient("sk_test_1234567890")

	res, err := client.FetchTransaction("9k2f3k4")
	assert.Nil(t, err)
	assert.NoError(t, err)
	assert.Equal(t, res.Status, true)
	assert.Equal(t, res.Message, "Transaction fetched")
	assert.Equal(t, res.Data.Amount, 10000)
	assert.Equal(t, res.Data.Reference, "9k2f3k4")
	assert.Equal(t, res.Data.TransactionStatus, "success")
	assert.Equal(t, res.Data.TransactionDate, "2020-12-12T12:12:12")
	assert.NotEqual(t, res.Data.TransactionDate, "2020-12-12T12:12:13")
	
}

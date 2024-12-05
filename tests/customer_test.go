package tests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aglili/gopaystack/config"
	"github.com/aglili/gopaystack/paystack"
	"github.com/stretchr/testify/assert"
)

func TestCreateCustomer(t *testing.T) {
	// mock the response
	Response := paystack.CustomerResponse{
		Status:  true,
		Message: "Customer created",
		Data: struct {
			ID           int                    `json:"id"`
			FirstName    string                 `json:"first_name"`
			LastName     string                 `json:"last_name"`
			Email        string                 `json:"email"`
			Phone        string                 `json:"phone"`
			CustomerCode string                 `json:"customer_code"`
			Metadata     map[string]interface{} `json:"metadata"`
		}{
			ID:           1,
			CustomerCode: "CUS_1234567890",
			FirstName:    "John",
			LastName:     "Doe",
			Email:        "test@test.com",
			Phone:        "1234567890",
		},
	}

	// create a mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "POST")
		assert.Equal(t, r.URL.Path, "/customer")

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Response)
	},
	))

	defer server.Close()

	config.BaseURL = server.URL

	client := paystack.NewClient("sk_test_1234567890")

	req := &paystack.CreateCustomerRequest{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "test@test.com",
		Phone:     "1234567890",
	}

	res, err := client.CreateCustomer(req)
	assert.Nil(t, err)
	assert.NoError(t, err)
	assert.Equal(t, res.Status, true)
	assert.Equal(t, res.Message, "Customer created")
	assert.Equal(t, res.Data.ID, 1)
	assert.Equal(t, res.Data.CustomerCode, "CUS_1234567890")
	assert.Equal(t, res.Data.FirstName, "John")
	assert.Equal(t, res.Data.LastName, "Doe")
}

func TestListCustomers(t *testing.T) {
	// mock the response
	Response := paystack.ListCustomersResponse{
		Status:  true,
		Message: "Customers fetched",
		Data: []struct {
			ID           int                    `json:"id"`
			FirstName    string                 `json:"first_name"`
			LastName     string                 `json:"last_name"`
			Email        string                 `json:"email"`
			Phone        string                 `json:"phone"`
			CustomerCode string                 `json:"customer_code"`
			Metadata     map[string]interface{} `json:"metadata"`
		}{
			{
				ID:           1,
				CustomerCode: "CUS_1234567890",
				FirstName:    "John",
				LastName:     "Doe",
				Email:        "test@test.com",
				Phone:        "1234567890",
			},
			{
				ID:           2,
				CustomerCode: "CUS_0987654321",
				FirstName:    "Jane",
				LastName:     "Doe",
				Email:        "jane@test.com",
				Phone:        "0987654321",
			},
		},
	}

	// create a mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "GET")
		assert.Equal(t, r.URL.Path, "/customer")

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Response)
	}))

	defer server.Close()

	config.BaseURL = server.URL

	client := paystack.NewClient("sk_test_1234567890")

	req := &paystack.ListCustomersRequest{
		PerPage: 10,
	}

	res, err := client.ListCustomers(req)
	assert.Nil(t, err)
	assert.NoError(t, err)
	assert.Equal(t, res.Status, true)
	assert.Equal(t, res.Message, "Customers fetched")
	assert.Equal(t, len(res.Data), 2)
	assert.Equal(t, res.Data[0].ID, 1)
	assert.Equal(t, res.Data[0].CustomerCode, "CUS_1234567890")
	assert.Equal(t, res.Data[0].FirstName, "John")
	assert.Equal(t, res.Data[0].LastName, "Doe")
	assert.Equal(t, res.Data[1].ID, 2)
}

func TestGetCustomer(t *testing.T) {
	customerCodeOrEmail := "CUS_1234567890"
	
	Response := paystack.GetCustomerResponse{
		Status:  true,
		Message: "Customer fetched",
		Data: struct {
			ID            int    `json:"id"`
			FirstName     string `json:"first_name"`
			LastName      string `json:"last_name"`
			Email         string `json:"email"`
			Phone         string `json:"phone"`
			CustomerCode  string `json:"customer_code"`
			Subscriptions []struct {
				ID               int    `json:"id"`
				Plan             string `json:"plan"`
				SubscriptionCode string `json:"subscription_code"`
			} `json:"subscriptions"`
			Transaction []struct {
				ID        int    `json:"id"`
				Amount    int    `json:"amount"`
				Reference string `json:"reference"`
			} `json:"transactions"`
		}{
			ID:           1,
			FirstName:    "John",
			LastName:     "Doe",
			Email:        "test@test.com",
			Phone:        "1234567890",
			CustomerCode: "CUS_1234567890",
			Subscriptions: []struct {
				ID               int    `json:"id"`
				Plan             string `json:"plan"`
				SubscriptionCode string `json:"subscription_code"`
			}{
				{
					ID:               1,
					Plan:             "Premium",
					SubscriptionCode: "SUB_0123456",
				},
				{
					ID:               2,
					Plan:             "VIP",
					SubscriptionCode: "SUB_4875158",
				},
			},
			Transaction: []struct {
				ID        int    `json:"id"`
				Amount    int    `json:"amount"`
				Reference string `json:"reference"`
			}{
				{
					ID:        1,
					Amount:    500,
					Reference: "PremiumPay",
				},
				{
					ID:        2,
					Amount:    1000,
					Reference: "VIPPay",
				},
			},
		},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "GET")
		assert.Equal(t, r.URL.Path, fmt.Sprintf("/customer/%s", customerCodeOrEmail))

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Response)
	}))

	defer server.Close()

	config.BaseURL = server.URL

	client := paystack.NewClient("sk_test_1234567890")

	res, err := client.GetCustomer(customerCodeOrEmail)
	assert.Nil(t, err)
	assert.NoError(t, err)
	assert.Equal(t, res.Status, true)
	assert.Equal(t, res.Message, "Customer fetched")
	assert.Equal(t, len(res.Data.Subscriptions), 2)
	assert.Equal(t, len(res.Data.Transaction), 2)
	assert.Equal(t, res.Data.ID, 1)
	assert.Equal(t, res.Data.CustomerCode, "CUS_1234567890")
	assert.Equal(t, res.Data.FirstName, "John")
	assert.Equal(t, res.Data.LastName, "Doe")
	assert.Equal(t, res.Data.Subscriptions[0].ID, 1)
	assert.Equal(t, res.Data.Subscriptions[1].ID, 2)
	assert.Equal(t, res.Data.Transaction[0].ID, 1)
	assert.Equal(t, res.Data.Transaction[1].ID, 2)
}


func TestUpdateCustomer(t *testing.T) {
	customerCode := "CUS_1234567890"

	Response := paystack.CustomerResponse {
		Status: true,
		Message: "Customer updated",
		Data: struct {
			ID int `json:"id"`
			FirstName string `json:"first_name"`
			LastName string `json:"last_name"`
			Email string `json:"email"`
			Phone string `json:"phone"`
			CustomerCode string `json:"customer_code"`
			Metadata map[string]interface{} `json:"metadata"`
		}{
			ID: 1,
			CustomerCode: "CUS_1234567890",
			FirstName: "Jane",
			LastName: "Smith",
			Email: "test@test.com",
			Phone: "54481255651",
		},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "PUT")
		assert.Equal(t, r.URL.Path, fmt.Sprintf("/customer/%s", customerCode))

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Response)
	}))

	defer server.Close()

	config.BaseURL = server.URL

	client := paystack.NewClient("sk_test_1234567890")

	req := &paystack.UpdateCustomerRequest{
		FirstName: "Jane",
		LastName: "Smith",
		Phone: "54481255651",
	}

	res, err := client.UpdateCustomer(customerCode, req)
	assert.Nil(t, err)
	assert.NoError(t, err)
	assert.Equal(t, res.Status, true)
	assert.Equal(t, res.Message, "Customer updated")
	assert.Equal(t, res.Data.ID, 1)
	assert.Equal(t, res.Data.CustomerCode, customerCode)
	assert.Equal(t, res.Data.FirstName, "Jane")
	assert.Equal(t, res.Data.LastName, "Smith")
	assert.Equal(t, res.Data.Phone, "54481255651")
}
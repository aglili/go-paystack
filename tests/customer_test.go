package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aglili/go-paystack/config"
	"github.com/aglili/go-paystack/pkg/paystack"
	"github.com/aglili/go-paystack/schema"
	"github.com/stretchr/testify/assert"
)





func TestCreateCustomer(t *testing.T) {
	// mock the response
	Response := schema.CustomerResponse{
		Status:  true,
		Message: "Customer created",
		Data: struct {
			ID          int                    `json:"id"`
			FirstName   string                 `json:"first_name"`
			LastName    string                 `json:"last_name"`
			Email       string                 `json:"email"`
			Phone       string                 `json:"phone"`
			CustomerCode string                `json:"customer_code"`
			Metadata    map[string]interface{} `json:"metadata"`
		}{
			ID:    1,
			CustomerCode: "CUS_1234567890",
			FirstName:   "John",
			LastName:    "Doe",
			Email:       "test@test.com",
			Phone:       "1234567890",
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


	req := &schema.CreateCustomerRequest{
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
	Response := schema.ListCustomersResponse{
		Status:  true,
		Message: "Customers fetched",
		Data: []struct {
			ID          int                    `json:"id"`
			FirstName   string                 `json:"first_name"`
			LastName    string                 `json:"last_name"`
			Email       string                 `json:"email"`
			Phone       string                 `json:"phone"`
			CustomerCode string                `json:"customer_code"`
			Metadata    map[string]interface{} `json:"metadata"`
		}{
			{
				ID:    1,
				CustomerCode: "CUS_1234567890",
				FirstName:   "John",
				LastName:    "Doe",
				Email:       "test@test.com",
				Phone:       "1234567890",
			},
			{
				ID:    2,
				CustomerCode: "CUS_0987654321",
				FirstName:   "Jane",
				LastName:    "Doe",
				Email:       "jane@test.com",
				Phone:       "0987654321",
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

	req := &schema.ListCustomersRequest{
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



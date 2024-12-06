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

func TestCreatePlan(t *testing.T) {

	// mock the response
	Response := paystack.PlanResponse{
		Status:  true,
		Message: "Plan created",
		Data: paystack.CreatePlanRequest{
			Name:         "Basic",
			Amount:       10000,
			Interval:     "monthly",
			Description:  "Basic plan",
			SendInvoices: true,
		},
	}

	// create a mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "POST")
		assert.Equal(t, r.URL.Path, "/plan")

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Response)
	}))

	defer server.Close()

	config.BaseURL = server.URL

	client := paystack.NewClient("sk_test_1234567890")

	req := &paystack.CreatePlanRequest{
		Name:         "Basic",
		Amount:       10000,
		Interval:     "monthly",
		Description:  "Basic plan",
		SendInvoices: true,
	}

	res, err := client.CreatePlan(req)
	assert.Nil(t, err)
	assert.Equal(t, res.Status, true)
	assert.Equal(t, res.Message, "Plan created")
	assert.Equal(t, res.Data.Name, "Basic")
	assert.Equal(t, res.Data.Amount, 10000)
	assert.Equal(t, res.Data.Interval, "monthly")
	assert.Equal(t, res.Data.Description, "Basic plan")
}

func TestListPlans(t *testing.T) {

	// mock the response
	Response := paystack.ListPlansResponse{
		Status:  true,
		Message: "Plans retrieved",
		Data: []struct {
			Name        string `json:"name"`
			Amount      int    `json:"amount"`
			Description string `json:"description"`
			PlanCode    string `json:"plan_code"`
		}{
			{
				Name:        "Basic",
				Amount:      10000,
				Description: "Basic plan",
				PlanCode:    "PLN_1234567890",
			},
			{
				Name:        "Pro",
				Amount:      20000,
				Description: "Pro plan",
				PlanCode:    "PLN_0987654321",
			},
		},
	}

	// create a mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "GET")
		assert.Equal(t, r.URL.Path, "/plan")

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Response)
	}))

	defer server.Close()

	config.BaseURL = server.URL

	client := paystack.NewClient("sk_test_1234567890")

	res, err := client.ListPlans()

	assert.Nil(t, err)
	assert.Equal(t, res.Status, true)
	assert.Equal(t, res.Message, "Plans retrieved")
	assert.Equal(t, len(res.Data), 2)
	assert.Equal(t, res.Data[0].Name, "Basic")
	assert.Equal(t, res.Data[0].Amount, 10000)
	assert.Equal(t, res.Data[0].Description, "Basic plan")
}

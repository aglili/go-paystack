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






func TestCreatePlan(t *testing.T){

	// mock the response
	Response := paystack.PlanResponse{
		Status:  true,
		Message: "Plan created",
		Data: paystack.CreatePlanRequest{
			Name: "Basic",
			Amount: 10000,
			Interval: "monthly",
			Description: "Basic plan",
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
		Name: "Basic",
		Amount: 10000,
		Interval: "monthly",
		Description: "Basic plan",
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
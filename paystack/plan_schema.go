package paystack

// CreatePlanRequest represents the body parameters for the CreatePlan API.
type CreatePlanRequest struct {
	Name         string `json:"name"`
	Amount       int    `json:"amount"`
	Interval     string `json:"interval"`
	Description  string `json:"description,omitempty"`
	SendInvoices bool   `json:"send_invoices,omitempty"`
	SendSMS      bool   `json:"send_sms,omitempty"`
	Currency     string `json:"currency,omitempty"`
}

// PlanResponse represents the response from the CreatePlan API.
type PlanResponse struct {
	Status  bool              `json:"status"`
	Message string            `json:"message"`
	Data    CreatePlanRequest `json:"data"`
}

// ListPlansResponse represents the response from the ListPlans API.
type ListPlansResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    []struct {
		Name        string `json:"name"`
		Amount      int    `json:"amount"`
		Description string `json:"description"`
		PlanCode    string `json:"plan_code"`
	} `json:"data"`
}

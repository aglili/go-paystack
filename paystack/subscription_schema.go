package paystack

// CreateSubscriptionRequest represents the body parameters for the CreateSubscription API.
type CreateSubscriptionRequest struct {
	Customer      string `json:"customer"`
	Plan          string `json:"plan"`
	Authorization string `json:"authorization"`
	StartDate     string `json:"start_date,omitempty"`
}

// SubscriptionResponse represents the response body for the CreateSubscription API.
type SubscriptionResponse struct {
	Status  bool                      `json:"status"`
	Message string                    `json:"message"`
	Data    CreateSubscriptionRequest `json:"data"`
}

// ListSubscriptionsResponse represents the response from the ListSubscriptions API.
type ListSubscriptionsResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    []struct {
		ID       int `json:"id"`
		Customer struct {
			FirstName string                 `json:"first_name"`
			LastName  string                 `json:"last_name"`
			Email     string                 `json:"email"`
			Phone     string                 `json:"phone,omitempty"`
			Metadata  map[string]interface{} `json:"metadata"`
		} `json:"customer"`
		Plan struct {
			Name        string `json:"name"`
			Amount      int    `json:"amount"`
			Description string `json:"description"`
			PlanCode    string `json:"plan_code"`
		} `json:"plan"`
		SubscriptionCode string `json:"subscription_code"`
		Status           string `json:"status"`
		Quantity         int    `json:"quantity"`
		Amount           int    `json:"amount"`
		EmailToken       string `json:"email_token"`
	} `json:"data"`
}

// GetSubscriptionResponse represents the response from the FetchSubscriptions API.
type GetSubscriptionResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    struct {
		ID       int `json:"id"`
		Customer struct {
			FirstName string                 `json:"first_name"`
			LastName  string                 `json:"last_name"`
			Email     string                 `json:"email"`
			Phone     string                 `json:"phone,omitempty"`
			Metadata  map[string]interface{} `json:"metadata"`
		} `json:"customer"`
		Plan struct {
			Name        string `json:"name"`
			Amount      int    `json:"amount"`
			Description string `json:"description"`
			PlanCode    string `json:"plan_code"`
		} `json:"plan"`
		SubscriptionCode string `json:"subscription_code"`
		Status           string `json:"status"`
		Quantity         int    `json:"quantity"`
		Amount           int    `json:"amount"`
		EmailToken       string `json:"email_token"`
	} `json:"data"`
}

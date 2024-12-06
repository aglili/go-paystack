package paystack

// CreateCustomerRequest represents the body parameters for the CreateCustomer API.
type CreateCustomerRequest struct {
	FirstName string                 `json:"first_name"`
	LastName  string                 `json:"last_name"`
	Email     string                 `json:"email"`
	Phone     string                 `json:"phone,omitempty"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
}

// CustomerResponse represents the response body for the CreateCustomer API.
type CustomerResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    struct {
		ID           int                    `json:"id"`
		FirstName    string                 `json:"first_name"`
		LastName     string                 `json:"last_name"`
		Email        string                 `json:"email"`
		Phone        string                 `json:"phone"`
		CustomerCode string                 `json:"customer_code"`
		Metadata     map[string]interface{} `json:"metadata"`
	} `json:"data"`
}

type ListCustomersRequest struct {
	PerPage int    `json:"perPage,omitempty"`
	Page    int    `json:"page,omitempty"`
	From    string `json:"from,omitempty"`
	To      string `json:"to,omitempty"`
}

type ListCustomersResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    []struct {
		ID           int                    `json:"id"`
		FirstName    string                 `json:"first_name"`
		LastName     string                 `json:"last_name"`
		Email        string                 `json:"email"`
		Phone        string                 `json:"phone"`
		CustomerCode string                 `json:"customer_code"`
		Metadata     map[string]interface{} `json:"metadata"`
	} `json:"data"`
}

type GetCustomerResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    struct {
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
	} `json:"data"`
}

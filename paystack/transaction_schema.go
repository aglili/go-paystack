package paystack

// InitializeTransactionRequest represents the body parameters for the InitializeTransaction API.
type InitializeTransactionRequest struct {
	Reference   string                 `json:"reference"`
	Amount      int                    `json:"amount"`
	Email       string                 `json:"email"`
	CallbackURL string                 `json:"callback_url,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

// TransactionResponse represents the response body for the InitializeTransaction API.
type TransactionResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    struct {
		AuthorizationURL string `json:"authorization_url"`
		AccessCode       string `json:"access_code"`
		Reference        string `json:"reference"`
	} `json:"data"`
}

// VerifyTransactionResponse represents the response body for the VerifyTransaction API.
type VerifyTransactionResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Amount            int    `json:"amount"`
		TransactionDate   string `json:"transaction_date"`
		TransactionStatus string `json:"status"`
		Reference         string `json:"reference"`
	} `json:"data"`
}

// ListTransactionsRequest represents the body parameters for the ListTransactions API.
type ListTransactionsRequest struct {
	PerPage int    `json:"perPage,omitempty"`
	Page    int    `json:"page,omitempty"`
	From    string `json:"from,omitempty"`
	To      string `json:"to,omitempty"`
}

// ListTransactionsResponse represents the response body for the ListTransactions API.
type ListTransactionsResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    []struct {
		ID              int    `json:"id"`
		TransactionDate string `json:"transaction_date"`
		Amount          int    `json:"amount"`
		Currency        string `json:"currency"`
		Channel         string `json:"channel"`
		Reference       string `json:"reference"`
		Status          string `json:"status"`
	} `json:"data"`
}

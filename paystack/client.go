package paystack

import "net/http"

type Client struct {
	secretKey  string
	httpClient *http.Client
}

func NewClient(secretKey string) *Client {
	return &Client{
		secretKey:  secretKey,
		httpClient: http.DefaultClient,
	}
}

package paystack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/aglili/gopaystack/config"
)

func (c Client) CreateSubscription(req *CreateSubscriptionRequest) (*SubscriptionResponse, error) {
	url := config.BaseURL + "/subscription"

	payload, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request: %v", err)
	}

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	request.Header.Set("Authorization", "Bearer"+c.secretKey)
	request.Header.Set("Content-Type", "application/json")

	response, err := c.httpClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %v", err)
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error creating plan: %s", body)
	}

	var subscriptionResponse SubscriptionResponse
	err = json.Unmarshal(body, &subscriptionResponse)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %v", err)
	}

	return &subscriptionResponse, nil
}

func (c Client) ListSubscriptions() (*ListSubscriptionsResponse, error) {
	url := config.BaseURL + "/subscription"

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	request.Header.Set("Authorization", "Bearer "+c.secretKey)

	response, err := c.httpClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %v", err)
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error listing plans: %s", body)
	}

	var listSubscriptionsResponse ListSubscriptionsResponse
	err = json.Unmarshal(body, &listSubscriptionsResponse)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %v", err)
	}

	return &listSubscriptionsResponse, nil
}

func (c Client) GetSubscription(subscriptionIdOrCode string) (*GetSubscriptionResponse, error) {
	url := config.BaseURL + "/subscription/" + subscriptionIdOrCode

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	request.Header.Set("Authorization", "Bearer "+c.secretKey)

	response, err := c.httpClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API error: %s", body)
	}

	var getSubscriptionResponse GetSubscriptionResponse
	err = json.Unmarshal(body, &getSubscriptionResponse)
	if err != nil {
		return nil, fmt.Errorf("error parsing response: %v", err)
	}

	return &getSubscriptionResponse, nil
}

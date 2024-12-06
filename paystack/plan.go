package paystack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/aglili/gopaystack/config"
)

func (c Client) CreatePlan(req *CreatePlanRequest) (*PlanResponse, error) {
	url := config.BaseURL + "/plan"

	payload, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request: %v", err)
	}

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	request.Header.Set("Authorization", "Bearer "+c.secretKey)
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

	var planResponse PlanResponse
	err = json.Unmarshal(body, &planResponse)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %v", err)
	}

	return &planResponse, nil
}

func (c Client) ListPlans() (*ListPlansResponse, error) {
	url := config.BaseURL + "/plan"

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

	var listPlansResponse ListPlansResponse
	err = json.Unmarshal(body, &listPlansResponse)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %v", err)
	}

	return &listPlansResponse, nil
}

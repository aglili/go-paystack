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


	fmt.Println(string(body))
	fmt.Println(response.StatusCode)

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



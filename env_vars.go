package eyc

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GetEnvVars - /environments - Returns list of environment_variables under the account
func (c *Client) GetEnvVars() ([]EnvVar, error) {
	fullURL := fmt.Sprintf("%s/environment_variables", c.HostURL)

	fmt.Printf("fullURL: %v\n", fullURL)
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return nil, err
	}
	fmt.Printf("req: %v\n", req)

	body, err := c.doRequest(req, nil)

	fmt.Printf("body: %v\n", body)

	env_vars := []EnvVar{}

	err = json.Unmarshal(body["environment_variables"], &env_vars)
	fmt.Printf("envVars: %v\n", env_vars)
	if err != nil {
		return nil, err
	}

	return env_vars, nil
}

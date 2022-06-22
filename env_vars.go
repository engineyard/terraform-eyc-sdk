package eyc

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GetEnvVars - /environments - Returns list of environment_variables under the account
func (c *Client) GetEnvVars() (EnvVars, error) {
	fullURL := fmt.Sprintf("%s/environment_variables", c.HostURL)

	fmt.Printf("fullURL: %v\n", fullURL)
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return {}, err
	}
	fmt.Printf("req: %v\n", req)

	body, err := c.doRequest(req, nil)

	fmt.Printf("body: %v\n", body)

	env_vars := EnvVars{}

	err = json.Unmarshal(body, &env_vars)
	fmt.Printf("envVars: %v\n", env_vars)
	if err != nil {
		return {}, err
	}

	return env_vars, nil
}

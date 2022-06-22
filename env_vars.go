package eyc

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GetEnvVars - /environments - Returns list of environment_variables under the account
func (c *Client) GetEnvVars() ([]EnvVar, error) {
	fullURL := fmt.Sprintf("%s/environment_variables", c.HostURL)
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)

	env_vars := []EnvVar{}

	err = json.Unmarshal(body, &env_vars)
	if err != nil {
		return nil, err
	}

	return env_vars, nil
}

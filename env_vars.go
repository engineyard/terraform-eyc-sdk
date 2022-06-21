package eyc

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GetEnvVars - /environments - Returns list of environment_variables under the accout
func (c *Client) GetEnvVars() ([]EnvVar, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/environment_variables", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	fmt.Printf(body)

	env_vars := []EnvVar{}
	err = json.Unmarshal(body, &env_vars)
	if err != nil {
		return nil, err
	}

	return env_vars, nil
}

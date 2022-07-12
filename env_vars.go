package eyc

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// GetEnvVars - /environments - Returns list of environment_variables under the account
func (c *Client) GetEnvVars() (map[string]interface{}, error) {
	fullURL := fmt.Sprintf("%s/environment_variables", c.HostURL)
	// env_vars := EnvVars{}
	var env_vars map[string]interface{}

	fmt.Printf("fullURL: %v\n", fullURL)
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return env_vars, err
	}
	fmt.Printf("req: %v\n", req)

	body, err := c.doRequest(req, nil)

	fmt.Printf("body from GetEnvVars: %v\n", body)

	err = json.Unmarshal(body, &env_vars)
	fmt.Printf("envVars: %v\n", env_vars)
	if err != nil {
		return env_vars, err
	}

	return env_vars, nil
}

// GetEnvVarsByEnv - GET /environments/:environment_id/environment_variables - Returns list of environment_variables under an environment
func (c *Client) GetEnvVarsByEnv(env_id int) (map[string]interface{}, error) {
	fullURL := fmt.Sprintf("%s/environments/%d/environment_variables", c.HostURL, env_id)
	// env_vars := EnvVars{}
	var env_vars map[string]interface{}

	fmt.Printf("fullURL: %v\n", fullURL)
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return env_vars, err
	}
	fmt.Printf("req: %v\n", req)

	body, err := c.doRequest(req, nil)

	fmt.Printf("body from GetEnvVarsByEnv: %v\n", body)

	err = json.Unmarshal(body, &env_vars)
	fmt.Printf("envVars: %v\n", env_vars)
	if err != nil {
		return env_vars, err
	}

	return env_vars, nil
}

// CreateEnvVar - POST /environment_variables
func (c *Client) CreateEnvVar(envVarParam EnvVarParam) (map[string]interface{}, error) {
	rb, err := json.Marshal(envVarParam)

	fullURL := fmt.Sprintf("%s/environment_variables", c.HostURL)

	req, err := http.NewRequest("POST", fullURL, strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	ev := EnvVar{}
	err = json.Unmarshal(body, &ev)
	if err != nil {
		return nil, err
	}
	return &ev, nil
}

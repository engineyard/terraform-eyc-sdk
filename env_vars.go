package eyc

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// Geterr - /environments - Returns list of environment_variables under the account
func (c *Client) GetEnvVars() (map[string]interface{}, error) {
	fullURL := fmt.Sprintf("%s/environment_variables", c.HostURL)
	// env_vars := EnvVars{}
	var env_vars map[string]interface{}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return env_vars, err
	}

	body, err := c.doRequest(req, nil)

	err = json.Unmarshal(body, &env_vars)
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

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return env_vars, err
	}

	body, err := c.doRequest(req, nil)

	err = json.Unmarshal(body, &env_vars)

	if err != nil {
		return nil, err
	}

	return env_vars, nil
}

// CRUD for each env var
// CreateEnvVar - POST /environment_variables
func (c *Client) CreateEnvVar(envVarParam EnvVarParam) (map[string]EnvVar, error) {
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

	var ev map[string]EnvVar

	err = json.Unmarshal(body, &ev)

	return ev, err
}

// GetEnvVarByID - GET /environment_variables/:id
func (c *Client) GetEnvVarByID(id int) (map[string]EnvVar, error) {
	fullURL := fmt.Sprintf("%s/environment_variables/%d", c.HostURL, id)

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)

	if err != nil {
		return nil, err
	}

	var ev map[string]EnvVar

	err = json.Unmarshal(body, &ev)

	return ev, err
}

// UpdateEnvVar - PUT /environment_variables/:id
func (c *Client) UpdateEnvVar(envVarParam EnvVarParam, id int) (map[string]EnvVar, error) {

	rb, err := json.Marshal(envVarParam)

	fullURL := fmt.Sprintf("%s/environment_variables/%v", c.HostURL, id)

	req, err := http.NewRequest("PUT", fullURL, strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)

	if err != nil {
		return nil, err
	}

	var ev map[string]EnvVar

	err = json.Unmarshal(body, &ev)

	return ev, err
}

// DeleteEnvVar - DELETE /environment_variables/:id
func (c *Client) DeleteEnvVar(id int) (map[string]EnvVar, error) {

	fullURL := fmt.Sprintf("%s/environment_variables/%d", c.HostURL, id)

	req, err := http.NewRequest("DELETE", fullURL, nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)

	if err != nil {
		return nil, err
	}

	var ev map[string]EnvVar

	err = json.Unmarshal(body, &ev)

	return ev, err
}

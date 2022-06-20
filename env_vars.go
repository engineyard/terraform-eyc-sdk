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

// GetCoffee - Returns specific coffee (no auth required)
func (c *Client) GetCoffee(coffeeID string) ([]Coffee, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/coffees/%s", c.HostURL, coffeeID), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	coffees := []Coffee{}
	err = json.Unmarshal(body, &coffees)
	if err != nil {
		return nil, err
	}

	return coffees, nil
}

// GetCoffeeIngredients - Returns list of coffee ingredients (no auth required)
func (c *Client) GetCoffeeIngredients(coffeeID string) ([]Ingredient, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/coffees/%s/ingredients", c.HostURL, coffeeID), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	ingredients := []Ingredient{}
	err = json.Unmarshal(body, &ingredients)
	if err != nil {
		return nil, err
	}

	return ingredients, nil
}

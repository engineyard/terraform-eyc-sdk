package eyc

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// HostURL - Default Core API URL
const HostURL string = "https://api.engineyard.com"

// Client -
type Client struct {
	HostURL    string
	HTTPClient *http.Client
	Token      string
}

// NewClient -
func NewClient(host, token *string) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		// Default Hashicups URL
		HostURL: HostURL,
	}

	if host != nil {
		c.HostURL = *host
	}

	// If token not provided, return empty client
	if token == nil {
		return &c, nil
	}
	c.Token = *token

	return &c, nil
}

func (c *Client) doRequest(req *http.Request, authToken *string) ([]byte, error) {
	token := c.Token

	if authToken != nil {
		token = *authToken
	}

	req.Header = http.Header{
		"Content-Type": []string{"application/json"},
		"X-EY-TOKEN":   []string{token},
		"Accept":       []string{"application/vnd.engineyard.v3+json"},
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	fmt.Printf("string(body): %v\n", string(body))
	fmt.Printf("body from client.doRequest: %v\n", body)

	// if res.StatusCode != http.StatusOK {
	// 	return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	// }

	return body, err
}

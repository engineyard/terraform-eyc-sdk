package eyc

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/user"
	"strings"
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

	// If token not provided, fetch from ~/.ey-core
	if token == nil {
		usr, err := user.Current()
		if err != nil {
			return &c, nil
		}

		eycore_path := fmt.Sprintf("%s/.ey-core", usr.HomeDir)
		eycore_data, err := os.ReadFile(eycore_path)
		if err != nil {
			return &c, nil
		}

		eycore_token := strings.Split(string(eycore_data), ": ")[1]
		token = &eycore_token

		if token == nil {
			return &c, nil
		}
	}
	c.Token = *token

	return &c, nil
}

func (c *Client) doRequest(req *http.Request, authToken *string) ([]byte, error) {
	token := c.Token
	if err != nil {
		log.Fatal(err)
	}

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

	if res.StatusCode == 404 {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}

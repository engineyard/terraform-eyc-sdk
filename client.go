package eyc

import (
	"fmt"
	"io/ioutil"
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
	fmt.Printf("inside new client")

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
		fmt.Printf("empty token")

		usr, err := user.Current()
		fmt.Printf("a")
		if err != nil {
			return &c, nil
		}
		fmt.Printf("b")

		eycore_path := fmt.Sprintf("%s/.ey-core", usr.HomeDir)
		eycore_data, err := os.ReadFile(eycore_path)
		if err != nil {
			return &c, nil
		}
		fmt.Printf("c")

		eycore_token := strings.Split(string(eycore_data), ": ")[1]
		fmt.Printf("eycore_token %s", eycore_token)
		eycore_token = strings.TrimSuffix(eycore_token, "\r\n")
		fmt.Printf("eycore_token %s", eycore_token)
		eycore_token = strings.ReplaceAll(eycore_token, " ", "")
		token = &eycore_token

		if token == nil {
			return &c, nil
		}
		fmt.Printf("d")

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

	if res.StatusCode == 404 {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}

package common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

// JikanClient handles API requests to Jikan API with rate limiting
type JikanClient struct {
	BaseURL    string
	HttpClient *http.Client
	mutex      sync.Mutex
	LastCall   time.Time
}

// NewJikanClient creates a new client for Jikan API
func NewJikanClient() *JikanClient {
	return &JikanClient{
		BaseURL:    "https://api.jikan.moe/v4",
		HttpClient: &http.Client{Timeout: 10 * time.Second},
		LastCall:   time.Now().Add(-2 * time.Second),
	}
}

// Get performs a GET request to Jikan API with rate limiting
func (c *JikanClient) Get(endpoint string) ([]byte, error) {
	c.mutex.Lock()
	// Respect rate limiting (at least 1 second between requests)
	elapsed := time.Since(c.LastCall)
	if elapsed < time.Second {
		time.Sleep(time.Second - elapsed)
	}
	c.LastCall = time.Now()
	c.mutex.Unlock()

	url := c.BaseURL + endpoint
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", "JikanAPIWrapper/1.0")

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error executing request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned error status: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	return body, nil
}

// GetAndUnmarshal makes a GET request and unmarshals the response into the provided interface
func (c *JikanClient) GetAndUnmarshal(endpoint string, target interface{}) error {
	data, err := c.Get(endpoint)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, target); err != nil {
		return fmt.Errorf("error unmarshaling response: %w", err)
	}

	return nil
}

// Global client instance
var DefaultClient = NewJikanClient()

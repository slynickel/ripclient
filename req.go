package ripclient

import (
	"encoding/json"
	"net/http"
	"time"
)

const (
	AuthHeader     = "X-Slynickel-Auth"
	DefaultTimeout = time.Second * 5
)

type IPState struct {
	CachedIP           string `json:"cachedIP"`
	CachedIPlastUpdate string `json:"cachedIPlastUpdate"`
	CurrentIP          string `json:"currentIP"`
	Updated            bool   `json:"updated"`
}

type Config struct {
	Hostname string
	Auth     string
}

func (c *Config) Post() (IPState, error) {
	var remote IPState
	client := &http.Client{
		Timeout: DefaultTimeout,
	}

	req, err := http.NewRequest("POST", "https://"+c.Hostname, nil)
	if err != nil {
		return remote, err
	}
	req.Header.Add(AuthHeader, c.Auth)

	resp, err := client.Do(req)
	if err != nil {
		return remote, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&remote)

	return remote, err
}

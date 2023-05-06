package provider

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Client -
type Client struct {
	HostURL    string
	HTTPClient *http.Client
	Token      string
	Auth       AuthStruct
}

// AuthStruct -
type AuthStruct struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// AuthResponse -
type AuthResponse struct {
	UserID   int    `json:"user_id`
	Username string `json:"username`
	Token    string `json:"token"`
}

type RealServiceRequestModel struct {
	Name                string `json:"name"`
	Address             string `json:"address"`
	Port                string `json:"port"`
	Weight              string `json:"weight,omitempty"`
	ConnectionLimit     string `json:"connectionLimit,omitempty"`
	ConnectionRateLimit string `json:"connectionRateLimit,omitempty"`
	RecoveryTime        string `json:"recoveryTime,omitempty"`
	WarmTime            string `json:"warmTime,omitempty"`
	Monitor             string `json:"monitor,omitempty"`
	MonitorList         string `json:"monitorList,omitempty"`
	LeastNumber         string `json:"leastNumber,omitempty"`
	Priority            string `json:"priority,omitempty"`
	MonitorLog          string `json:"monitorLog,omitempty"`
	SimulTunnelsLimit   string `json:"simulTunnelsLimit,omitempty"`
	CpuWeight           string `json:"cpuWeight,omitempty"`
	MemoryWeight        string `json:"memoryWeight,omitempty"`
	State               string `json:"state,omitempty"`
	VsysName            string `json:"vsysName,omitempty"`
}

func NewClient(host *string, auth *AuthStruct) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		// Default Hashicups URL
		HostURL: *host,
		Auth:    *auth,
	}

	// req, err := http.NewRequest("POST", c.HostURL, nil)
	// req.Header.Add("Content-type", "application/json")
	// req.Header.Set("Accept", "application/json")
	// req.SetBasicAuth(c.Auth.Username, c.Auth.Password)
	// if err != nil {
	// 	return nil, err
	// }
	// body, err := c.doRequest(req)
	// if err != nil {
	// 	return nil, err
	// }
	// ar := AuthResponse{}
	// err = json.Unmarshal(body, &ar)
	// if err != nil {
	// 	return nil, err
	// }

	return &c, nil
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}

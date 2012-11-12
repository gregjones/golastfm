package lastfm

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	DefaultHost       = "http://ws.audioscrobbler.com/2.0/"
	DefaultHttpClient *http.Client
)

type Client struct {
	APIKey    string
	APISecret string
	Host      string
	*http.Client
}

func NewClient(APIKey, APISecret string) *Client {
	c := &Client{APIKey: APIKey, APISecret: APISecret, Host: DefaultHost, Client: DefaultHttpClient}
	return c
}

func (c *Client) Execute(query map[string]string, v interface{}) error {
	query["api_key"] = c.APIKey
	query["format"] = "json"
	url := fmt.Sprintf("%s?%s", c.Host, Urlencode(query))
	resp, err := c.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	d := json.NewDecoder(resp.Body)
	err = d.Decode(v)
	return err
}

func (c *Client) Album() *albumMethods {
	return &albumMethods{Client: c}
}

func init() {
	DefaultHttpClient = &http.Client{}
}

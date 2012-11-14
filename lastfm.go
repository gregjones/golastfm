// Package lastfm provided methods for accessing the Last.FM API.
//
// Documentation for the Last.FM API can be found here: http://www.last.fm/api/intro
// 
// To use the API you need your own Developer Key, get one from http://www.last.fm/api/account/create
//
// Some methods required an authorized user, this is noted in the docs.
package lastfm

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
)

var (
	DefaultHost       = "http://ws.audioscrobbler.com/2.0/"
	DefaultHttpClient = http.DefaultClient
	ErrInvalidMethod  = errors.New("Invalid method")
	Errors            = map[int]error{
		3: ErrInvalidMethod,
	}
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
	url := fmt.Sprintf("%s?%s", c.Host, Urlencode(query))
	resp, err := c.Get(url)

	defer resp.Body.Close()
	if err != nil {
		return err
	}

	failureResponse := struct {
		Status string `xml:"status,attr"`
		Error  struct {
			Code    int    `xml:"code,attr"`
			Message string `xml:",chardata"`
		} `xml:"error"`
	}{}

	var b bytes.Buffer
	io.Copy(&b, resp.Body)
	err = xml.Unmarshal(b.Bytes(), &failureResponse)
	if failureResponse.Status != "ok" {
		if err, ok := Errors[failureResponse.Error.Code]; ok {
			return err
		}
	}
	err = xml.Unmarshal(b.Bytes(), v)
	return err
}

func (c *Client) Album() *AlbumMethods {
	return &AlbumMethods{Client: c}
}

// func init() {
// 	DefaultHttpClient = 
// }

package app

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"hvvApi/config"
	"net/http"
)

type Client struct {
	httpClient *http.Client
	testUrl    string
}

// func NewClient(optional... Optional) (*Client, error)
func NewClient() *Client {
	c := &Client{}
	return c
}

func NewClientWithTestUrl(testUrl string) *Client {
	c := &Client{}
	c.testUrl = testUrl
	return c
}

// Make a http POST Request to the Server
func (c *Client) post(config *config.ApiConfig, apiReq interface{}, resp interface{}, reqBody []byte, header map[string]string) error {
	host := config.Host
	//Only for testing purposes
	if c.testUrl != "" {
		host = c.testUrl
	}

	req, err := http.NewRequest(http.MethodPost, host+config.Path, bytes.NewBuffer(reqBody))
	if err != nil {
		return errors.New("failed to make a request to" + host + config.Path)
	}

	//Set Headers
	for i, v := range header {
		req.Header.Set(i, v)
	}

	if err := c.do(req, resp); err != nil {
		return err
	}

	return nil
}

func (c *Client) do(req *http.Request, resp interface{}) error {
	client := c.httpClient
	if client == nil {
		client = http.DefaultClient
	}

	//s, _ := ioutil.ReadAll(req.Body)
	//fmt.Println(string(s))

	httpResp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("client failed to do request %v", err)
	}
	defer httpResp.Body.Close()

	return json.NewDecoder(httpResp.Body).Decode(&resp)

}

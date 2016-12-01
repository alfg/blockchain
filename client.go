package blockchain

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	API_ROOT = "https://blockchain.info"
)

type Client struct {
	*http.Client
}

func (c *Client) loadResponse(path string, i interface{}, formatJson bool) error {
	full_path := API_ROOT + path
	if formatJson {
		full_path = API_ROOT + path + "?format=json"
	}

	fmt.Println("querying..." + full_path)
	rsp, e := c.Get(full_path)
	if e != nil {
		return e
	}

	defer rsp.Body.Close()

	b, e := ioutil.ReadAll(rsp.Body)
	if e != nil {
		return e
	}
	if rsp.Status[0] != '2' {
		return fmt.Errorf("expected status 2xx, got %s: %s", rsp.Status, string(b))
	}

	return json.Unmarshal(b, &i)
}

func New() (*Client, error) {
	return &Client{Client: &http.Client{}}, nil
}

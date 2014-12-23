package blockchain

import (
	"fmt"
)

type Chart struct {
	Values []*Value `json:"values"`
}

type Value struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func (c *Client) GetChart(chartType string) (*Chart, error) {
	rsp := &Chart{}
	var path = "/charts/" + chartType
	e := c.loadResponse(path, rsp, true)

	if e != nil {
		fmt.Print(e)
	}
	return rsp, e
}

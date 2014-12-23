package blockchain

import (
	"fmt"
	"testing"
)

func TestGetChart(t *testing.T) {
	fmt.Println("===== TESTING CHART =====")

	c, e := New()
	resp, e := c.GetChart("total-bitcoins")
	if e != nil {
		fmt.Print(e)
	}

	for i := range resp.Values {
		fmt.Println(resp.Values[i].X)
	}
}

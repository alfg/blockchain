package blockchain

import (
	"fmt"
	"testing"
)

func TestGetBlock(t *testing.T) {
	fmt.Println("===== TESTING BLOCK =====")

	c, e := New()
	resp, e := c.GetBlock("0000000000000bae09a7a393a8acded75aa67e46cb81f7acaa5ad94f9eacd103")
	if e != nil {
		fmt.Print(e)
	}

	fmt.Println(resp.Hash)

	for i := range resp.Tx {

		for j := range resp.Tx[i].Inputs {
			fmt.Println(resp.Tx[i].Inputs[j].Sequence)
		}
	}
}

func TestGetBlockHeight(t *testing.T) {
	fmt.Println("===== TESTING BLOCK HEIGHT =====")

	c, e := New()
	resp, e := c.GetBlockHeight("160778")
	if e != nil {
		fmt.Print(e)
	}

	for i := range resp.Blocks {
		fmt.Println(resp.Blocks[i].Hash)
		fmt.Println(resp.Blocks[i].Time)
	}
}

func TestGetLatestBlock(t *testing.T) {
	fmt.Println("===== TESTING LATEST BLOCK =====")

	c, e := New()
	resp, e := c.GetLatestBlock()
	if e != nil {
		fmt.Print(e)
	}

	fmt.Println(resp.Hash)
	fmt.Println(resp.Height)

	for _, v := range resp.TxIndexes {
		fmt.Println(v)
	}
}

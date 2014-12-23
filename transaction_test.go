package blockchain

import (
	"fmt"
	"testing"
)

func TestGetTransaction(t *testing.T) {
	fmt.Println("===== TESTING TRANSACTION =====")

	c, e := New()
	resp, e := c.GetTransaction("b6f6991d03df0e2e04dafffcd6bc418aac66049e2cd74b80f14ac86db1e3f0da")
	if e != nil {
		fmt.Print(e)
	}

	fmt.Println(resp.Hash)

	for i := range resp.Out {
		fmt.Println(resp.Out[i].Spent)
		fmt.Println(resp.Out[i].TxIndex)
	}
}

func TestGetUnconfirmedTransactions(t *testing.T) {
	fmt.Println("===== TESTING UNCONFIRMED TRANSACTION =====")

	c, e := New()
	resp, e := c.GetUnconfirmedTransactions()
	if e != nil {
		fmt.Print(e)
	}

	for i := range resp.Transactions {
		fmt.Println(resp.Transactions[i].Hash)
	}
}

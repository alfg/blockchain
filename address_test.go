package blockchain

import (
	"fmt"
	"testing"
)

func TestGetAddress(t *testing.T) {
	fmt.Println("===== TESTING ADDRESS =====")

	c, e := New()
	resp, e := c.GetAddress("162FjqU7RYdojnejCDe6zrPDUpaLcv9Hhq")
	if e != nil {
		fmt.Print(e)
	}

	fmt.Println(resp.Hash160)
	fmt.Println(resp.Address)
	fmt.Println(resp.NTx)
	fmt.Println(resp.TotalReceived)
	fmt.Println(resp.TotalSent)
	fmt.Println(resp.FinalBalance)

	for i := range resp.Txs {
		fmt.Println(resp.Txs[i].Result)

		for j := range resp.Txs[i].Inputs {
			fmt.Println(resp.Txs[i].Inputs[j].Sequence)
			fmt.Println(resp.Txs[i].Inputs[j].PrevOut.Spent)
		}
	}
}

func TestGetAddresses(t *testing.T) {
	fmt.Println("===== TESTING ADDRESSES =====")

	c, e := New()
	if e != nil {
		fmt.Print(e)
	}

	addresses := []string{
		"162FjqU7RYdojnejCDe6zrPDUpaLcv9Hhq",
		"1K3Vs8tPu2YkAoWmrkjUQVJuxr7wgPP3Wf"}

	resp, _ := c.GetAddresses(addresses)
	for i := range resp.Addresses {
		fmt.Println(resp.Addresses[i].Address)

		for j := range resp.Txs[i].Inputs {
			fmt.Println(resp.Txs[i].Inputs[j].Sequence)
			fmt.Println(resp.Txs[i].Inputs[j].PrevOut.Spent)
		}
	}
}

package blockchain

import (
	"fmt"
)

type Transaction struct {
	Hash        string    `json:"hash"`
	Ver         int       `json:"ver"`
	VinSz       int       `json:"vin_sz"`
	VoutSz      int       `json:"vout_sz"`
	LockTime    int       `json:"lock_time"`
	Size        int       `json:"size"`
	RelayedBy   string    `json:"relayed_by"`
	BlockHeight int       `json:"block_height"`
	TxIndex     int       `json:"tx_index"`
	Inputs      []*Inputs `json:"inputs"`
	Out         []*Out    `json:"out"`
}

type Transactions struct {
	Transactions []*Transaction `json:"txs"`
}

func (c *Client) GetTransaction(transaction string) (*Transaction, error) {
	rsp := &Transaction{}
	var path = "/rawtx/" + transaction
	e := c.loadResponse(path, rsp, false)

	if e != nil {
		fmt.Print(e)
	}
	return rsp, e
}

func (c *Client) GetUnconfirmedTransactions() (*Transactions, error) {
	rsp := &Transactions{}
	var path = "/unconfirmed-transactions"
	e := c.loadResponse(path, rsp, true)

	if e != nil {
		fmt.Print(e)
	}
	return rsp, e
}

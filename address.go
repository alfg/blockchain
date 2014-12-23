package blockchain

import (
	"fmt"
	"strings"
)

type Address struct {
	Hash160       string `json:"hash160"`
	Address       string `json:"address"`
	NTx           int    `json:"n_tx"`
	TotalReceived int    `json:"total_received"`
	TotalSent     int    `json:"total_sent"`
	FinalBalance  int    `json:"final_balance"`
	Txs           []*Tx  `json:"txs"`
}

type MultiAddr struct {
	Addresses []*Address `json:"addresses"`
	Txs       []*Tx      `json:"txs"`
}

type Tx struct {
	Result      int       `json:"result"`
	Ver         int       `json:"ver"`
	Size        int       `json:"size"`
	Inputs      []*Inputs `json:"inputs"`
	Time        int       `json:"time"`
	BlockHeight int       `json:"block_height"`
	TxIndex     int       `json:"tx_index"`
	VinSz       int       `json:"vin_sz"`
	Hash        string    `json:"hash"`
	VoutSz      int       `json:"vout_sz"`
	RelayedBy   string    `json:"relayed_by"`
	Out         []*Out    `json:"out"`
}

type Inputs struct {
	Sequence int      `json:"sequence"`
	Script   string   `json:"script"`
	PrevOut  *PrevOut `json:"prev_out"`
}

type PrevOut struct {
	Spent   bool   `json:"spent"`
	TxIndex int    `json:"tx_index"`
	Type    int    `json:"type"`
	Addr    string `json:"addr"`
	Value   int    `json:"value"`
	N       int    `json:"n"`
	Script  string `json:"script"`
}

type Out struct {
	Spent   bool   `json:"spent"`
	TxIndex int    `json:"tx_index"`
	Type    int    `json:"type"`
	Addr    string `json:"addr"`
	Value   int    `json:"value"`
	N       int    `json:"n"`
	Script  string `json:"script"`
}

func (c *Client) GetAddress(address string) (*Address, error) {
	rsp := &Address{}
	var path = "/address/" + address
	e := c.loadResponse(path, rsp, true)

	if e != nil {
		fmt.Print(e)
	}
	return rsp, e
}

func (c *Client) GetAddresses(addresses []string) (*MultiAddr, error) {
	rsp := &MultiAddr{}
	var path = "/multiaddr?active=" + strings.Join(addresses, "|")
	e := c.loadResponse(path, rsp, false)

	if e != nil {
		fmt.Print(e)
	}
	return rsp, e
}

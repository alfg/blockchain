package blockchain

import (
	"fmt"
)

type Block struct {
	Hash         string `json:"hash"`
	Ver          int    `json:"ver"`
	PrevBlock    string `json:"prev_block"`
	MrklRoot     string `json:"mrkl_root"`
	Time         int    `json:"time"`
	Bits         int    `json:"bits"`
	Nonce        int    `json:"nonce"`
	NTx          int    `json:"n_tx"`
	Size         int    `json:"size"`
	BlockIndex   int    `json:"block_index"`
	MainChain    bool   `json:"main_chain"`
	Height       int    `json:"height"`
	ReceivedTime int    `json:"received_time"`
	RelayedBy    string `json:"relayed_by"`
	Tx           []*Tx  `json:"tx"`
	TxIndexes    []int  `json:"txIndexes"`
}

type BlockHeight struct {
	Blocks []*Block `json:"blocks"`
}

func (c *Client) GetBlock(block string) (*Block, error) {
	rsp := &Block{}
	var path = "/rawblock/" + block
	e := c.loadResponse(path, rsp, true)

	if e != nil {
		fmt.Print(e)
	}
	return rsp, e
}

func (c *Client) GetBlockHeight(blockHeight string) (*BlockHeight, error) {
	rsp := &BlockHeight{}
	var path = "/block-height/" + blockHeight
	e := c.loadResponse(path, rsp, true)

	if e != nil {
		fmt.Print(e)
	}
	return rsp, e
}

func (c *Client) GetLatestBlock() (*Block, error) {
	rsp := &Block{}
	var path = "/latestblock"
	e := c.loadResponse(path, rsp, false)

	if e != nil {
		fmt.Print(e)
	}
	return rsp, e
}

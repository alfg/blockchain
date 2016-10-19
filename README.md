blockchain
==========

[![Build Status](https://travis-ci.org/alfg/blockchain.svg?branch=master)](https://travis-ci.org/alfg/blockchain)

Go Client for the Blockchain Data API https://blockchain.info/api/blockchain_api

Documentation: https://godoc.org/github.com/alfg/blockchain

## Installation

```
go get github.com/alfg/blockchain
```

## Example

```go
package main

import (
  "github.com/alfg/blockchain"
  "fmt"
)

func main() {

	c, e := blockchain.New()
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
```

## License

MIT

package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	conn, err := ethclient.Dial("https://mainnet.infura.io")
	if err != nil {
		log.Fatalf("Whoops, something went wrong!", err)
	}

	ctx := context.Background()
	tx, pending, err := conn.TransactionByHash(ctx, common.HexToHash("0x3278cca4db786dc32e1523c666f24679ec4f11c74bb0241a36130cdd51db5946"))
	if !pending {
		fmt.Printf("%+v\n", tx)
	}
}

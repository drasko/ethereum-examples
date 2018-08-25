package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// Create an IPC based RPC connection to a remote node and an authorized transactor
	//conn, err := rpc.NewIPCClient("/home/karalabe/.ethereum/testnet/geth.ipc")
	conn, err := ethclient.Dial("http://localhost:7545")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// Put here address of the deployed contract
	hello, err := NewHello(common.HexToAddress(<contact_address>), conn)

	for i := 0; i < 5; i++ {
		msg, err := hello.Nb(&bind.CallOpts{})
		if err != nil {
			log.Fatalf("Failed to retrieve nb var: %v", err)
		}
		fmt.Println("Hello:", msg)
	}
}

package main

import (
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var key, _ = crypto.HexToECDSA("4fa85d098eab0352dd9bfc3ca67d1ae31901ba366839650981bd87b6d0186557")

func main() {
	// Create an IPC based RPC connection to a remote node and an authorized transactor
	//conn, err := rpc.NewIPCClient("/home/karalabe/.ethereum/testnet/geth.ipc")
	conn, err := ethclient.Dial("http://localhost:7545")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	auth := bind.NewKeyedTransactor(key)

	// Deploy a new awesome contract for the binding demo
	address, tx, auction, err := DeployAuction(auth, conn, common.HexToAddress("0x92bA762c0495F9c48a7b8a9d2975ECC1426f8b32"), big.NewInt(1))
	if err != nil {
		log.Fatalf("Failed to deploy new auction contract: %v", err)
	}
	fmt.Printf("Contract pending deploy: 0x%x\n", address)
	fmt.Printf("Transaction waiting to be mined: 0x%x\n\n", tx.Hash())

	// Don't even wait, check its presence in the local pending state
	time.Sleep(250 * time.Millisecond) // Allow it to be processed by the local node :P

	ended1, err := auction.Ended(&bind.CallOpts{})
	if err != nil {
		log.Fatalf("Failed to retrieve ended var: %v", err)
	}
	fmt.Println("Ended:", ended1)

	time.Sleep(2 * time.Second)


	auction.EndAuction(&bind.TransactOpts{})

	time.Sleep(250 * time.Millisecond) // Allow it to be processed by the local node :P

	ended2, err := auction.Ended(&bind.CallOpts{})
	if err != nil {
		log.Fatalf("Failed to retrieve ended var: %v", err)
	}
	fmt.Println("Ended:", ended2)
}

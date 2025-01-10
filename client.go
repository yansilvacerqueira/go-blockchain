package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// conection with public node of cloudFlare
var client, errClient = ethclient.Dial("http://localhost:8545")

func main() {

	if errClient != nil {
		log.Fatalf("Failed to connect node: %v", errClient)
	}

	handleAddress()
	fmt.Println("connected...")
	_ = client
}

func handleAddress() {
	// wallet address, using as unique "id". That make reference to the wallet that is going receive or send transactions
	account := common.HexToAddress("0x52e2f0beea740e1b0b3470b82dad18240c92f220982d841fbce54fc1d1c90a5c")

	blockWalletBalance(account)
}

func walletBalance(account common.Address) {

	balance, err := client.BalanceAt(context.Background(), account, nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(balance)
}

func blockWalletBalance(account common.Address) {
	blockNumber := big.NewInt(5532993)

	balance, err := client.BalanceAt(context.Background(), account, blockNumber)

	if err != nil {
		log.Fatal(err)
	}

	convertWeiToEth(balance)

	fmt.Println(balance)
}

func convertWeiToEth(balance *big.Int) {
	fbalance := new(big.Float)

	// value in wei (wei = fraction of eth)
	fbalance.SetString(balance.String())

	// convert wei to eth using wei balance dividing for 10 to the power of 18
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))

	fmt.Println(ethValue)
}

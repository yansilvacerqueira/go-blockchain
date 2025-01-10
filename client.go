package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/crypto/sha3"
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

func pendingBalance(account common.Address) {
	pendingBalance, err := client.PendingBalanceAt(context.Background(), account)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(pendingBalance) // 25729324269165216042
}

func generateWallet() {
	privateKey, err := crypto.GenerateKey()

	if err != nil {
		log.Fatal(err)
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	fmt.Println(hexutil.Encode(privateKeyBytes)[2:]) // fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println(hexutil.Encode(publicKeyBytes)[4:]) // 9a7df67f79246283fdc93af76d4f8cdd62c4886e8cd870944e817dd0b97934fdd7719d0810951e03418205868a5c1b40b192451367f28e0088dd75e15de40c05

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println(address) // 0x96216849c49358B10257cb55b28eA603c874b05E

	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:])

	fmt.Println(hexutil.Encode(hash.Sum(nil)[12:]))
}

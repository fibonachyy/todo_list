package main

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	multicall "github.com/handler/contract"
)

func main() {
	client, err := ethclient.Dial("http://your-ethereum-node-url:8545")
	if err != nil {
		fmt.Println("Failed to connect to Ethereum node:", err)
		return
	}
	contractAddress := common.HexToAddress("0x1234567890abcdef...")

	multiCallContract, err := multicall.NewMulticall(contractAddress, client)
	if err != nil {
		panic("cant get contract multiCall")
	}
	privateKey, err := crypto.HexToECDSA("your-private-key-hex")
	if err != nil {
		fmt.Println("Failed to parse private key:", err)
		return
	}
	auth := bind.NewKeyedTransactor(privateKey)

	// Set other transaction options as needed (e.g., gas limit, gas price, etc.)
	auth.GasLimit = uint64(300000)
	auth.GasPrice = big.NewInt(20000000000) // 20 Gwei (in Wei)

	multiCallContract.Aggregate(auth)
}

func newMultiCall2Call(abiString string, methodName string, args ...interface{}) ([]byte, error) {
	contractABI, err := abi.JSON(strings.NewReader(abiString))
	if err != nil {
		return nil, fmt.Errorf("Error load ilks by Multicall %w", err)
	}
	// Encode the method call data
	data, err := contractABI.Pack(methodName, args...)
	if err != nil {
		return nil, fmt.Errorf("unpacking loaded ilks by Multicall %w", err)
	}

	return data, nil
}

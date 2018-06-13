package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"medchain/medchain-network/chaincode/storage"
	"medchain/medchain-network/chaincode/provider"
)

func main() {
	// Start the chaincode and make it ready for futures requests
	errStorage := shim.Start(new(storage.StorageChaincode))
	if errStorage != nil {
		fmt.Printf("Error starting Storage chaincode: %s", errStorage)
	}
	errProvider := shim.Start(new(provider.ProviderChaincode))
	if errProvider != nil {
		fmt.Printf("Error starting Provider chaincode: %s", errProvider)
	}
}

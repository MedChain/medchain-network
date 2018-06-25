package main

import (
	"fmt"
	"medchain/medchain-network/blockchain"
	"medchain/medchain-network/web"
	"medchain/medchain-network/web/controllers"
	"os"
)

func main() {
	// Definition of the Fabric SDK properties
	fSetup := blockchain.FabricSetup{
		// Channel parameters
		ChannelID:     "emr",
		ChannelConfig: os.Getenv("GOPATH") + "/src/medchain/medchain-network/fixtures/artifacts/emr.channel.tx",

		OrgAdmin:   "Admin",
		OrgName:    "Providers",
		ConfigFile: "config.yaml",

		// User parameters
		UserName: "User1",
	}
	// Chaincode parameters for Storage
	ccSetupStorage := blockchain.ChaincodeSetup{
		ChainCodeID:      "StorageChainCode",
		ChaincodeGoPath:  os.Getenv("GOPATH"),
		ChaincodePath:    "medchain/medchain-network/chaincode/storage",
		ChaincodeVersion: "1.07",
	}
	// Chaincode parameters for Provider
	ccSetupProvider := blockchain.ChaincodeSetup{
		ChainCodeID:      "ProviderChainCode",
		ChaincodeGoPath:  os.Getenv("GOPATH"),
		ChaincodePath:    "medchain/medchain-network/chaincode/provider",
		ChaincodeVersion: "2.07",
	}

	// Initialization of the Fabric SDK from the previously set properties
	err := fSetup.Initialize()
	if err != nil {
		fmt.Printf("Unable to initialize the Fabric SDK: %v\n", err)
	}
	fmt.Printf("%+v\n", fSetup)
	// Install and instantiate the chaincode
	err = fSetup.InstallAndInstantiateCC(ccSetupStorage)
	if err != nil {
		fmt.Printf("Unable to install and instantiate the chaincode: %v\n", err)
	}

	// Install and instantiate the chaincode
	err = fSetup.InstallAndInstantiateCC(ccSetupProvider)
	if err != nil {
		fmt.Printf("Unable to install and instantiate the chaincode: %v\n", err)
	}

	// Launch the web application listening
	app := &controllers.Application{
		Fabric:     &fSetup,
		CcStorage:  ccSetupStorage,
		CcProvider: ccSetupProvider,
	}
	web.Serve(app)
}

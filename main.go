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

		// Chaincode parameters
		ChainCodeID:     "medchain-network",
		ChaincodeGoPath: os.Getenv("GOPATH"),
		ChaincodePath:   "medchain/medchain-network/chaincode/",
		OrgAdmin:        "Admin",
		OrgName:         "Org1",
		ConfigFile:      "config.yaml",

		// User parameters
		UserName: "User1",
	}

	// Initialization of the Fabric SDK from the previously set properties
	err := fSetup.Initialize()
	if err != nil {
		fmt.Printf("Unable to initialize the Fabric SDK: %v\n", err)
	}

	// Install and instantiate the chaincode
	err = fSetup.InstallAndInstantiateCC()
	if err != nil {
		fmt.Printf("Unable to install and instantiate the chaincode: %v\n", err)
	}

	// Launch the web application listening
	app := &controllers.Application{
		Fabric: &fSetup,
	}
	web.Serve(app)
}

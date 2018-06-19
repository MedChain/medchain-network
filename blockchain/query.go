package blockchain

import (
	"fmt"

	"github.com/hyperledger/fabric-sdk-go/api/apitxn/chclient"
)

// QueryHello query the chaincode to get the state of hello
func (setup *FabricSetup) QueryHello(cc ChaincodeSetup, field string) (string, error) {

	// Prepare arguments
	var args []string
	args = append(args, "invoke")
	args = append(args, "query")
	args = append(args, field)

	response, err := setup.client.Query(chclient.Request{ChaincodeID: cc.ChainCodeID, Fcn: args[0], Args: [][]byte{[]byte(args[1]), []byte(args[2])}})
	if err != nil {
		return "", fmt.Errorf("failed to query: %v", err)
	}

	return string(response.Payload), nil
}

package blockchain

import (
	"fmt"
	"time"

	"github.com/hyperledger/fabric-sdk-go/api/apitxn/chclient"
)

// QueryHello query the chaincode to get the state of hello
func (setup *FabricSetup) QueryHello() (string, error) {

	// Prepare arguments
	queryRequest := chclient.Request{
		ChaincodeID: "ProviderChainCode",
		Fcn:         "queryHello",
		Args: [][]byte{
			[]byte("test"),
		},
	}

	response, err := setup.client.Query(queryRequest)
	if err != nil {
		return "", fmt.Errorf("failed to query: %v", err)
	}

	return string(response.Payload), nil
}

// QueryHello query the chaincode to get the state of hello
func (setup *FabricSetup) QueryHelloProvider() (string, error) {

	// Prepare arguments
	queryRequest := chclient.Request{
		ChaincodeID: "ProviderChainCode",
		Fcn:         "queryHelloProvider",
		Args:        nil,
	}

	response, err := setup.client.Query(queryRequest)
	if err != nil {
		return "", fmt.Errorf("failed to query: %v", err)
	}

	return string(response.Payload), nil
}

// QueryPatientData query the chaincode to get the list of EMR for patientAddr
func (setup *FabricSetup) QueryPatientData(patientAddr string) (string, error) {

	// Prepare arguments
	queryRequest := chclient.Request{
		ChaincodeID: "ProviderChainCode",
		Fcn:         "queryPatientData",
		Args: [][]byte{
			[]byte(patientAddr),
		},
	}

	response, err := setup.client.Query(queryRequest)
	if err != nil {
		return "", fmt.Errorf("failed to query: %v", err)
	}

	return string(response.Payload), nil
}

// QueryPatientRecords queries the chaincode to get the full EMR for each recordId
func (setup *FabricSetup) QueryPatientRecords(recordIds []string) (string, error) {

	// Prepare arguments
	var ids [][]byte
	// TODO: security concern to make sure all recordIds are of correct format
	for _, id := range recordIds {
		ids = append(ids, []byte(id))
	}

	queryRequest := chclient.Request{
		ChaincodeID: "ProviderChainCode",
		Fcn:         "queryPatientRecords",
		Args:        ids,
	}

	response, err := setup.client.Query(queryRequest)
	if err != nil {
		return "", fmt.Errorf("failed to query: %v", err)
	}

	return string(response.Payload), nil
}

// InvokeHello
func (setup *FabricSetup) InvokeHello(cc ChaincodeSetup, field string, value string) (string, error) {

	// Prepare arguments
	var args []string
	args = append(args, "invoke")
	args = append(args, "invoke")
	args = append(args, field)
	args = append(args, value)

	eventID := "eventInvoke"

	// Add data that will be visible in the proposal, like a description of the invoke request
	transientDataMap := make(map[string][]byte)
	transientDataMap["result"] = []byte("Transient data in hello invoke")

	// Register a notification handler on the client
	notifier := make(chan *chclient.CCEvent)
	rce, err := setup.client.RegisterChaincodeEvent(notifier, cc.ChainCodeID, eventID)
	if err != nil {
		return "", fmt.Errorf("failed to register chaincode evet: %v", err)
	}

	// Create a request (proposal) and send it
	response, err := setup.client.Execute(chclient.Request{ChaincodeID: cc.ChainCodeID, Fcn: args[0], Args: [][]byte{[]byte(args[1]), []byte(args[2]), []byte(args[3])}, TransientMap: transientDataMap})
	if err != nil {
		return "", fmt.Errorf("failed to move funds: %v", err)
	}

	// Wait for the result of the submission
	select {
	case ccEvent := <-notifier:
		fmt.Printf("Received CC event: %s\n", ccEvent)
	case <-time.After(time.Second * 20):
		return "", fmt.Errorf("did NOT receive CC event for eventId(%s)", eventID)
	}

	// Unregister the notification handler previously created on the client
	err = setup.client.UnregisterChaincodeEvent(rce)

	return response.TransactionID.ID, nil
}

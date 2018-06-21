package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// ChannelChaincode implementation of Chaincode
type ChannelChaincode struct {
}

// Init of the chaincode
// This function is called only one when the chaincode is instantiated.
// So the goal is to prepare the ledger to handle future requests.
func (t *ChannelChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("########### ChannelChaincode Init (provider) ###########")

	// Get the function and arguments from the request
	function, _ := stub.GetFunctionAndParameters()

	// Check if the request is the init function
	if function != "init" {
		return shim.Error("Unknown function call")
	}

	// Put in the ledger the key/value
	err := stub.PutState("helloprovider", []byte("provider1"))
	if err != nil {
		return shim.Error(err.Error())
	}

	// Put in the ledger the key/value
	err2 := stub.PutState("permission", []byte("true"))
	if err2 != nil {
		return shim.Error(err2.Error())
	}

	// Return a successful message
	return shim.Success(nil)
}

// Invoke all requests
func (t *ChannelChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	// Get the function and arguments from the request
	function, args := stub.GetFunctionAndParameters()

	fmt.Println("Provider invoke function: " + function)
	permission := t.queryPermission(stub)

	// Handle different functions
	if permission {
		if function == "queryHello" {
			return t.queryHello(stub, args)
		} else if function == "queryHelloProvider" {
			return t.queryHelloProvider(stub, args)
		} else if function == "queryPatientData" {
			return t.queryPatientData(stub, args)
		} else if function == "queryPatientRecords" {
			return t.queryPatientRecords(stub, args)
		}
	}

	fmt.Println("invoke did not find func: " + function) //error
	return shim.Error("Received unknown function invocation")

	// // Check whether the number of arguments is sufficient
	// if len(args) < 1 {
	// 	return shim.Error("The number of arguments is insufficient.")
	// }
	//
	// // In order to manage multiple type of request, we will check the first argument.
	// // Here we have one possible argument: query (every query request will read in the ledger without modification)
	// if args[0] == "query" {
	// 	return t.query(stub, args)
	// }
	//
	// // The update argument will manage all updates in the ledger
	// if args[0] == "invoke" {
	// 	return t.invoke(stub, args)
	// }
	//
	// // If the arguments given don’t match any function, we return an error
	// return shim.Error("Unknown action, check the first argument")
}

func (t *ChannelChaincode) queryPermission(stub shim.ChaincodeStubInterface) bool {
	fmt.Println("getting permission to read")

	// Get the state of the value matching the key hello in the ledger
	state, err := stub.GetState("permission")
	if err != nil {
		fmt.Println("Failed to get state of permission")
		return false
	}

	if string(state) == "true" {
		return true
	}
	return false
}

// queryHello: returns value of hello state
func (t *ChannelChaincode) queryHello(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("########### queryHello ###########")

	// Get the state of the value matching the key hello in the ledger
	state, err := stub.GetState("hello")
	if err != nil {
		return shim.Error("Failed to get state of hello")
	}

	// Return this value in response
	return shim.Success(state)
}

// queryHello: returns value of hello state
func (t *ChannelChaincode) queryHelloProvider(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("########### queryHelloProvider ###########")

	// Get the state of the value matching the key hello in the ledger
	state, err := stub.GetState("helloprovider")
	if err != nil {
		return shim.Error("Failed to get state of helloprovider")
	}

	// Return this value in response
	return shim.Success(state)
}

// queryPatientData: returns an array of DiscoveryBlocks for patientAddr
func (t *ChannelChaincode) queryPatientData(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("########### queryPatientData ###########")

	// Check whether the number of arguments is sufficient
	if len(args) != 1 {
		return shim.Error("The number of arguments is insufficient.")
	}

	// Get the state of the value matching the patientAddr in the ledger
	fmt.Println("looking up all directory blocks for patientAddr")
	state, err := stub.GetState("helloprovider")
	if err != nil {
		return shim.Error("Failed to queryPatientData")
	}

	// Return this value in response
	return shim.Success(state)
}

// queryPatientData: returns an array of DiscoveryBlocks for patientAddr
func (t *ChannelChaincode) queryPatientRecords(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("########### queryPatientRecords ###########")

	// Check whether the number of arguments is sufficient
	if len(args) < 1 {
		return shim.Error("The number of arguments is insufficient.")
	}

	// Get the state of the value matching the patientAddr in the ledger
	fmt.Println("looking up directory blocks for all recordIds")
	fmt.Println("running assembleFile on each recordId")
	fmt.Println("getting list of fileFragments from directory block")
	fmt.Println("calling Storage's queryFileFragment(fragmentId string)")
	fmt.Println("getting back an encrypted fragment")
	fmt.Println("assembling all encryptedFrag into encryptedFile")
	fmt.Println("returning encryptedFile")
	state, err := stub.GetState("helloprovider")
	if err != nil {
		return shim.Error("Failed to queryPatientRecords")
	}

	// Return this value in response
	return shim.Success(state)
}

// invoke
// Every functions that read and write in the ledger will be here
func (t *ChannelChaincode) invoke(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("########### ChannelChaincode invoke ###########")

	if len(args) < 2 {
		return shim.Error("The number of arguments is insufficient.")
	}

	// Check if the ledger key is "hello" and process if it is the case. Otherwise it returns an error.
	if args[1] == "helloprovider" && len(args) == 3 {

		// Write the new value in the ledger
		err := stub.PutState("helloprovider", []byte(args[2]))
		if err != nil {
			return shim.Error("Failed to update state of helloprovider")
		}

		// Notify listeners that an event "eventInvoke" have been executed (check line 19 in the file invoke.go)
		err = stub.SetEvent("eventInvoke", []byte{})
		if err != nil {
			return shim.Error(err.Error())
		}

		// Return this value in response
		return shim.Success(nil)
	}

	// If the arguments given don’t match any function, we return an error
	return shim.Error("Unknown invoke action, check the second argument.")
}

func main() {
	// Start the chaincode and make it ready for futures requests
	err := shim.Start(new(ChannelChaincode))
	if err != nil {
		fmt.Printf("Error starting provider chaincode: %s", err)
	}
}

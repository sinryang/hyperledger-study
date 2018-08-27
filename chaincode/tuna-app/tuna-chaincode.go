// SPDX-License-Identifier: Apache-2.0

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Define the Smart Contract structure
type SmartContract struct {
}

/* Define Tuna structure, with 4 properties.
Structure tags are used by encoding/json library
*/
type Tuna struct {
	Vessel string `json:"vessel"`
	Timestamp string `json:"timestamp"`
	Location  string `json:"location"`
	Holder  string `json:"holder"`
}


func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}


func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	return shim.Success(nil)
}

func (s *SmartContract) queryTuna(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	return shim.Success(nil)
}

func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	tuna := []Tuna{
		Tuna{Vessel: "923F", Location: "67.0006, -70.5476", Timestamp: "1504054225", Holder: "Miriam"},
		Tuna{Vessel: "M83T", Location: "91.2395, -49.4594", Timestamp: "1504057825", Holder: "Dave"},
		Tuna{Vessel: "T012", Location: "58.0148, 59.01391", Timestamp: "1493517025", Holder: "Igor"},
		Tuna{Vessel: "P490", Location: "-45.0945, 0.7949", Timestamp: "1496105425", Holder: "Amalea"},
		Tuna{Vessel: "S439", Location: "-107.6043, 19.5003", Timestamp: "1493512301", Holder: "Rafa"},
		Tuna{Vessel: "J205", Location: "-155.2304, -15.8723", Timestamp: "1494117101", Holder: "Shen"},
		Tuna{Vessel: "S22L", Location: "103.8842, 22.1277", Timestamp: "1496104301", Holder: "Leila"},
		Tuna{Vessel: "EI89", Location: "-132.3207, -34.0983", Timestamp: "1485066691", Holder: "Yuan"},
		Tuna{Vessel: "129R", Location: "153.0054, 12.6429", Timestamp: "1485153091", Holder: "Carlo"},
		Tuna{Vessel: "49W4", Location: "51.9435, 8.2735", Timestamp: "1487745091", Holder: "Fatima"},
	}

}

func (s *SmartContract) recordTuna(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	return shim.Success(nil)
}

func (s *SmartContract) queryAllTuna(APIstub shim.ChaincodeStubInterface) sc.Response {

	startKey := "0"
	endKey := "999"

	resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add comma before array members,suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- queryAllTuna:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

func (s *SmartContract) changeTunaHolder(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	return shim.Success(nil)
}

func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}

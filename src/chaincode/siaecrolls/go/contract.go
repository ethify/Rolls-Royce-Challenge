/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

/*
 * The sample smart contract for documentation topic:
 * Writing Your First Blockchain Application
 */

package main

/* Imports
 * 4 utility libraries for formatting, handling bytes, reading and writing JSON, and string manipulation
 * 2 specific Hyperledger Fabric specific libraries for Smart Contracts
 */
import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
	"strconv"
)

// Define the Smart Contract structure
type SmartContract struct {
}

type Request struct{
	EntryID string `json:"EntryID"`
	Requestor string `json:"Requestor"`
	UrgentRequest bool `json:"UrgentRequest"`
	LocationTask string `json: "LocationTask"`
	technicalDetails TechnicalDetails `json:"TechnicalDetails"`
	action Action `json:"Action"`
	requestDetails RequestDetails `json:"RequestDetails"`



}
type TechnicalDetails struct {
	creationDate string `json:"creationDate"`
	AcTrail string `json:"AcTrail"`
	EnginePositon string `json:"EnginePositon"`
	ESN string `json:"ESN"`
}
type Action struct {
	LineAdditions int `json:"lineAdditions"`
	Actions string `json:"Actions"`
	MaterialsNeeded string `json:"MaterialsNeeded"`
	SpecialInstruction string `json:"SpecialInstruction"`
	HangerAdditon string `json:"HangerAdditon"`
	AircraftCheck string `json:"AircraftCheck"`
}


type RequestDetails struct {
	subject string `json:"subject"`
	summary string `json:"summary"`
	docRef string  `json:"docRef"`
	e4Notif string `json:"e4Notif"`
	e8Notif string `json:"e8Notif"`
	componentChanges string `json:"componentChanges"`
	equipmentNumber int `json:"equipmentNumber"`
	pnbool bool `json:"pnbool"`
	pnString bool `json:"pnString"`
	bfeShortage string `json:"bfeShortage"`


}




/*
 * The Init method is called when the Smart Contract "fabcar" is instantiated by the blockchain network
 * Best practice is to have any Ledger initialization in separate function -- see initLedger()
 */
func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

/*
 * The Invoke method is called as a result of an application request to run the Smart Contract "fabcar"
 * The calling application program has also specified the particular smart contract function to be called, with arguments
 */
func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	fmt.Printf("len=%d cap=%d %v\n", len(args), cap(args), args)
	//return shim.Error("name:"+function)
	//Route to the appropriate handler function to interact with the ledger appropriately
	if function == "CreateWorkRequest" {
		return s.CreateWorkRequest(APIstub,args)
	}else if function == "queryRequests" {
		return s.queryRequests(APIstub,args)
	}

	return shim.Error("Invalid Smart Contract function name.")
}
func (s *SmartContract) CreateWorkRequest(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1{
		return shim.Error("Incorrect number of arguments. Expecting 2")//First argument will be entryId and second will be JSON having entryID again
	}
	byteToInt, _ := strconv.Atoi(string(count))
	//var car = Car{Make: args[1], Model: args[2], Colour: args[3], Owner: args[4]}
	count := APIstub.GetState("RequestCount")
	if len(count)>0 {
		byteToInt = byteToInt +1
		APIstub.PutState("RequestCount", []byte(byteToInt))
	}
	else{
		APIstub.PutState("RequestCount", []byte( byteToInt))
	}
	APIstub.PutState("req"+strconv.Itoa(byteToInt), []byte( args[1]))
	return shim.Success(nil)
}

func (s *SmartContract) queryRequests(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	R, _ := APIstub.GetState(args[0])
	return shim.Success(R)
	
	
}
func (s *SmartContract) CreateWorkInstructions(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1{
		return shim.Error("Incorrect number of arguments. Expecting 2")//First argument will be entryId and second will be JSON having entryID again
	}
	byteToInt, _ := strconv.Atoi(string(count))
	//var car = Car{Make: args[1], Model: args[2], Colour: args[3], Owner: args[4]}
	count := APIstub.GetState("workCount")
	if len(count)>0 {
		byteToInt = byteToInt +1
		APIstub.PutState("workCount", []byte(byteToInt))
	}
	else{
		APIstub.PutState("workCount", []byte( byteToInt))
	}
	APIstub.PutState("work"+strconv.Itoa(byteToInt), []byte( args[1]))
	return shim.Success(nil)
}

func (s *SmartContract) queryWork(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	R, _ := APIstub.GetState(args[0])
	return shim.Success(R)
	
}
func (s *SmartContract) queryRequestcount(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	R, _ := APIstub.GetState([]byte("RequestCount"))
	return shim.Success(R)
	
}
func (s *SmartContract) queryWorkCount(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	R, _ := APIstub.GetState([]byte("workCount"))
	return shim.Success(R)
	
}
func (s *SmartContract) assign (APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	f len(args) != 1{
		return shim.Error("Incorrect number of arguments. Expecting 2")//First argument will be request no and second will be rrsquestor
	}
	byteToInt, _ := strconv.Atoi(string(count))
	//var car = Car{Make: args[1], Model: args[2], Colour: args[3], Owner: args[4]}
	
	APIstub.PutState("select"+strconv.Itoa(byteToInt), []byte( args[1]))
	return shim.Success(nil)
	
}
func (s *SmartContract) showSelected (APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	f len(args) != 1{
		return shim.Error("Incorrect number of arguments. Expecting 2")//First argument will be request no and second will be rqruestor
	}
	R, _ := APIstub.GetState([]byte(args[0]))
	return shim.Success(R)
}
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}

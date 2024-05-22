package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type TestContract struct {
	contractapi.Contract
}

type Asset struct {
	Id    string `json:Id`
	Owner string `json:Owner`
}

// Struct For Storing Transaction Metadata
type TransactionReceipt struct {
	Transaction_Creator   string
	Transaction_Timestamp *timestamp.Timestamp
	Transaction_Id        string
	Client_Id             string
	Channel_Id            string
}

func (contract *TestContract) InitAsset(ctx contractapi.TransactionContextInterface) error {

	/*
		InitAsset() Method Is Write According To The Chaincode
		Becasue It is Writing To The Ledger And A Write Transaction In HLF
		Only Returns Ststus Code Or Error In Case Of Failure

		func Init_Assets() is method of the Asset Struct
		Usage: It Adds Dummy  Assets To The Legder
		@param: 'ctx'  is type of 'contractapi.TransactionContextInterface' Which Provide all the
		ShimAPI Methods To Interact Within The Networks

		returns: error / nil as PutState() returns only error or nil


	*/

	// dummy asset
	assets := []Asset{
		{
			Id:    "1",
			Owner: "Alesso",
		},

		{
			Id:    "2",
			Owner: "Coldplay",
		},
	}

	for _, asset := range assets {

		json_Asset, err := json.Marshal(asset)

		if err != nil {
			return fmt.Errorf("Failed To Marshal Asset %v", err)
		}

		// it calls the shimpapi method over getstub() method to putstate()
		// to add the asset to ledger
		// PutState() method retquires two params
		// one is the key and its type must be a string
		// it is the key by which we can query over the ledger for an asset
		// another is the whole reprentation of the asset in json Format
		// It returns error or nil

		err = ctx.GetStub().PutState(asset.Id, json_Asset)
		if err != nil {
			return fmt.Errorf("Failed To Add Asset To The Ledger")
		}
	}

	return nil
}

func main() {

	test_chaincode, err := contractapi.NewChaincode(&TestContract{})

	if err != nil {
		log.Panicf("Failed To Initiate Chaincode Error: %v", err)
	}
	err = test_chaincode.Start()

	if err != nil {
		log.Panicf("Failed To Start The Chaincode Error: %v", err)

	}

}

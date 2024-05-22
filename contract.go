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
	Transaction_Creator   string               `json:trx_Creator`
	Transaction_Timestamp *timestamp.Timestamp `json: trx_timestamp`
	Transaction_Id        string               `json: trx_id`
	Client_Id             string               `json: client_id`
	Channel_Id            string               `json:channel_id`
}

func (contract *TestContract) Init_Asset(ctx contractapi.TransactionContextInterface) error {

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

func (contract *TestContract) Create_Asset(ctx contractapi.TransactionContextInterface, _id string, _owner string) (*TransactionReceipt, error) {
	/* 	Create Asset Chaincode Method To Create New Asset
	@params _id is the id of the asset
	@params _owner owner of the asset
	return: on successfull it will return the transationreceipt object which will contain all information for this transaction
	*/
	is_exists, err := contract.Has_Asset(ctx, _id)

	if err != nil || is_exists == true {
		return nil, fmt.Errorf("Asset Already Exists With The Id : %v", _id)
	}

	asset := Asset{
		Id:    _id,
		Owner: _owner,
	}

	// Json Parsing

	json_asset, err := json.Marshal(asset)

	if err != nil {
		return nil, fmt.Errorf("Failed to parsing json error : %v", err)
	}

	err = ctx.GetStub().PutState(asset.Id, json_asset)

	if err != nil {
		return nil, fmt.Errorf("Failed to put asset to ledger : %v", err)
	}

	trx_creator, err := ctx.GetClientIdentity().GetID()

	if err != nil {
		return nil, fmt.Errorf("Failed to get transaction creator : %v", err)
	}

	trx_timestamp, err := ctx.GetStub().GetTxTimestamp()

	if err != nil {
		return nil, fmt.Errorf("Failed to get transaction timestamp : %v", err)
	}

	trx_id := ctx.GetStub().GetTxID()

	client_id, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return nil, fmt.Errorf("Failed to get Client : %v", err)
	}

	channel_id := ctx.GetStub().GetChannelID()

	transaction_receipt := TransactionReceipt{
		Transaction_Creator:   trx_creator,
		Transaction_Timestamp: trx_timestamp,
		Transaction_Id:        trx_id,
		Client_Id:             client_id,
		Channel_Id:            channel_id,
	}

	return &transaction_receipt, nil

}

func (contract *TestContract) Has_Asset(ctx contractapi.TransactionContextInterface, _id string) (bool, error) {
	/*
		Has_Asset is a helper function Which checks if a asset exists or not in the ledger
		@params _id is the id of the asset which is being checked for its existence

		returns: true and nil if the asset exists else false if the asset does not exists
	*/
	var found bool
	asset, err := ctx.GetStub().GetState(_id)

	if err != nil {
		return found, fmt.Errorf("Failed To Read From The Ledger")
	}

	if asset != nil {
		found = true
	}

	return found, nil
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

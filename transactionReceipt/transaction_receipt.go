package transactionReceipt

import (
	"fmt"
	"time"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type TransactionReceipt struct {
	Transaction_Creator   string    `json:trx_Creator`
	Transaction_Timestamp time.Time `json: trx_timestamp`
	Transaction_Id        string    `json: trx_id`
	Client_Id             string    `json: client_id`
	Channel_Id            string    `json:channel_id`
}

func Get_Transaction_Receipt(ctx contractapi.TransactionContextInterface) (*TransactionReceipt, error) {
	trx_creator, err := ctx.GetClientIdentity().GetID()

	if err != nil {
		return nil, fmt.Errorf("Failed to get transaction creator : %v", err)
	}

	trx_timestamp, err := ctx.GetStub().GetTxTimestamp()

	if err != nil {
		return nil, fmt.Errorf("Failed to get transaction timestamp : %v", err)
	}
	utc_time := time.Unix(trx_timestamp.Seconds, 0)

	trx_id := ctx.GetStub().GetTxID()

	client_id, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return nil, fmt.Errorf("Failed to get Client : %v", err)
	}

	channel_id := ctx.GetStub().GetChannelID()

	transaction_receipt := TransactionReceipt{
		Transaction_Creator:   trx_creator,
		Transaction_Timestamp: utc_time,
		Transaction_Id:        trx_id,
		Client_Id:             client_id,
		Channel_Id:            channel_id,
	}

	return &transaction_receipt, nil

}

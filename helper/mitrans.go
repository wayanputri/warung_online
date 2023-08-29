package helper

import (
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"strconv"
	"warung_online/app/config"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

func RequestCreditCard(harga string,orderId string, bank string) *coreapi.ChargeResponse{

	cfg := config.InitConfig()
	midtrans.ServerKey = cfg.KEY_SERVER_MIDTRANS
	authString := EncodeAuthString(midtrans.ServerKey, "")
	fmt.Println("AUTH_STRING:", authString)
	midtrans.Environment = midtrans.Sandbox

	totalHarga,_ := strconv.Atoi(harga)

	bankTransferReq := &coreapi.ChargeReq{
		PaymentType:        coreapi.PaymentTypeBankTransfer,
		TransactionDetails: midtrans.TransactionDetails{OrderID: orderId, GrossAmt: int64(totalHarga)}, 
		BankTransfer:       &coreapi.BankTransferDetails{Bank: midtrans.Bank(bank)},
		Metadata:           nil,
	}

	coreApiRes, errCore := coreapi.ChargeTransaction(bankTransferReq)
	if errCore != nil {
		log.Fatal("Failed to charge transaction:", errCore)
	}
	return coreApiRes
}
func EncodeAuthString(username, password string) string {
	auth := username + ":" + password
	authBytes := []byte(auth)
	encodedAuth := base64.StdEncoding.EncodeToString(authBytes)
	return encodedAuth
}

func SetMitrans()coreapi.Client{
	cfg := config.InitConfig()
	var client = coreapi.Client{}
	client.New(cfg.KEY_SERVER_MIDTRANS, midtrans.Sandbox)

	client.Options.SetPaymentAppendNotification("YOUR-APPEND-NOTIFICATION-ENDPOINT")
	return client
}

func OrderIdMitrans(notificationPayload map[string]interface{},client coreapi.Client)(*coreapi.TransactionStatusResponse,string,error){
	orderID, exists := notificationPayload["order_id"].(string)
	if !exists {
	
		return &coreapi.TransactionStatusResponse{},"",errors.New("failed get orderId")
	}
	transactionStatusResp, errTrans := client.CheckTransaction(orderID) 
	if errTrans != nil {
		return &coreapi.TransactionStatusResponse{},"",errors.New("failed check status transaction")
	}
	return transactionStatusResp,orderID,nil
}

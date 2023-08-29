package service

import (
	"errors"
	"strconv"
	"warung_online/features/payment"
	"warung_online/features/structsEntity"
	"warung_online/helper"
)

type PaymentService struct {
	payment payment.PaymentDataInterface
}

// Notification implements payment.PaymentServiceInterface.
func (service *PaymentService) Notification(notificationPayload map[string]interface{}) (structsEntity.PaymentEntity, error) {
	client:=helper.SetMitrans()

	response,orderId,errOsrder:=helper.OrderIdMitrans(notificationPayload,client)
	if errOsrder != nil{
		return structsEntity.PaymentEntity{},errors.New("failed order id")
	}
	id,err:=service.payment.UpdatePayment(response.TransactionStatus,orderId)
	if err != nil{
		return structsEntity.PaymentEntity{},errors.New("update payment failed")
	}
	data,errGet:=service.payment.SelectPaymentById(id)
	if errGet != nil{
		return structsEntity.PaymentEntity{},errors.New("failed get payment update")
	}
	return data,nil
}

// Add implements payment.PaymentServiceInterface.
func (service *PaymentService) Add(input structsEntity.PaymentEntity) (structsEntity.PaymentEntity, error) {
	dataTransactionDetil, errTransactionDetil := service.payment.SelectTransactionDetilById(input.TransactionFinalID)
	if errTransactionDetil != nil {
		return structsEntity.PaymentEntity{}, errors.New("error get data transaction detil")
	}
	totalHargaConv := strconv.Itoa(dataTransactionDetil.TotalHarga)
	response := helper.RequestCreditCard(totalHargaConv, dataTransactionDetil.OrderID, input.Bank)
	dataInputResponse := structsEntity.ResponseMitrans(input, input.TransactionFinalID, response)
	id, err := service.payment.Insert(dataInputResponse)
	if err != nil {
		return structsEntity.PaymentEntity{}, errors.New("error create payment")
	}
	dataPayment, errPayment := service.payment.SelectPaymentById(id)
	if errPayment != nil {
		return structsEntity.PaymentEntity{}, errors.New("error get payment")
	}
	return dataPayment, nil
}

func New(payment payment.PaymentDataInterface) payment.PaymentServiceInterface {
	return &PaymentService{
		payment: payment,
	}
}

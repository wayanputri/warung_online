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

// UpdateStok implements payment.PaymentServiceInterface.
func (service *PaymentService) UpdateStok(idTransactionFinal uint) error {
	idTransaction,errTransaction:=service.payment.SelectTransactionDetil(idTransactionFinal)
	if errTransaction != nil{
		return errTransaction
	}
	dataTransaction,idProduct,errProduct:=service.payment.SelectTransaction(idTransaction)
	if errProduct !=nil{
		return errProduct
	}

	_,dataProduct,errStok:=service.payment.SelectProduct(idProduct)
	if errStok != nil{
		return errStok
	}
	var stokNew []int
	for _,value:=range dataProduct{
		for _,value1:=range dataTransaction{
			if value1.ProductID==value.Id{
				stok:=value.Stok-value1.Jumlah
				stokNew=append(stokNew, stok)
			}
		}
	}
	var productEntity []structsEntity.ProductEntity
	for _,value4:=range stokNew{
		for _,value5:=range dataProduct{
			value5.Stok=value4
			productEntity = append(productEntity, value5)
		}
	}
	errUp:=service.payment.UpdateProduct(productEntity)
	if errUp != nil{
		return errUp
	}
	return nil
}

// Notification implements payment.PaymentServiceInterface.
func (service *PaymentService) Notification(notificationPayload map[string]interface{}) (structsEntity.PaymentEntity, error) {
	client := helper.SetMitrans()

	response, orderId, errOsrder := helper.OrderIdMitrans(notificationPayload, client)
	if errOsrder != nil {
		return structsEntity.PaymentEntity{}, errors.New("failed order id")
	}
	id, err := service.payment.UpdatePayment(response.TransactionStatus, orderId)
	if err != nil {
		return structsEntity.PaymentEntity{}, errors.New("update payment failed")
	}
	data, errGet := service.payment.SelectPaymentById(id)
	if errGet != nil {
		return structsEntity.PaymentEntity{}, errors.New("failed get payment update")
	}

	// if data.Status != "settlement"{
	// 	return structsEntity.PaymentEntity{}, errors.New("pembayaran gagal")
	// }

	idTransaction,errTransaction:=service.payment.SelectTransactionDetil(data.TransactionFinalID)
	if errTransaction != nil{
		return structsEntity.PaymentEntity{},errTransaction
	}
	dataTransaction,idProduct,errProduct:=service.payment.SelectTransaction(idTransaction)
	if errProduct !=nil{
		return structsEntity.PaymentEntity{},errProduct
	}

	_,dataProduct,errStok:=service.payment.SelectProduct(idProduct)
	if errStok != nil{
		return structsEntity.PaymentEntity{},errStok
	}
	var stokNew []int
	for _,value:=range dataProduct{
		for _,value1:=range dataTransaction{
			if value1.ProductID==value.Id{
				stok:=value.Stok-value1.Jumlah
				stokNew=append(stokNew, stok)
			}
		}
	}
	var productEntity []structsEntity.ProductEntity
	for _,value4:=range stokNew{
		for _,value5:=range dataProduct{
			value5.Stok=value4
			productEntity = append(productEntity, value5)
		}
	}
	errUp:=service.payment.UpdateProduct(productEntity)
	if errUp != nil{
		return structsEntity.PaymentEntity{},errUp
	}
	return data, nil
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

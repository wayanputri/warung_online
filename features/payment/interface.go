package payment

import "warung_online/features/structsEntity"

type PaymentDataInterface interface {
	Insert(input structsEntity.PaymentEntity)(uint,error)
	SelectTransactionDetilById(idTransaction uint)(structsEntity.TransactionFinalEntity,error)
	SelectTransactionDetil(idTransaction uint)([]uint,error)
	SelectTransaction(id []uint)([]int,[]uint,error)
	SelectProduct(idProduct []uint)([]int,[]structsEntity.ProductEntity,error)
	UpdateProduct(input []int,id []uint)(error)
	SelectPaymentById(id uint)(structsEntity.PaymentEntity,error)
	UpdatePayment(accept string,orderId string)(uint,error)
}

type PaymentServiceInterface interface{
	Add(input structsEntity.PaymentEntity)(structsEntity.PaymentEntity,error)
	Notification(notificationPayload map[string]interface{})(structsEntity.PaymentEntity,error)
}
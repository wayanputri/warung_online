package review

import "warung_online/features/structsEntity"

type ReviewDataInterface interface {
	Insert(input structsEntity.ReviewEntity,idProduct uint)(uint,error)
	SelectReview(idReview uint)(structsEntity.ReviewEntity,error)
	SelectPaymentId(idPayment uint)(uint,string,error)
	SelectKeranjang(idTransactionDetil uint)([]uint,error)
	SelectTransaction(idTransaction []uint)([]uint,error)
	UpdateProduct(input float64,idUser uint,idproduct uint)(uint,error)
	SelectProduct(idProduct []uint,idPrd uint)(error)
	SelectProductReviewById(idProduct uint)([]structsEntity.ReviewEntity,error)
	Delete(idReview uint)error
}

type ReviewServiceInterface interface{
	Add(input structsEntity.ReviewEntity,idUser uint)(structsEntity.ReviewEntity,error)
	Delete(idReview uint,userId uint)error
}
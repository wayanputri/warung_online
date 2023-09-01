package service

import (
	"errors"
	"warung_online/features/review"
	"warung_online/features/structsEntity"
)

type ReviewService struct {
	reviewService review.ReviewDataInterface
}

// Delete implements review.ReviewServiceInterface.
func (service *ReviewService) Delete(idReview uint,userId uint) error {
	data,errReview:=service.reviewService.SelectReview(idReview)
	if errReview != nil{
		return errReview
	}
	dataAll,errAll:=service.reviewService.SelectProductReviewById(data.ProductID)
	if errAll != nil{
		return errAll
	}
	var count int=0
	var rating float64=0
	for _,value:=range dataAll{
		count++
		rating +=value.Rating
	}
	average:=(rating-data.Rating)/float64(count-1)
	_,errUpdate:=service.reviewService.UpdateProduct(average,userId,data.ProductID)
	if errUpdate != nil{
		return errUpdate
	}
	err:=service.reviewService.Delete(idReview)
	if err != nil{
		return err
	}
	return nil
}

// Add implements review.ReviewServiceInterface.
func (service *ReviewService) Add(input structsEntity.ReviewEntity, idUser uint) (structsEntity.ReviewEntity, error) {
	idTransFinal, status, errPayment := service.reviewService.SelectPaymentId(input.PaymentID)
	if errPayment != nil {
		return structsEntity.ReviewEntity{}, errPayment
	}
	if status != "settlement" {
		return structsEntity.ReviewEntity{}, errors.New("tidak dapat menambah review, product blm dibayar")
	}

	idTrans, errKeranjang := service.reviewService.SelectKeranjang(idTransFinal)
	if errKeranjang != nil {
		return structsEntity.ReviewEntity{}, errKeranjang
	}
	idProduct, errProduct := service.reviewService.SelectTransaction(idTrans)
	if errProduct != nil {
		return structsEntity.ReviewEntity{}, errProduct
	}
	errPrd := service.reviewService.SelectProduct(idProduct, input.ProductID)
	if errPrd != nil {
		return structsEntity.ReviewEntity{}, errPrd
	}
	idReview, errAdd := service.reviewService.Insert(input, input.ProductID)
	if errAdd != nil {
		return structsEntity.ReviewEntity{}, errAdd
	}
	dataReview, errReview := service.reviewService.SelectReview(idReview)
	if errReview != nil {
		return structsEntity.ReviewEntity{}, errReview
	}
	dataallReview, errAllReview := service.reviewService.SelectProductReviewById(dataReview.ProductID)
	if errAllReview != nil {
		return structsEntity.ReviewEntity{}, errAllReview
	}
	var count int = 0
	var rating float64 = 0
	for _, value := range dataallReview {
		count++
		rating += value.Rating
	}
	average := rating / float64(count)
	_, errUpdateRating := service.reviewService.UpdateProduct(average, idUser, dataReview.ProductID)
	if errUpdateRating != nil {
		return structsEntity.ReviewEntity{}, errUpdateRating
	}
	return dataReview, nil
}

func New(review review.ReviewDataInterface) review.ReviewServiceInterface {
	return &ReviewService{
		reviewService: review,
	}
}

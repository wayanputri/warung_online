package data

import (
	"errors"
	"warung_online/features/review"
	"warung_online/features/structsEntity"

	"gorm.io/gorm"
)

type ReviewData struct {
	db *gorm.DB
}

// Delete implements review.ReviewDataInterface.
func (repo *ReviewData) Delete(idReview uint) error {
	var review structsEntity.Review
	tx:=repo.db.Delete(&review,idReview)
	if tx.Error != nil{
		return errors.New("error deleted")
	}
	return nil
}

// SelectProductReviewById implements review.ReviewDataInterface.
func (repo *ReviewData) SelectProductReviewById(idProduct uint) ([]structsEntity.ReviewEntity, error) {
	var review []structsEntity.Review
	tx := repo.db.Where("product_id=?", idProduct).Find(&review)
	if tx.Error != nil {
		return nil, errors.New("failed review all by product id")
	}
	var output []structsEntity.ReviewEntity
	for _, value := range review {
		output = append(output, structsEntity.ReviewModelToEntity(value))
	}
	return output, nil
}

// SelectProduct implements review.ReviewDataInterface.
func (repo *ReviewData) SelectProduct(idProduct []uint, idPrd uint) error {
	var product []structsEntity.Product
	tx := repo.db.Where("id IN ?", idProduct).Find(&product)
	if tx.Error != nil {
		return errors.New("failed get select produt all")
	}
	txx := repo.db.First(&product, idPrd)
	if txx.Error != nil {
		return errors.New("failed product not found")
	}
	return nil
}

// Insert implements review.ReviewDataInterface.
func (repo *ReviewData) Insert(input structsEntity.ReviewEntity, idProduct uint) (uint, error) {
	inputModel := structsEntity.ReviewEntityToModel(input)
	inputModel.ProductID = idProduct
	tx := repo.db.Create(&inputModel)
	if tx.Error != nil {
		return 0, errors.New("failed insert review")
	}
	return inputModel.ID, nil
}

// SelectKeranjang implements review.ReviewDataInterface.
func (repo *ReviewData) SelectKeranjang(idTransactionDetil uint) ([]uint, error) {
	var keranjang []structsEntity.TransactionKeranjang
	tx := repo.db.Where("transaction_final_id = ?", idTransactionDetil).Find(&keranjang)
	if tx.Error != nil {
		return nil, errors.New("error get id transaction")
	}
	var idTransaction []uint
	for _, value := range keranjang {
		idTransaction = append(idTransaction, value.TransactionID)
	}
	return idTransaction, nil
}

// SelectPaymentId implements review.ReviewDataInterface.
func (repo *ReviewData) SelectPaymentId(idPayment uint) (uint, string, error) {
	var payment structsEntity.Payment
	tx := repo.db.First(&payment, idPayment)
	if tx.Error != nil {
		return 0, "", errors.New("failed select payment")
	}
	return payment.TransactionFinalID, payment.Status, nil
}

// SelectReview implements review.ReviewDataInterface.
func (repo *ReviewData) SelectReview(idReview uint) (structsEntity.ReviewEntity, error) {
	var review structsEntity.Review
	tx := repo.db.Preload("ImageReviews").First(&review, idReview)
	if tx.Error != nil {
		return structsEntity.ReviewEntity{}, errors.New("failed select review")
	}
	output := structsEntity.ReviewModelToEntity(review)
	return output, nil
}

// SelectTransaction implements review.ReviewDataInterface.
func (repo *ReviewData) SelectTransaction(idTransaction []uint) ([]uint, error) {
	var transaction []structsEntity.Transaction
	tx := repo.db.Where("id IN ?", idTransaction).Find(&transaction)
	if tx.Error != nil {
		return nil, errors.New("failed get id product")
	}
	var idProduct []uint
	for _, value := range transaction {
		idProduct = append(idProduct, value.ProductID)
	}
	return idProduct, nil
}

// UpdateProduct implements review.ReviewDataInterface.
func (repo *ReviewData) UpdateProduct(input float64, idUser uint, idproduct uint) (uint, error) {
	var transaction structsEntity.Transaction
	txx := repo.db.First(&transaction, "user_id=?", idUser)
	if txx.Error != nil {
		return 0, errors.New("anda harus transaksi dulu agar bisa mengedit product")
	}
	var updateReview structsEntity.Product
	updateReview.Ratings = input
	tx := repo.db.Model(&structsEntity.Product{}).Where("id=?", idproduct).Updates(updateReview)
	if tx.Error != nil {
		return 0, errors.New("failed update product")
	}
	return idproduct, nil
}

func New(db *gorm.DB) review.ReviewDataInterface {
	return &ReviewData{
		db: db,
	}
}

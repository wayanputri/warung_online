package data

import (
	"warung_online/features/payment"
	"warung_online/features/structsEntity"

	"gorm.io/gorm"
)

type PaymentData struct {
	db *gorm.DB
}

// UpdatePayment implements payment.PaymentDataInterface.
func (repo *PaymentData) UpdatePayment(accept string, orderId string) (uint, error) {
	var transaction structsEntity.TransactionFinal
	txx:=repo.db.Where("order_id=?",orderId).First(&transaction)
	if txx.Error !=nil{
		return 0,txx.Error
	}
	var payment structsEntity.Payment
	payment.Status=accept
	tx:=repo.db.Model(&structsEntity.Payment{}).Where("transaction_final_id=?",transaction.ID).Updates(payment)
	if tx.Error !=nil{
		return 0,txx.Error
	}
	return payment.TransactionFinalID,nil
}

// Insert implements payment.PaymentDataInterface.
func (repo *PaymentData) Insert(input structsEntity.PaymentEntity) (uint, error) {
	inputModel := structsEntity.PaymentEntityToModel(input)
	tx := repo.db.Create(&inputModel)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return inputModel.ID, nil
}

// SelectPaymentById implements payment.PaymentDataInterface.
func (repo *PaymentData) SelectPaymentById(id uint) (structsEntity.PaymentEntity, error) {
	var payment structsEntity.Payment
	tx := repo.db.Preload("TransactionFinals").First(&payment, id)
	if tx.Error != nil {
		return structsEntity.PaymentEntity{}, tx.Error
	}
	output := structsEntity.PaymentModelToEntity(payment)
	return output, nil
}

// SelectProduct implements payment.PaymentDataInterface.
func (repo *PaymentData) SelectProduct(idProduct []uint) ([]int, []structsEntity.ProductEntity, error) {
	var productModel []structsEntity.Product
	tx := repo.db.Where("id IN ?", idProduct).Find(&productModel)
	if tx.Error != nil {
		return nil, nil, tx.Error
	}
	var jumlahStok []int
	for _, value1 := range productModel {
		jumlahStok = append(jumlahStok, value1.Stok)
	}
	var output []structsEntity.ProductEntity
	for _, value2 := range productModel {
		output = append(output, structsEntity.ProductModelToEntity(value2))
	}
	return jumlahStok, output, nil
}

// SelectTransaction implements payment.PaymentDataInterface.
func (repo *PaymentData) SelectTransaction(id []uint) ([]int, []uint, error) {
	var transansactionModel []structsEntity.Transaction
	tx := repo.db.Where("id IN ?", id).Find(&transansactionModel)
	if tx.Error != nil {
		return nil, nil, tx.Error
	}
	var jumlah []int
	for _, value1 := range transansactionModel {
		jumlah = append(jumlah, value1.Jumlah)
	}
	var productId []uint
	for _, value2 := range transansactionModel {
		productId = append(productId, value2.ProductID)
	}
	return jumlah, productId, nil
}

// SelectTransactionDetil implements payment.PaymentDataInterface.
func (repo *PaymentData) SelectTransactionDetil(idTransaction uint) ([]uint, error) {
	var transansactionModel []structsEntity.TransactionKeranjang
	tx := repo.db.Where("transaction_final_id=?", idTransaction).Find(&transansactionModel)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var ids []uint
	for _, value := range transansactionModel {
		ids = append(ids, value.TransactionID)
	}
	//id transaction
	return ids, nil
}

// SelectTransactionDetilById implements payment.PaymentDataInterface.
func (repo *PaymentData) SelectTransactionDetilById(idTransaction uint) (structsEntity.TransactionFinalEntity, error) {
	var transaction structsEntity.TransactionFinal
	tx := repo.db.First(&transaction, idTransaction)
	if tx.Error != nil {
		return structsEntity.TransactionFinalEntity{}, tx.Error
	}
	output := structsEntity.TransactionFinalModelToEntity(transaction)
	return output, nil
}

// UpdateProduct implements payment.PaymentDataInterface.
func (repo *PaymentData) UpdateProduct(input []int, id []uint) error {
	var product []structsEntity.Product
	tx := repo.db.Where("id IN ?", id).Find(&product)
	if tx.Error != nil {
		return tx.Error
	}
	for i, value := range input {
		if i >= len(product) {
			break
		}
		product[i].Stok = value
		txx := repo.db.Save(&product[i])
		if txx.Error != nil {
			return txx.Error
		}
	}
	return nil
}

func New(db *gorm.DB) payment.PaymentDataInterface {
	return &PaymentData{
		db: db,
	}
}

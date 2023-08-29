package data

import (
	"errors"
	"warung_online/features/structsEntity"
	"warung_online/features/transaction"

	"gorm.io/gorm"
)

type TransactionData struct {
	db *gorm.DB
}

// Delete implements transaction.TransactionDataInterface.
func (repo *TransactionData) Delete(id uint, idUser uint) error {
	var transaction structsEntity.Transaction
	tx:=repo.db.Where("user_id=? and id=?",idUser,id).Delete(&transaction)
	if tx.Error != nil {
		return errors.New("failed delete transaction")
	}
	if tx.RowsAffected == 0 {
		return errors.New("not row affected")
	}
	return nil
}

// Update implements transaction.TransactionDataInterface.
func (repo *TransactionData) Update(input structsEntity.TransactionEntity) (uint, error) {
	inputModel := structsEntity.TransactionEntityToModel(input)
	tx := repo.db.Model(&structsEntity.Transaction{}).Where("user_id=? and id=?", input.UserID, input.Id).Updates(inputModel)
	if tx.Error != nil {
		return 0, errors.New("failed update transaction")
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("not row affected")
	}
	return input.Id, nil
}

// SelectAll implements transaction.TransactionDataInterface.
func (repo *TransactionData) SelectAll(UserId uint) ([]structsEntity.TransactionEntity, error) {
	var transaction []structsEntity.Transaction
	tx := repo.db.Preload("Users").Preload("Products").Preload("Products.Users").Where("user_id=?", UserId).Find(&transaction)
	if tx.Error != nil {
		return nil, errors.New("failed get data transaction")
	}
	var transactionEntity []structsEntity.TransactionEntity
	for _, value := range transaction {
		transactionEntity = append(transactionEntity, structsEntity.TransactionModelToEntity(value))
	}
	return transactionEntity, nil
}

// Insert implements transaction.TransactionDataInterface.
func (repo *TransactionData) Insert(input structsEntity.TransactionEntity) (uint, error) {
	inputModel := structsEntity.TransactionEntityToModel(input)
	tx := repo.db.Create(&inputModel)
	if tx.Error != nil {
		return 0, errors.New("failed insert data transaction")
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("not row affected")
	}
	var keranjang structsEntity.TransactionKeranjang
	keranjang.TransactionID = inputModel.ID
	txx := repo.db.Create(&keranjang)
	if txx.Error != nil {
		return 0, errors.New("failed insert data keranjang")
	}
	if txx.RowsAffected == 0 {
		return 0, errors.New("not row affected")
	}
	return inputModel.ID, nil
}

// SelectById implements transaction.TransactionDataInterface.
func (repo *TransactionData) SelectById(idTransaction uint, UserId uint) (structsEntity.TransactionEntity, error) {
	var transaction structsEntity.Transaction
	tx := repo.db.Preload("Users").Preload("Products").Preload("Products.Users").Where("user_id=?", UserId).First(&transaction, idTransaction)
	if tx.Error != nil {
		return structsEntity.TransactionEntity{}, errors.New("failed get data transaction")
	}
	output := structsEntity.TransactionModelToEntity(transaction)
	return output, nil
}

// SelectByProduct implements transaction.TransactionDataInterface.
func (repo *TransactionData) SelectByProduct(idProduct uint) (structsEntity.ProductEntity, error) {
	var product structsEntity.Product
	tx := repo.db.First(&product, idProduct)
	if tx.Error != nil {
		return structsEntity.ProductEntity{}, errors.New("failed get data product")
	}
	output := structsEntity.ProductModelToEntity(product)
	return output, nil
}

func New(db *gorm.DB) transaction.TransactionDataInterface {
	return &TransactionData{
		db: db,
	}
}

package data

import (
	"errors"
	"warung_online/features/product"
	"warung_online/features/structsEntity"

	"gorm.io/gorm"
)

type ProductData struct {
	db *gorm.DB
}

// Searching implements product.ProductDataInterface.
func (data ProductData) Searching(name string) ([]structsEntity.ProductEntity, error) {
	var productModel []structsEntity.Product
	tx := data.db.Preload("Users").Preload("Image").Preload("Transaction").Where("nama_barang like ?", "%"+name+"%").Find(&productModel)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var productCore []structsEntity.ProductEntity
	for _, value := range productModel {
		productCore = append(productCore, structsEntity.ProductModelToEntity(value))
	}
	return productCore, nil

}

// SelectById implements product.ProductDataInterface.
func (data ProductData) SelectById(id uint) (structsEntity.ProductEntity, error) {
	var productModel structsEntity.Product
	tx := data.db.Preload("Users").Preload("Image").Preload("Transaction").Preload("Reviews").First(&productModel, id)
	if tx.Error != nil {
		return structsEntity.ProductEntity{}, tx.Error
	}

	productCore := structsEntity.ProductModelToEntity(productModel)

	return productCore, nil
}

// SelectAll implements product.ProductDataInterface.
func (data ProductData) SelectAll() ([]structsEntity.ProductEntity, error) {
	var productModel []structsEntity.Product
	tx := data.db.Preload("Users").Preload("Image").Preload("Transaction").Find(&productModel)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var productCore []structsEntity.ProductEntity
	for _, value := range productModel {
		productCore = append(productCore, structsEntity.ProductModelToEntity(value))
	}
	return productCore, nil
}

// Delete implements product.ProductDataInterface.
func (data ProductData) Delete(idProduct uint) error {
	var product structsEntity.Product
	tx := data.db.Delete(&product, idProduct)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("failed error delete data")
	}
	return nil
}

// Update implements product.ProductDataInterface.
func (data ProductData) Update(input structsEntity.ProductEntity, id uint) (uint, error) {
	productModel := structsEntity.ProductEntityToModel(input)
	tx := data.db.Model(&structsEntity.Product{}).Where("id=?", id).Updates(productModel)
	if tx.Error != nil {
		return 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("failed error update data")
	}
	return productModel.ID, nil
}

// SelectUser implements product.ProductDataInterface.
func (data ProductData) SelectUser(id uint) (structsEntity.UserEntity, error) {
	var userModel structsEntity.User
	tx := data.db.First(&userModel, id)
	if tx.Error != nil {
		return structsEntity.UserEntity{}, tx.Error
	}
	dataUser := structsEntity.UserModelToEntity(userModel)
	return dataUser, nil
}

// Insert implements product.ProductDataInterface.
func (data ProductData) Insert(input structsEntity.ProductEntity) (uint, error) {
	productModel := structsEntity.ProductEntityToModel(input)

	tx := data.db.Create(&productModel)
	if tx.Error != nil {
		return 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("failed error create")
	}
	return productModel.ID, nil
}

func New(db *gorm.DB) product.ProductDataInterface {
	return ProductData{
		db: db,
	}
}

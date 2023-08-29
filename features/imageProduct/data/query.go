package data

import (
	"errors"
	imageproduct "warung_online/features/imageProduct"
	"warung_online/features/structsEntity"

	"gorm.io/gorm"
)

type ImageProductData struct {
	db *gorm.DB
}

// Delete implements imageproduct.ImageProductDataInterface.
func (repo *ImageProductData) Delete(idImage uint) error {
	var image structsEntity.ImageProduct
	tx:=repo.db.Delete(&image,idImage)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("not row affected")
	}
	return nil
}

// Update implements imageproduct.ImageProductDataInterface.
func (repo *ImageProductData) Update(input structsEntity.ImageProductEntity, idImage uint) (uint, error) {
	inputModel := structsEntity.ImageProductEntityToModel(input)
	tx := repo.db.Model(&structsEntity.ImageProduct{}).Where("id=? and product_id=?", idImage, input.ProductID).Updates(inputModel)
	if tx.Error != nil {
		return 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("not row affected")
	}
	return idImage, nil
}

// Insert implements imageproduct.ImageProductDataInterface.
func (repo *ImageProductData) Insert(input structsEntity.ImageProductEntity) (uint, error) {
	inputModel := structsEntity.ImageProductEntityToModel(input)
	tx := repo.db.Create(&inputModel)
	if tx.Error != nil {
		return 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("not row affected")
	}
	return inputModel.ID, nil
}

// SelectById implements imageproduct.ImageProductDataInterface.
func (repo *ImageProductData) SelectById(idImage uint) (structsEntity.ImageProductEntity, error) {
	var imageModel structsEntity.ImageProduct
	tx := repo.db.First(&imageModel, idImage)
	if tx.Error != nil {
		return structsEntity.ImageProductEntity{}, tx.Error
	}
	output := structsEntity.ImageProductModelToEntity(imageModel)
	return output, nil
}

// SelectUser implements imageproduct.ImageProductDataInterface.
func (repo *ImageProductData) SelectUser(idUser uint) (structsEntity.UserEntity, error) {
	var user structsEntity.User
	tx := repo.db.First(&user, idUser)
	if tx.Error != nil {
		return structsEntity.UserEntity{}, tx.Error
	}
	output := structsEntity.UserModelToEntity(user)
	return output, nil
}

func New(db *gorm.DB) imageproduct.ImageProductDataInterface {
	return &ImageProductData{
		db: db,
	}
}

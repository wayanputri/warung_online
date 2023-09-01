package data

import (
	"errors"
	imagereview "warung_online/features/imageReview"
	"warung_online/features/structsEntity"

	"gorm.io/gorm"
)

type ImageReviewData struct {
	db *gorm.DB
}

// Delete implements imagereview.ImageReviewDataInterface.
func (repo *ImageReviewData) Delete(id uint) error {
	var image structsEntity.ImageReview
	tx:=repo.db.Delete(&image,id)
	if tx.Error != nil {
		return errors.New("failed delete data")
	}
	if tx.RowsAffected == 0 {
		return errors.New("row not affected")
	}
	return nil
}

// Insert implements imagereview.ImageReviewDataInterface.
func (repo *ImageReviewData) Insert(input structsEntity.ImageReviewEntity) (uint, error) {
	var review structsEntity.Review
	txx := repo.db.First(&review, input.ReviewID)
	if txx.Error != nil {
		return 0, errors.New("failed get review data")
	}
	inputModel := structsEntity.ImageReviewEntityToModel(input)
	tx := repo.db.Create(&inputModel)
	if tx.Error != nil {
		return 0, errors.New("failed create data")
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("row not affected")
	}
	return inputModel.ID, nil
}

// SelectById implements imagereview.ImageReviewDataInterface.
func (repo *ImageReviewData) SelectById(id uint) (structsEntity.ImageReviewEntity, error) {
	var image structsEntity.ImageReview
	tx := repo.db.First(&image, id)
	if tx.Error != nil {
		return structsEntity.ImageReviewEntity{}, errors.New("failed read data")
	}
	output := structsEntity.ImageReviewModelToEntity(image)
	return output, nil
}

func New(db *gorm.DB) imagereview.ImageReviewDataInterface {
	return &ImageReviewData{
		db: db,
	}
}

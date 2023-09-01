package service

import (
	imagereview "warung_online/features/imageReview"
	"warung_online/features/structsEntity"
)

type ImageReviewService struct {
	imageReviewService imagereview.ImageReviewDataInterface
}

// Delete implements imagereview.ImageReviewServiceInterface.
func (service *ImageReviewService) Delete(id uint) error {
	err:=service.imageReviewService.Delete(id)
	if err != nil{
		return err
	}
	return nil
}

// Add implements imagereview.ImageReviewServiceInterface.
func (service *ImageReviewService) Add(input structsEntity.ImageReviewEntity) (structsEntity.ImageReviewEntity, error) {
	id, errAdd := service.imageReviewService.Insert(input)
	if errAdd != nil {
		return structsEntity.ImageReviewEntity{}, errAdd
	}
	data, errGet := service.imageReviewService.SelectById(id)
	if errGet != nil {
		return structsEntity.ImageReviewEntity{}, errGet
	}
	return data, nil
}

func New(service imagereview.ImageReviewDataInterface) imagereview.ImageReviewServiceInterface {
	return &ImageReviewService{
		imageReviewService: service,
	}
}

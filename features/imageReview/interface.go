package imagereview

import "warung_online/features/structsEntity"

type ImageReviewDataInterface interface {
	Insert(input structsEntity.ImageReviewEntity)(uint,error)
	SelectById(id uint)(structsEntity.ImageReviewEntity,error)
	Delete(id uint)error
}

type ImageReviewServiceInterface interface{
	Add(input structsEntity.ImageReviewEntity)(structsEntity.ImageReviewEntity,error)
	Delete(id uint)error
}
package imageproduct

import "warung_online/features/structsEntity"

type ImageProductDataInterface interface {
	Insert(input structsEntity.ImageProductEntity)(uint,error)
	SelectById(idImage uint)(structsEntity.ImageProductEntity,error)
	SelectUser(idUser uint)(structsEntity.UserEntity,error)
	Update(input structsEntity.ImageProductEntity,idImage uint)(uint,error)
	Delete(idImage uint)error
}

type ImageProductServiceInterface interface{
	Add(input structsEntity.ImageProductEntity,idUser uint)(structsEntity.ImageProductEntity,error)
	Edit(input structsEntity.ImageProductEntity,idImage uint,idUser uint)(structsEntity.ImageProductEntity,error)
	Delete(idImage uint)error
}
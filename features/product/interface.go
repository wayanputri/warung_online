package product

import (
	"warung_online/features/structsEntity"
)

type ProductDataInterface interface {
	Insert(input structsEntity.ProductEntity)(uint,error)
	SelectUser(id uint)(structsEntity.UserEntity,error)
	Update(input structsEntity.ProductEntity,id uint)(uint,error)
	Delete(idProduct uint)error
	SelectAll()([]structsEntity.ProductEntity,error)
	SelectById(id uint)(structsEntity.ProductEntity,error)
	Searching(name string)([]structsEntity.ProductEntity,error)
}

type ProductServiceInterface interface{
	Add(input structsEntity.ProductEntity)error
	Edit(input structsEntity.ProductEntity,id uint)(error)
	Delete(idUser uint,idProduct uint)error
	GetAll(name string)([]structsEntity.ProductEntity,error)
	SelectById(id uint)(structsEntity.ProductEntity,error)
}
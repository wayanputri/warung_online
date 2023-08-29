package user

import (
	"warung_online/features/structsEntity"
)

type UserDataInterface interface {
	Insert(user structsEntity.UserEntity)(uint,error)
	Login(user structsEntity.UserEntity)(uint,string,error)
	SelectAll()([]structsEntity.UserEntity,error)
	SelectById(id uint)(structsEntity.UserEntity,error)
	Update(input structsEntity.UserEntity,id uint)(error)
	Delete(id uint)error
	Upgrade(input structsEntity.UserEntity,id uint)error
	UpdateProfil(input structsEntity.UserEntity,id uint)error
}
type UserServiceInterface interface{
	Add(user structsEntity.UserEntity)(uint,error)
	Login(user structsEntity.UserEntity)(uint,error)
	GetAll()([]structsEntity.UserEntity,error)
	SelectById(id uint)(structsEntity.UserEntity,error)
	Edit(input structsEntity.UserEntity,id uint)error
	Delete(id uint)error
	Upgrade(input structsEntity.UserEntity,id uint)error
	UpdateProfil(input structsEntity.UserEntity,id uint)error
}
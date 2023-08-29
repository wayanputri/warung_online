package transaction

import "warung_online/features/structsEntity"

type TransactionDataInterface interface {
	Insert(input structsEntity.TransactionEntity)(uint,error)
	SelectByProduct(idProduct uint)(structsEntity.ProductEntity,error)
	SelectById(idTransaction uint,IdUser uint)(structsEntity.TransactionEntity,error)
	SelectAll(UserId uint)([]structsEntity.TransactionEntity,error)
	Update(input structsEntity.TransactionEntity)(uint,error)
	Delete(id uint,idUser uint)error
}

type TransactionServiceInterface interface {
	Add(input structsEntity.TransactionEntity)(structsEntity.TransactionEntity,error)
	GetAll(UserId uint)([]structsEntity.TransactionEntity,error)
	GetById(idTransaction uint,IdUser uint)(structsEntity.TransactionEntity,error)
	Edit(input structsEntity.TransactionEntity)(structsEntity.TransactionEntity,error)
	Delete(id uint,idUser uint)error
}
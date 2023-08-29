package transactiondetil

import "warung_online/features/structsEntity"

type TransactionDetilDataInterface interface {
	Insert(inputKeranjang []structsEntity.TransactionKeranjangEntity, input structsEntity.TransactionFinalEntity)(uint,error)
	SelectAllTansaction(idUser uint)([]uint,error)
	SelectAllTansactionKeranjang(idTransaction []uint)([]uint ,error)
	SelectIdTansactionKeranjang(idKeranjang []uint,idKer []uint)(int,[]structsEntity.TransactionKeranjangEntity ,error)
	SelectById(id uint)(structsEntity.TransactionFinalEntity,error)
}

type TransactionDetilServiceInterface interface{
	Add(idKeranjang []uint,idUser uint,input structsEntity.TransactionFinalEntity)(structsEntity.TransactionFinalEntity,error)
}
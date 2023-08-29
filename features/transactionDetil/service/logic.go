package service

import (
	"errors"
	"fmt"
	"warung_online/features/structsEntity"
	transactiondetil "warung_online/features/transactionDetil"
	"warung_online/helper"
)

type TransactionDetilService struct {
	transactionDetilService transactiondetil.TransactionDetilDataInterface
}

// Add implements transactiondetil.TransactionDetilServiceInterface.
func (service *TransactionDetilService) Add(idKeranjang []uint,idUser uint, input structsEntity.TransactionFinalEntity) (structsEntity.TransactionFinalEntity, error) {
	
	idTransaction,errTransaction:=service.transactionDetilService.SelectAllTansaction(idUser)
	if errTransaction != nil{
		return structsEntity.TransactionFinalEntity{},errors.New("tidak dapat menambahkan transaction user lain")
	}
	idKer,errKeranjang:=service.transactionDetilService.SelectAllTansactionKeranjang(idTransaction)
	if errKeranjang != nil{
		return structsEntity.TransactionFinalEntity{},errors.New("data transaksi tidak ditemukan")
	}
	harga,dataKeranjang,erridKeranjang:=service.transactionDetilService.SelectIdTansactionKeranjang(idKeranjang,idKer)
	if erridKeranjang != nil{
		return structsEntity.TransactionFinalEntity{},errors.New("data keranjang tidak ditemukan")
	}
	orderId,errOrder:=helper.GenerateUUID()
	if errOrder != nil{
		return structsEntity.TransactionFinalEntity{},errors.New("failed generate orderId")
	}
	input.OrderID = orderId
	input.TotalHarga = harga
	idTranDet,err:=service.transactionDetilService.Insert(dataKeranjang,input)
	if err != nil{
		return structsEntity.TransactionFinalEntity{},err
	}
	fmt.Println("id",idTranDet)
	data,errGet:=service.transactionDetilService.SelectById(idTranDet)
	if errGet != nil{
		return structsEntity.TransactionFinalEntity{},errors.New("data transaction detail tidak ditemukan")
	}
	return data,nil	
}

func New(transaction transactiondetil.TransactionDetilDataInterface) transactiondetil.TransactionDetilServiceInterface {
	return &TransactionDetilService{
		transactionDetilService: transaction,
	}
}

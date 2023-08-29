package service

import (
	"errors"
	"fmt"
	"warung_online/features/structsEntity"
	"warung_online/features/transaction"
)

type TransactionService struct {
	transactionService transaction.TransactionDataInterface
}

// Delete implements transaction.TransactionServiceInterface.
func (service *TransactionService) Delete(id uint, idUser uint) error {
	err:=service.transactionService.Delete(id,idUser)
	if err != nil {
		return err
	}
	return nil
}

// Edit implements transaction.TransactionServiceInterface.
func (service *TransactionService) Edit(input structsEntity.TransactionEntity) (structsEntity.TransactionEntity, error) {
	data, errGet := service.transactionService.SelectByProduct(input.ProductID)
	if errGet != nil {
		return structsEntity.TransactionEntity{}, errors.New("data product tidak ditemukan")
	}
	if data.Stok < input.Jumlah {
		return structsEntity.TransactionEntity{}, errors.New("stok tidak cukup, harap periksa stok terupdate")
	}
	fmt.Println("id user:", input.Id)
	totalHarga := input.Jumlah * data.Harga
	input.TotalHarga = totalHarga

	id, err := service.transactionService.Update(input)
	if err != nil {
		return structsEntity.TransactionEntity{}, err
	}
	dataTransaction, errData := service.transactionService.SelectById(id, input.UserID)
	if errData != nil {
		return structsEntity.TransactionEntity{}, errData
	}
	return dataTransaction, nil
}

// GetById implements transaction.TransactionServiceInterface.
func (service *TransactionService) GetById(idTransaction uint, IdUser uint) (structsEntity.TransactionEntity, error) {
	data, err := service.transactionService.SelectById(idTransaction, IdUser)
	return data, err
}

// GetAll implements transaction.TransactionServiceInterface.
func (service *TransactionService) GetAll(UserId uint) ([]structsEntity.TransactionEntity, error) {
	data, err := service.transactionService.SelectAll(UserId)
	if err != nil {
		return nil, err
	}
	return data, err
}

// Add implements transaction.TransactionServiceInterface.
func (service *TransactionService) Add(input structsEntity.TransactionEntity) (structsEntity.TransactionEntity, error) {
	data, errGet := service.transactionService.SelectByProduct(input.ProductID)
	if errGet != nil {
		return structsEntity.TransactionEntity{}, errors.New("data product tidak ditemukan")
	}
	if data.Stok < input.Jumlah {
		return structsEntity.TransactionEntity{}, errors.New("stok tidak cukup, harap periksa stok terupdate")
	}
	totalHarga := input.Jumlah * data.Harga
	input.TotalHarga = totalHarga
	id, err := service.transactionService.Insert(input)
	if err != nil {
		return structsEntity.TransactionEntity{}, err
	}
	dataTransaction, errData := service.transactionService.SelectById(id, input.UserID)
	if errData != nil {
		return structsEntity.TransactionEntity{}, errData
	}
	return dataTransaction, nil
}

func New(service transaction.TransactionDataInterface) transaction.TransactionServiceInterface {
	return &TransactionService{
		transactionService: service,
	}
}

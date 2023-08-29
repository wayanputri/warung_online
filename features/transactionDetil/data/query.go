package data

import (
	"errors"
	"warung_online/features/structsEntity"
	transactiondetil "warung_online/features/transactionDetil"

	"gorm.io/gorm"
)

type TransactionDetilData struct {
	db *gorm.DB
}

// Insert implements transactiondetil.TransactionDetilDataInterface.
func (repo *TransactionDetilData) Insert(inputKeranjang []structsEntity.TransactionKeranjangEntity, input structsEntity.TransactionFinalEntity) (uint, error) {
	inputTransaction:=structsEntity.TransactionFinalEntityToModel(input)
	tx:=repo.db.Create(&inputTransaction)
	if tx.Error != nil{
		return 0,errors.New("failed error create transaction")
	}
	if tx.RowsAffected ==0{
		return 0,errors.New("row not affected")
	}
	for _,value:= range inputKeranjang{
		updateKeranjang:=structsEntity.TransactionKeranjangEntityToModel(value)
		updateKeranjang.TransactionFinalID = inputTransaction.ID
		txx:=repo.db.Model(&structsEntity.TransactionKeranjang{}).Where("id=?",value.Id).Updates(updateKeranjang)
		if txx.Error != nil{
			return 0,errors.New("failed error update keranjang")
		}
		if txx.RowsAffected ==0{
			return 0,errors.New("row not affected")
		}
	}
	return inputTransaction.ID,nil
}

// SelectAllTansaction implements transactiondetil.TransactionDetilDataInterface.
func (repo *TransactionDetilData) SelectAllTansaction(idUser uint) ([]uint, error) {
	var transaction []structsEntity.Transaction
	tx:=repo.db.Where("user_id=?",idUser).Find(&transaction)
	if tx.Error != nil{
		return nil,errors.New("failed error get transaction")
	}
	var transactionId []uint
	for _,value:=range transaction{
		transactionId = append(transactionId, value.ID)
	}
	return transactionId,nil
}

// SelectAllTansactionKeranjang implements transactiondetil.TransactionDetilDataInterface.
func (repo *TransactionDetilData) SelectAllTansactionKeranjang(idTransaction []uint) ([]uint, error) {
	var keranjang []structsEntity.TransactionKeranjang
	tx:=repo.db.Where("transaction_id IN ?",idTransaction).Find(&keranjang)
	if tx.Error != nil{
		return nil,errors.New("failed error get transaction keranjang")
	}
	var transactionId []uint
	for _,value:=range keranjang{
		transactionId = append(transactionId, value.ID)
	}
	return transactionId,nil
}

// SelectById implements transactiondetil.TransactionDetilDataInterface.
func (repo *TransactionDetilData) SelectById(id uint) (structsEntity.TransactionFinalEntity, error) {
	var transaction structsEntity.TransactionFinal
	tx:=repo.db.Preload("TransactionKeranjang").Preload("TransactionKeranjang.Transaction").Preload("TransactionKeranjang.Transaction.Users").Preload("TransactionKeranjang.Transaction.Products").Preload("TransactionKeranjang.Transaction.Products.Users").Preload("Payment").First(&transaction,id)
	if tx.Error != nil{
		return structsEntity.TransactionFinalEntity{},errors.New("failed error get transaction detil")
	}
	output:=structsEntity.TransactionFinalModelToEntity(transaction)
	return output,nil
}

// SelectIdTansactionKeranjang implements transactiondetil.TransactionDetilDataInterface.
func (repo *TransactionDetilData) SelectIdTansactionKeranjang(idKeranjang []uint, idKer []uint) (int,[]structsEntity.TransactionKeranjangEntity, error) {
	var keranjangmodel []structsEntity.TransactionKeranjang
	for _,value:=range idKer{
		for _,val:=range idKeranjang{
			if value==val{
				var keranjang structsEntity.TransactionKeranjang
				tx:=repo.db.Where("id = ?",val).First(&keranjang)
				if tx.Error != nil{
					return 0,nil,errors.New("failed error get transaction keranjang model")
				}
				keranjangmodel = append(keranjangmodel, keranjang)
			}
		}
	}
	var harga int=0
	
	for _,value1:=range keranjangmodel{
		var transaction structsEntity.Transaction
		txxx:=repo.db.Where("id=?",value1.TransactionID).First(&transaction)
		if txxx.Error != nil{
			return 0,nil,errors.New("failed error get transaction")
		}
		harga +=transaction.TotalHarga
	}
	
	var output []structsEntity.TransactionKeranjangEntity
	for _,value:=range keranjangmodel{
		output = append(output, structsEntity.TransactionKeranjangModelToEntity(value))
	}
	
	return harga,output,nil
}

func New(db *gorm.DB) transactiondetil.TransactionDetilDataInterface {
	return &TransactionDetilData{
		db: db,
	}
}

package service

import (
	"errors"
	"fmt"
	"warung_online/features/product"
	"warung_online/features/structsEntity"
)

type ProductService struct {
	productService product.ProductDataInterface
}

// SelectById implements product.ProductServiceInterface.
func (service ProductService) SelectById(id uint) (structsEntity.ProductEntity, error) {
	data,err:=service.productService.SelectById(id)
	if err != nil {
		return structsEntity.ProductEntity{}, err
	}
	return data, nil
}

// GetAll implements product.ProductServiceInterface.
func (service ProductService) GetAll(name string) ([]structsEntity.ProductEntity, error) {

	if name ==""{
		data, err := service.productService.SelectAll()
		return data, err
	}else{
		data, err := service.productService.Searching(name)
		return data, err
	}	
}

// Delete implements product.ProductServiceInterface.
func (service ProductService) Delete(idUser uint, idProduct uint) error {
	data, errUser := service.productService.SelectUser(idUser)
	if errUser != nil {
		return errUser
	}
	fmt.Println("id_user:", idUser)
	if data.Role != "pedagang" {
		return errors.New("hanya pedagang yang bisa menghapus data product")
	}
	err := service.productService.Delete(idProduct)
	if err != nil {
		return err
	}
	return nil
}

// Edit implements product.ProductServiceInterface.
func (service ProductService) Edit(input structsEntity.ProductEntity, id uint) error {
	data, errUser := service.productService.SelectUser(input.UserID)
	if errUser != nil {
		return errUser
	}
	if data.Role != "pedagang" {
		return errors.New("hanya pedagang yang bisa mengedit data product")
	}
	_, err := service.productService.Update(input, id)
	if err != nil {
		return err
	}

	return nil
}

// Add implements product.ProductServiceInterface.
func (service ProductService) Add(input structsEntity.ProductEntity) error {
	if input.Nama == "" || input.Harga == 0 || input.Stok == 0 {
		return errors.New("error validation. nama/harga/stok required")
	}
	data, errUser := service.productService.SelectUser(input.UserID)
	if errUser != nil {
		return errUser
	}
	if data.Role != "pedagang" {
		return errors.New("hanya pedagang yang bisa menambah data product")
	}
	_, err := service.productService.Insert(input)
	if err != nil {
		return err
	}

	return nil
}

func New(service product.ProductDataInterface) product.ProductServiceInterface {
	return ProductService{
		productService: service,
	}
}

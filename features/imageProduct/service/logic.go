package service

import (
	"errors"
	imageproduct "warung_online/features/imageProduct"
	"warung_online/features/structsEntity"
)

type ImageProductService struct {
	imageProductService imageproduct.ImageProductDataInterface
}

// Delete implements imageproduct.ImageProductServiceInterface.
func (service *ImageProductService) Delete(idImage uint) error {
	err:=service.imageProductService.Delete(idImage)
	if err != nil {
		return err
	}
	return nil
}

// Edit implements imageproduct.ImageProductServiceInterface.
func (service *ImageProductService) Edit(input structsEntity.ImageProductEntity, idImage uint, idUser uint) (structsEntity.ImageProductEntity, error) {
	data, err := service.imageProductService.SelectUser(idUser)
	if err != nil {
		return structsEntity.ImageProductEntity{}, err
	}
	if data.Role != "pedagang" {
		return structsEntity.ImageProductEntity{}, errors.New("access not validate,hanya pedagang yang bisa insert image")
	}
	id, err := service.imageProductService.Update(input, idImage)
	if err != nil {
		return structsEntity.ImageProductEntity{}, err
	}
	dataGet, errGet := service.imageProductService.SelectById(id)
	if errGet != nil {
		return structsEntity.ImageProductEntity{}, errGet
	}
	return dataGet, nil
}

// Add implements imageproduct.ImageProductServiceInterface.
func (service *ImageProductService) Add(input structsEntity.ImageProductEntity, idUser uint) (structsEntity.ImageProductEntity, error) {
	data, err := service.imageProductService.SelectUser(idUser)
	if err != nil {
		return structsEntity.ImageProductEntity{}, err
	}
	if data.Role != "pedagang" {
		return structsEntity.ImageProductEntity{}, errors.New("access not validate,hanya pedagang yang bisa insert image")
	}
	id, errInsert := service.imageProductService.Insert(input)
	if errInsert != nil {
		return structsEntity.ImageProductEntity{}, errInsert
	}
	dataGet, errGet := service.imageProductService.SelectById(id)
	if errGet != nil {
		return structsEntity.ImageProductEntity{}, errGet
	}
	return dataGet, nil
}

func New(service imageproduct.ImageProductDataInterface) imageproduct.ImageProductServiceInterface {
	return &ImageProductService{
		imageProductService: service,
	}
}

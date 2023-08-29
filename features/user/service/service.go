package service

import (
	"errors"
	"warung_online/features/structsEntity"
	"warung_online/features/user"
	"warung_online/helper"

	"github.com/go-playground/validator"
)

type UserService struct {
	userData user.UserDataInterface
	validate *validator.Validate
}

// UpdateProfil implements user.UserServiceInterface.
func (service *UserService) UpdateProfil(input structsEntity.UserEntity, id uint) error {
	
	err:=service.userData.UpdateProfil(input,id)
	if err != nil{
		return err
	}
	return nil
}

// Upgrade implements user.UserServiceInterface.
func (service *UserService) Upgrade(input structsEntity.UserEntity, id uint) error {
	if input.DataUpdate == "" || input.Role != "pedagang" {
		return errors.New("link error")
	}
	err := service.userData.Upgrade(input, id)
	if err != nil {
		return err
	}
	return nil
}

// Delete implements user.UserServiceInterface.
func (service *UserService) Delete(id uint) error {
	err := service.userData.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

// Edit implements user.UserServiceInterface.
func (service *UserService) Edit(input structsEntity.UserEntity, id uint) error {
	err := service.userData.Update(input, id)
	if err != nil {
		return err
	}
	return nil
}

// SelectById implements user.UserServiceInterface.
func (service *UserService) SelectById(id uint) (structsEntity.UserEntity, error) {
	data, err := service.userData.SelectById(id)
	if err != nil {
		return structsEntity.UserEntity{}, err
	}
	return data, nil
}

// GetAll implements user.UserServiceInterface.
func (service *UserService) GetAll() ([]structsEntity.UserEntity, error) {
	data, err := service.userData.SelectAll()
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Login implements user.UserServiceInterface.
func (service *UserService) Login(user structsEntity.UserEntity) (uint, error) {
	if errValidate := service.validate.Struct(user); errValidate != nil {
		return 0, errors.New("email not valid")
	}
	id, pw, err := service.userData.Login(user)
	if err != nil {
		return 0, err
	}
	match:=helper.CheckPassword(user.Password,pw)
	if !match{
		return 0,errors.New("password salah")
	}

	return id, nil
}

// Add implements user.UserServiceInterface.
func (service *UserService) Add(user structsEntity.UserEntity) (uint, error) {
	if errValidate := service.validate.Struct(user); errValidate != nil {
		return 0, errValidate
	}

	pass,errPas:=helper.HasPassword(user.Password)
	if errPas != nil{
		return 0,errors.New("error hash password")
	}
	user.Password = pass
	id, err := service.userData.Insert(user)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func New(userData user.UserDataInterface) user.UserServiceInterface {
	return &UserService{
		userData: userData,
		validate: validator.New(),
	}
}

package data

import (
	"errors"
	"warung_online/features/structsEntity"
	"warung_online/features/user"

	"gorm.io/gorm"
)

type UserData struct {
	db *gorm.DB
}

// UpdateProfil implements user.UserDataInterface.
func (data *UserData) UpdateProfil(input structsEntity.UserEntity, id uint) error {
	tx := data.db.Model(&structsEntity.User{}).Where("id=?", id).Updates(structsEntity.UserEntityToModel(input))
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("failed edit profil data user")
	}
	return nil
}

// Upgrade implements user.UserDataInterface.
func (data *UserData) Upgrade(input structsEntity.UserEntity, id uint) error {

	tx := data.db.Model(&structsEntity.User{}).Where("id=?", id).Updates(structsEntity.UserEntityToModel(input))
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("failed upgrade data user")
	}
	return nil
}

// Delete implements user.UserDataInterface.
func (data *UserData) Delete(id uint) error {
	var user structsEntity.User
	tx := data.db.Delete(&user, id)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("failed delete data user")
	}
	return nil
}

// Update implements user.UserDataInterface.
func (data *UserData) Update(input structsEntity.UserEntity, id uint) error {
	tx := data.db.Model(&structsEntity.User{}).Where("id=?", id).Updates(structsEntity.UserEntityToModel(input))
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("failed update data user")
	}
	return nil
}

// SelectById implements user.UserDataInterface.
func (data *UserData) SelectById(id uint) (structsEntity.UserEntity, error) {
	var userModel structsEntity.User
	tx := data.db.First(&userModel, id)
	if tx.Error != nil {
		return structsEntity.UserEntity{}, tx.Error
	}
	userEntity := structsEntity.UserModelToEntity(userModel)
	return userEntity, nil
}

// SelectAll implements user.UserDataInterface.
func (data *UserData) SelectAll() ([]structsEntity.UserEntity, error) {
	var user []structsEntity.User
	tx := data.db.Find(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var users []structsEntity.UserEntity
	for _, value := range user {
		users = append(users, structsEntity.UserModelToEntity(value))
	}
	return users, nil
}

// Login implements user.UserDataInterface.
func (data *UserData) Login(user structsEntity.UserEntity) (uint, string, error) {
	var userModal structsEntity.User
	tx := data.db.Where("email=?", user.Email).First(&userModal)
	if tx.Error != nil {
		return 0, "", errors.New("email not found")
	}
	userEntity := structsEntity.UserModelToEntity(userModal)
	return userEntity.Id, userEntity.Password, nil
}

// Insert implements user.UserDataInterface.
func (data *UserData) Insert(user structsEntity.UserEntity) (uint, error) {

	userData := structsEntity.UserEntityToModel(user)
	tx := data.db.Create(&userData)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return userData.ID, nil
}

func New(db *gorm.DB) user.UserDataInterface {
	return &UserData{
		db: db,
	}
}

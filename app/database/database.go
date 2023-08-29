package database

import (
	"fmt"
	"warung_online/app/config"
	"warung_online/features/structsEntity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMysql(cfg *config.AppConfig) *gorm.DB{
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.DBUsername,cfg.DBPassword,cfg.DBHost,cfg.DBPort,cfg.DBName)

		DB, err := gorm.Open(mysql.Open(connectionString),&gorm.Config{})
		if err != nil{
			panic(err)
		}
		return DB
}

func InittialMigration(db *gorm.DB){
	db.AutoMigrate(&structsEntity.User{},&structsEntity.Product{},&structsEntity.ImageProduct{},&structsEntity.Transaction{},&structsEntity.TransactionFinal{},&structsEntity.Payment{},&structsEntity.Review{},&structsEntity.ImageReview{},&structsEntity.TransactionKeranjang{})
}


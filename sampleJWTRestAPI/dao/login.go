package dao

import (
	"sampleJWTRestAPI/models"
	"time"

	"github.com/jinzhu/gorm"
)

func CreateLoginUser(db *gorm.DB, data *models.Login) error {
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()
	db.Table("login").NewRecord(data)
	if err := db.Table("login").Create(&data).Error; err != nil {
		return err
	}
	return nil
}

func GetLoginUserByIDPassword(db *gorm.DB, id string, password string) *models.Login {
	login := models.NewLogin()
	if db.Table("login").Where("ID = ? AND Password = ?", id, password).First(&login).RecordNotFound() {
		return nil
	}
	return login
}

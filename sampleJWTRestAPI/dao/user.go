package dao

import (
	"sampleJWTRestAPI/models"
	"time"

	"github.com/jinzhu/gorm"
)

func CreateUser(db *gorm.DB, data *models.User) error {
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()
	db.Table("user").NewRecord(data)
	if err := db.Table("user").Create(&data).Error; err != nil {
		return err
	}
	return nil
}

func GetUserByID(db *gorm.DB, id string) *models.User {
	user := models.NewUser()
	if db.Table("user").Where("ID = ?", id).First(&user).RecordNotFound() {
		return nil
	}
	return user
}

func UpdateUserByID(db *gorm.DB, data *models.User, id string) error {
	data.UpdatedAt = time.Now()
	user := models.NewUser()
	if err := db.Table("user").Where("ID = ?", id).First(&user).Omit("created_at").Updates(&data).Error; err != nil {
		return err
	}
	return nil
}

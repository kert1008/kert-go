package services

import (
	"sampleJWTRestAPI/dao"
	"sampleJWTRestAPI/db"
	"sampleJWTRestAPI/models"
)

func GetUserByID(id string) *models.User {
	db := db.GormConnect()
	defer db.Close()

	return dao.GetUserByID(db, id)
}

func UpdateUser(user *models.User) error {
	db := db.GormConnect()
	defer db.Close()

	if err := dao.UpdateUserByID(db, user, user.ID); err != nil {
		return err
	}
	return nil
}

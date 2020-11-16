package services

import (
	"crypto/md5"
	"encoding/hex"
	"sampleJWTRestAPI/dao"
	"sampleJWTRestAPI/db"
	"sampleJWTRestAPI/models"
)

func CreateLoginUser(login *models.Login) error {
	db := db.GormConnect()
	defer db.Close()

	tx := db.Begin()
	var err error

	login.Password = md5Encoding(login.Password)

	if err = dao.CreateLoginUser(tx, login); err != nil {
		tx.Rollback()
		return err
	}

	user := models.User{}
	user.ID = login.ID

	if err = dao.CreateUser(tx, &user); err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func GetLoginUserByIDPassword(id string, password string) *models.Login {
	db := db.GormConnect()
	defer db.Close()

	return dao.GetLoginUserByIDPassword(db, id, md5Encoding(password))
}

func md5Encoding(password string) string {
	md5 := md5.Sum([]byte(password))
	md5Password := hex.EncodeToString(md5[:])
	return md5Password
}

package models

import (

	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
    ID int `gorm:"primary_key" json:"id"`
    Username string `json:"username"`
    Password string `json:"password"`
}

type GetAuthForm struct {
	Username string `json:"username" binding:"required,gte=0,lte=50"`
    Password string `json:"password" binding:"required,gte=0,lte=50"`
}

func CheckAuth(username, password string) bool {
    var auth Auth
    db.Where("username = ?", username).First(&auth)

    if auth.ID > 0 && comparePasswords(auth.Password, password) {
        return true
    }

    return false
}

func hashPassword(password string) string {
    bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
    return string(bytes)
}

func comparePasswords(hash string, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(plainPassword))
    return err == nil
}
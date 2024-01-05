package src

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	Login    string
	Password string
}

func NewUserModel(login string, password string) *UserModel {
	return &UserModel{Login: login, Password: password}
}

func (m *UserModel) DBSave() {
	db, err := gorm.Open(sqlite.Open("Users.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.Create(m)
}

func (m *UserModel) DBUpdate() {
	db, err := gorm.Open(sqlite.Open("Users.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.Model(m).Updates(m)
}

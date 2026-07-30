package handlers

import (
	"github.com/glebarez/sqlite"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"size:255"`
	Currency int
}

func OpenDB() {
	db, err := gorm.Open(sqlite.Open("db.db"), &gorm.Config{})
	Logger("couldnt open db", logrus.FatalLevel, err)

	err = db.AutoMigrate(&User{})
	Logger("couldnt migrate db", logrus.FatalLevel, err)

	//	user := User{ID: 2, Name: "kıyma", Currency: 100}
	//	result := db.Create(&user)
	//	Logger("couldn't create user", logrus.ErrorLevel, result.Error)

	//	var findUser User
	//	db.Find(&findUser) // or db.First(&findUser, 1)

	//Info(fmt.Sprintf("%+v", findUser)) // Convert User to string first
}

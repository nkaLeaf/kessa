package handlers

import (
	"github.com/glebarez/sqlite"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey"` // Use uint for auto-incrementing IDs
	Name     string `gorm:"size:255"`
	Currency int
}

func Db(action string, context string, count int) {
	db, err := gorm.Open(sqlite.Open("db.db"), &gorm.Config{})
	Logger("couldnt open db", logrus.FatalLevel, err)

	err = db.AutoMigrate(&User{})
	Logger("couldnt migrate db", logrus.FatalLevel, err)

	switch action {
	case "create":
		user := User{ID: 1, Name: context, Currency: count}
		result := db.Create(&user)
		Logger("couldnt create user", logrus.ErrorLevel, result.Error)
	}
	//half done db TODO
}

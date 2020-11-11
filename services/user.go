package services

import (
	"fmt"
	"net/http"

	"github.com/tian372/serverTest/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	DB *gorm.DB
}

func NewUser(psqlInfo string) *User {
	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		fmt.Println("Cannot connect to database! Error!")
		fmt.Println(err)
		//TODO: Handle Error
	}
	user := User{
		DB: db,
	}

	return &user
}

//Create User Info

func (user *User) Create(w http.ResponseWriter, r *http.Request) {
	testData := models.Exercise{}
	user.DB.First(&testData)
	fmt.Fprintln(w, testData.Activities)
}

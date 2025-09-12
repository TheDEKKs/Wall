package database

import (
	"log"
	"os"
	"time"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


//Структура базы данных юзера
type User struct {
	// Данные акаунта 
	Id_User int `gorm:"unique;primaryKey"`
	Name_User string `gorm:"unique;"`
	Id_Telegram string  `gorm:"unique;"`
	Id_Wall int `gorm:"unique"`
}

// Стурктура базы данных стены
type Wall struct {
	// Настройка стены
	Id_Wall int `gorm:"unique;primaryKey"`
	Id_Creator string  `gorm:"unique;"`
	Id_Comment []int `gorm:"column:Id_Comment;type:int[]"`

	// Настройки для коментариев на стене
	Mat bool 
	Anonymously bool 
	
}

// Структура базы данных коментариев
type Comment struct {
	// Данные 
	Id_Comment int `gorm:"unique;primaryKey"`
	Id_Wall int `gorm:"unique;"`
	Id_Commentator int `gorm:"unique;"`

	Text_Comment string `gorm:"size:128"`
	CreatedAt time.Time
	UpdatedAt time.Time
}


var db *gorm.DB

func InitDB() {
	log.SetOutput(os.Stderr)
    	log.SetFlags(log.LstdFlags | log.Lshortfile)
//	dsn := "host=postgres user=postgres password=2242 dbname=mydb port=5432 sslmode=disable"
//	dsn := "postgres://localhost:2242@postgres:5432/mydb"
	dsn := os.Getenv("DATABASE_URL")
	var err error

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panicln("Error DB Conect: %v", err)
	}


	if err := db.AutoMigrate(&User{}, &Wall{}, &Comment{}); err != nil {
		fmt.Println("Error Migration %v", err)
	}



	//CreateBase()

}

func CreateBase() {
	if err := db.AutoMigrate(&User{}); err != nil {
		fmt.Println("Error Migration %v", err)
	}

	if err := db.AutoMigrate(&Wall{}); err != nil {
		fmt.Println("Error Migration Wall %v", err)
	}


	if err := db.AutoMigrate(&Comment{}); err != nil {
		fmt.Println("Error Migration Wall %v", err)
	}

}

func Add() {
	user := User{1112, "tessdt", "tedwst", 2111}
	if err := db.Create(&user).Error; err != nil {
		fmt.Println("Error add user %v", err)
	}
}
/*
func Search_User(user_search string) string {
	user := UsersWaka{User:user_search} 
	res_search := db.Find(&UsersWaka{}, "User = ?", user)
	if errors.Is(res_search.Error, gorm.ErrRecordNotFound) {
		if err := db.Create(&user).Error; err != nil {
			fmt.Println(err)
			return "Error Create"
		}

		return "Good Create New User"
	} else {
		fmt.Println("User in database")
		return "User in database"
	}

////	return "Good"

}
*/



package database

import (
	"errors"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


type UsersWaka struct {
	gorm.Model
	User string `gorm:"unique"`	
	Time int 
	LangueList []string `gorm:"type:text[]"`
	EditorList []string `gorm:"type:text[]"`
	ProjectList []string `gorm:"type:text[]"`

}

type Acaunt_User struct{
	gorm.Model
	GitHub string `gorm:"unique"`
	WakaTime string `gorm:"unique"`
	Permision string
	Premium bool
	Requests int 
}


var db *gorm.DB

func InitDB() {
//	dsn := "host=postgres user=postgres password=2242 dbname=mydb port=5432 sslmode=disable"
//	dsn := "postgres://localhost:2242@postgres:5432/mydb"
	dsn := os.Getenv("DATABASE_URL")
	var err error

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error DB Conect: %v", err)
	}

	if err := db.AutoMigrate(&UsersWaka{}); err != nil {
		log.Fatalf("Error Migration UsersWaka %v", err)
	}

	if err := db.AutoMigrate(&Acaunt_User{}); err != nil {
		log.Fatalf("Error Migration Acaunt_User %v", err)
	}

}

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



package database

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() {
	log.SetOutput(os.Stderr)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	//	dsn := "host=postgres user=postgres password=2242 dbname=mydb port=5432 sslmode=disable"
	//	dsn := "postgres://localhost:2242@postgres:5432/mydb"
	dsn := os.Getenv("DATABASE_URL")
	var err error

	// Открытие базы 
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Fatal Open DB %v", err)
	}
	//Миграция
	if err := db.AutoMigrate(&User{}, &Wall{}, &Comment{}); err != nil {
		log.Printf("Fatal Migration %v", err)
	}

	//CreateBase()

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



/*
func ID_Creator(id_comment int) (int, error) {
	//Выбираем в таблице User пользователя с именем user_name и возращаем его ID
	commentSearchId := Comment{}

	if err := db.Find(&commentSearchId, "Id_Comment = ?", id_comment).Error; err != nil {
		return 0, err
	}

	return commentSearchId.Id_Commentator, nil
}

*/




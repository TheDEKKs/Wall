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
	Name_User string `gorm:"unique;not null"`
	Id_Telegram int  `gorm:"unique;not null"`
	Id_Wall int `gorm:"unique"`
}

// Стурктура базы данных стены
type Wall struct {
	// Настройка стены
	Id_Wall int `gorm:"unique;primaryKey"`
	Id_Creator int  `gorm:"unique;not null"`

	// Настройки для коментариев на стене
	Mat bool 
	Anonymously bool 

	
}

// Структура базы данных коментариев
type Comment struct {
	// Данные 
	Id_Comment int `gorm:"unique;primaryKey"`
	Id_Wall int `gorm:"unique;not null"`
	Id_Commentator int `gorm:"unique; not null"`

	Text_Comment string `gorm:"size:128; not null"`
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

func CreateNewComment(ID_Creator, ID_Wall int, Comment string) bool {
	comments := Comment{Id_Wall: ID_Wall, Id_Commentator: ID_Creator, Text_Comment: Comment}
	if err := db.Create(&comments).Error; err != nil {
		fmt.Println("Error create comment - %v", err)
		return false
	}

	return true
}


//Создание стены
func CerateWall(ID_Creator int) int {
	//Создание строки с ID создателя которое
	wall := Wall{Id_Creator: ID_Creator}
	
	if err := db.Create(&wall).Error; err != nil {
		fmt.Println("Error Create Wall - %v", err)
		return 0
	}	
	//Получаем ID стены и возвращаем его
	return wall.Id_Wall
	/*var wall_inf Wall
	result := db.First(&wall_inf, "Id_Creator = ?", ID_Creator)	//Выбираем строку с ID создателя
	if result.Error != nil {
		fmt.Println("Error Search Wall - %v", result.Error)
	} else {
		return wall_inf.Id_Wall
	}
	
	return 0000 */
}

//Функция для добавления в таблицу нового пользователя
func AddUser(Telegram_ID int, Username string) {
	
	//Создание записи
	user := User{Name_User: Username, Id_Telegram: Telegram_ID, Id_Wall: 0000}
	if err := db.Create(&user).Error; err != nil {
		fmt.Println("Error add user %v", err)
	}
	/*
	var us User
	result := db.First(&us, "Id_Telegram = ?", Telegram_ID)
	if result.Error != nil {
		fmt.Println("error: %v", result.Error)
	} else {*/
		//Создание стены
		wall := CerateWall(user.Id_User)
		
		//Если успешно создалась стена 
		if wall != 0 {
			if err := db.Model(&User{}).Where("Id_Telegram = ?", Telegram_ID).Updates(User{Id_Wall: wall}).Error; err != nil {
				fmt.Println("Error Updates Id_Wall - %v", err)
			}
		} else {
			fmt.Println("Error Create wall, delet user...")
			db.Delete(&User{}, user.Id_User)
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



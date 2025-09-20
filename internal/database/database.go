package database

import (
	"log"
	"os"
	//"time"
	"fmt"
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

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panicln("Error DB Conect: %v", err)
	}


	if err := db.AutoMigrate(&User{}, &Wall{}, &Comment{}); err != nil {
		fmt.Println("Error Migration %v", err)
	}



	//CreateBase()

}

func CreateNewComment(ID_Creator, ID_Wall int, Comment_User string) bool {
	comments := Comment{Id_Wall: ID_Wall, Id_Commentator: ID_Creator, Text_Comment: Comment_User}
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
func AddUser(Telegram_ID int, Username, password_user string) string {
	
	//Создание записи
	user := User{Name_User: Username, Id_Telegram: Telegram_ID, Id_Wall: 0000, password: password_user}
	if err := db.Create(&user).Error; err != nil {
		fmt.Println("Error add user %v", err)
		var us User
		if err := db.First(&us, "Name_User = ?", Username); err != nil {
			return "Name User Not Unique"
		}
		if err := db.First(&us, "Id_Telegram = ?", Telegram_ID); err != nil {
			return "ID Telegram Not Unique"
		}
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
				return "Error Upadtes Id_Wall"
			}
		} else {
			fmt.Println("Error Create wall, delet user...")
			db.Delete(&User{}, user.Id_User)
			return "Error Create Wall"
		}
	
	return "Good"

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



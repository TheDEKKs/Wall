package database

import (
	"log"
	"os"

	//"time"
	"fmt"
	jsonstr "thedekk/webapp/internal/json"

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
		log.Printf("Fatal Open DB %v", err)
	}

	if err := db.AutoMigrate(&User{}, &Wall{}, &Comment{}); err != nil {
		log.Printf("Fatal Migration %v", err)
	}

	//CreateBase()

}

func CreateNewComment(ID_Creator, ID_Wall int, Comment_User string) (int, error) {
	comments := Comment{Id_Wall: ID_Wall, Id_Commentator: ID_Creator, Text_Comment: Comment_User}
	if err := db.Create(&comments).Error; err != nil {
		return 0, err
	}

	return comments.Id_Comment, nil
}

// Создание стены
func CerateWall(ID_Creator int) (int, error) {
	//Создание строки с ID создателя которое
	wall := Wall{Id_Creator: ID_Creator}

	if err := db.Create(&wall).Error; err != nil {
		return 0, err
	}
	//Получаем ID стены и возвращаем его
	return wall.Id_Wall, nil
	/*var wall_inf Wall
	result := db.First(&wall_inf, "Id_Creator = ?", ID_Creator)	//Выбираем строку с ID создателя
	if result.Error != nil {
		fmt.Println("Error Search Wall - %v", result.Error)
	} else {
		return wall_inf.Id_Wall
	}

	return 0000 */
}

// Функция для добавления в таблицу нового пользователя
func AddUser(Telegram_ID int, Username, password_user string) error {

	//Создание записи
	user := User{Name_User: Username, Id_Telegram: Telegram_ID, Id_Wall: 0000, password: password_user}
	if err := db.Create(&user).Error; err != nil {
		var us User
		if err := db.First(&us, "Name_User = ?", Username); err != nil {
			return err.Error
		}
		if err := db.First(&us, "Id_Telegram = ?", Telegram_ID); err != nil {
			return err.Error
		}
	}

	/*
		var us User
		result := db.First(&us, "Id_Telegram = ?", Telegram_ID)
		if result.Error != nil {
			fmt.Println("error: %v", result.Error)
		} else {*/
	//Создание стены
	wall, err := CerateWall(user.Id_User)

	//Если успешно создалась стена
	if err == nil {
		if err := db.Model(&User{}).Where("Id_Telegram = ?", Telegram_ID).Updates(User{Id_Wall: wall}).Error; err != nil {
			return err
		}
	} else {
		fmt.Println("Error Create wall, delet user...")
		db.Delete(&User{}, user.Id_User)
		return err
	}
		return nil

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

func SearchComment(id_wall_search string) ([]jsonstr.CommentRequest, error) {
	comment := []Comment{}

	if err := db.Find(&comment, "Id_Wall = ?", id_wall_search).Error; err != nil {
		return []jsonstr.CommentRequest{}, err
	}

	commentRequstUS := []jsonstr.CommentRequest{}

	for _, com := range comment {
		fmt.Println(com)
		fmt.Println(com.Text_Comment)
			commentRequstUS = append(commentRequstUS, jsonstr.CommentRequest{
			Id_Commentor: com.Id_Commentator,
			Comment:      com.Text_Comment,
		})

	}

	return commentRequstUS, nil

}

func ID_User(user_name string) (int, error) {
	userSearchId := User{}

	if err := db.Find(&userSearchId, "Name_User = ?", user_name).Error; err != nil {
		return 0, err
	}

	return userSearchId.Id_User, nil
}

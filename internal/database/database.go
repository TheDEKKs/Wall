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

func CreateNewComment(ID_Creator, ID_Wall int, Comment_User string) (int, error) {
	//Создаем структуру
	comments := Comment{Id_Wall: ID_Wall, Id_Commentator: ID_Creator, Text_Comment: Comment_User}
	//Создаем поле
	if err := db.Create(&comments).Error; err != nil {
		return 0, err
	}
	//Возращаем ID коментария 
	return comments.Id_Comment, nil
}

// Создание стены
func CerateWall(ID_Creator int) (int, error) {
	//Создание строки с ID создателя который создал стену
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
	//Создаем заппись
	if err := db.Create(&user).Error; err != nil {
		return err
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
		//В юзера добавляем ID стены
		if err := db.Model(&User{}).Where("Id_Telegram = ?", Telegram_ID).Updates(User{Id_Wall: wall}).Error; err != nil {
			return err
		}
	} else {
		//Если есть ошибка то удаляем юзера из базы
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
	//Выбиравем коментарии где Id_Wall == id_wall_search 
	if err := db.Find(&comment, "Id_Wall = ?", id_wall_search).Error; err != nil {
		return []jsonstr.CommentRequest{}, err
	}

	//Структура для возращение коментариев
	commentRequstUS := []jsonstr.CommentRequest{}

	//Перебираем коментарии и добавляем в comentRequstUS
	for _, com := range comment {
		fmt.Println(com)
		fmt.Println(com.Text_Comment)
			commentRequstUS = append(commentRequstUS, jsonstr.CommentRequest{
			Id_Commentor: com.Id_Commentator,
			Comment:      com.Text_Comment,
		})

	}
	//Возращаем все коментарии в JSON
	return commentRequstUS, nil

}

func ID_User(user_name string) (int, error) {
	//Выбираем в таблице User пользователя с именем user_name и возращаем его ID
	userSearchId := User{}

	if err := db.Find(&userSearchId, "Name_User = ?", user_name).Error; err != nil {
		return 0, err
	}

	return userSearchId.Id_User, nil
}

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

func UpdateComentDB(id_comment, id_user  int, new_comment string) error{
	fmt.Println("Start UPDATEDB")
	if err := db.Model(&Comment{}).Where("Id_Comment = ? AND Id_Commentator = ?", id_comment, id_user).Updates(Comment{Text_Comment: new_comment}).Error; err != nil {
		return err
	}
	fmt.Println("Stop UPDATEDB")

	return  nil

}
package database

import(
	jsonstr "thedekk/webapp/internal/json"

)

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
			commentRequstUS = append(commentRequstUS, jsonstr.CommentRequest{
			Id_Commentor: com.Id_Commentator,
			Comment:      com.Text_Comment,
		})

	}
	//Возращаем все коментарии в JSON
	return commentRequstUS, nil

}



func SearchAllComment(id int) ([]jsonstr.ReturnAllComment, error){
	commentSearch := []Comment{}
	//Поиск всех по ID косентариев
	if err := db.Find(&commentSearch, "Id_Commentator = ?", id).Error; err != nil {
		return []jsonstr.ReturnAllComment{}, err
	}

	commentAnswer := []jsonstr.ReturnAllComment{}

	//Перебор всех коментариев пользовтаелей
	for _, com := range commentSearch {
		commentAnswer = append(commentAnswer, jsonstr.ReturnAllComment{
			Id_Wall: com.Id_Wall,
			Id_Comment: com.Id_Comment,
			Text_Comment: com.Text_Comment,	
		})
	}

	return commentAnswer, nil

}

func UpdateComentDB(id_comment, id_user  int, new_comment string) error{
	if err := db.Model(&Comment{}).Where("Id_Comment = ? AND Id_Commentator = ?", id_comment, id_user).Updates(Comment{Text_Comment: new_comment}).Error; err != nil {
		return err
	}

	return  nil

}
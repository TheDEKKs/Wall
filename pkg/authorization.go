package pkg

import (
	database "thedekk/webapp/internal/database"
	jsonstr "thedekk/webapp/internal/json"
)

func NewCommentCreate(data jsonstr.NewCommentRequest) (int, error){
		
		//Получаем данные которые были в токене
		vallid, err := ValidateToken(data.Token)
		if err != nil {
			return 0, err
		}

		//Получаем ID пользователя из его имени 
		id_User, err := database.ID_User(vallid.Name) 
		if err != nil {
			return 0, err
		}

		//Создаем новый коментарий 
		id, err := database.CreateNewComment(id_User, data.ID_Wall, data.Text_coment)
		if err != nil{
			return 0, err
		}
		
		//Возращаем ID коментария
		return id, nil
}
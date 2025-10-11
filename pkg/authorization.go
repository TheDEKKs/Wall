package pkg

import (
	"fmt"
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


func UpdateComment(data jsonstr.EditComment) (bool, error) {
	dataToken, err := ValidateToken(data.Token)

	if err != nil {
		return false, err
	}


	id, err := database.ID_User(dataToken.Name)

	if err != nil {
		return false,err
	}

	if err := database.UpdateComentDB(data.Id_Comment, id, data.New_Comment); err != nil {
		return false, err
	}

	return true, nil
}


func ExaminationAfftion(token string, mat bool) error {
	data, err := ValidateToken(token)

	if err != nil {
		if err.Error() == "Token invalid" {
			return fmt.Errorf("Token invalid")
		}
		return err
		
	}

	id_User, err := database.ID_User(data.Name) 

	id_Wall, err := database.SearchWallUser(id_User)

	if err := database.UpdateSetingsWall(mat, id_Wall); err != nil {
		return err
	}



	return nil
}
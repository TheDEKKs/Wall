package repositories

import (
	database "thedekk/webapp/internal/database"
	jsonstr "thedekk/webapp/internal/json"
	pkg "thedekk/webapp/pkg"

)

func NewCommentCreate(data jsonstr.NewCommentRequest) (int, error){
		
		//Получаем данные которые были в токене
		vallid, err := pkg.ValidateToken(data.Token)
		if err != nil {
			return 0, err
		}


		//Создаем новый коментарий 
		id, err := database.CreateNewComment(vallid.Id, data.ID_Wall, data.Text_coment)
		if err != nil{
			return 0, err
		}
		
		//Возращаем ID коментария
		return id, nil
}


func UpdateComment(data jsonstr.EditComment) (bool, error) {
	//Снова проверяем на фалидность
	dataToken, err := pkg.ValidateToken(data.Token)

	if err != nil {
		return false, err
	}

	//Получаем ID 
	// ***Про это я как раз иговорил потом в данные токены нужно добавить ID
	if err != nil {
		return false,err
	}


	//Обновляем коментарий
	if err := database.UpdateComentDB(data.Id_Comment, dataToken.Id, data.New_Comment); err != nil {
		return false, err
	}

	return true, nil
}

//Функция которая обновляет натсройки стены
func ExaminationAfftion(token string, mat bool) error {
	//Снова проверяем на фалидность
	data, err := pkg.ValidateToken(token)

	if err != nil {
		return err
		
	}

	id_Wall, err := database.SearchWallUser(data.Id)

	if err := database.UpdateSetingsWall(mat, id_Wall); err != nil {
		return err
	}


	return nil
}
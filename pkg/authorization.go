package pkg

import (
	database "thedekk/webapp/internal/database"
	jsonstr "thedekk/webapp/internal/json"
)

func NewCommentCreate(data jsonstr.NewCommentRequest) (int, error){
		vallid, err := ValidateToken(data.Token)
		if err != nil {
			return 0, err
		}

		id_User, err := database.ID_User(vallid.Name) 
		if err != nil {
			return 0, err
		}
		id, err := database.CreateNewComment(id_User, data.ID_Wall, data.Text_coment)
		if err != nil{
			return 0, err
		}

		return id, nil
}
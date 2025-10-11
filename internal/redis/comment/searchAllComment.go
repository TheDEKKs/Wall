package comment

import (
	"encoding/json"
	jsonstr "thedekk/webapp/internal/json"
	r "thedekk/webapp/internal/redis"

)

//Функция для поиска по ID
func SerachAllCommentUser(id_filed string) ([]jsonstr.ReturnAllComment, error) {
	//Даные по ключу ID

	data, err := r.Rdb.HGetAll(r.Ctx, id_filed +":all").Result()

	if err != nil {
		return nil, err
	}

	//Записываем в нужный тип 
	var comment []jsonstr.ReturnAllComment
	for _, com := range data {
		var c []jsonstr.ReturnAllComment

		if err := json.Unmarshal([]byte(com), &c); err != nil {
			return nil, err
		}

		comment = append(comment, c...)
	}


	return comment, nil
}
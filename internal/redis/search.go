package redis

import (
	"encoding/json"
	jsonstr "thedekk/webapp/internal/json"
)

//Функция для поиска по ID
func Serach(id_filed int) ([]jsonstr.ReturnAllComment, error) {
	//Даные по ключу ID

	data, err := rdb.HGetAll(ctx, string(id_filed)).Result()

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
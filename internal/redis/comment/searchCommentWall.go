package comment


import (
	"encoding/json"
	jsonstr "thedekk/webapp/internal/json"
	r "thedekk/webapp/internal/redis"
)

func SearchAllCommentWall(id_wall string) ([]jsonstr.CommentRequest, error) {
	//Получаем результат поиска по ID стены
	data, err := r.Rdb.HGetAll(r.Ctx, id_wall).Result()
	if err != nil {
		return nil, err
	}
	
	//Перебираем и Записываем в нужный тип для дольнейшего удобства
	var comment []jsonstr.CommentRequest
	for _, com := range data {
		var c []jsonstr.CommentRequest

		if err := json.Unmarshal([]byte(com), &c); err != nil {
			return nil, err
		}

		comment = append(comment, c...)
	}


	return comment, nil
}
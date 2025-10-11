package comment


import (
	"encoding/json"
	jsonstr "thedekk/webapp/internal/json"
	r "thedekk/webapp/internal/redis"
)

func SearchAllCommentWall(id_wall string) ([]jsonstr.CommentRequest, error) {
	data, err := r.Rdb.HGetAll(r.Ctx, id_wall).Result()
	if err != nil {
		return nil, err
	}

	//Записываем в нужный тип 
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
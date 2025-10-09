package comment


import (
	"encoding/json"
	jsonstr "thedekk/webapp/internal/json"
	r "thedekk/webapp/internal/redis"

)


func NewRecordingWallComent(key string, data []jsonstr.CommentRequest) (error) {
	jsonData, err := json.Marshal(data)

	if err != nil {
		return err
	}

	if err := r.Rdb.HSet(r.Ctx, key, "wall_commnet",jsonData).Err(); err != nil {
		return err
	}

	return nil
}
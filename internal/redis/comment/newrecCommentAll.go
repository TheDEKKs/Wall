package comment


import (
	"encoding/json"
	jsonstr "thedekk/webapp/internal/json"
	r "thedekk/webapp/internal/redis"

)


func NewRecordingAllComent(key string, data []jsonstr.ReturnAllComment) (error) {
	jsonData, err := json.Marshal(data)

	if err != nil {
		return err
	}

	if err := r.Rdb.HSet(r.Ctx, key + ":all", "all_commnet",jsonData).Err(); err != nil {
		return err
	}

	return nil
}
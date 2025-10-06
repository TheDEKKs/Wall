package redis


import (
	"encoding/json"
	jsonstr "thedekk/webapp/internal/json"
)


func NewRecording(key string, data []jsonstr.ReturnAllComment) (error) {
	jsonData, err := json.Marshal(data)

	if err != nil {
		return err
	}

	if err := rdb.HSet(ctx, key, "all_commnet",jsonData).Err(); err != nil {
		return err
	}

	return nil
}
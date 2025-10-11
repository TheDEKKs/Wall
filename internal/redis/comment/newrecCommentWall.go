package comment

import (
	"encoding/json"
	jsonstr "thedekk/webapp/internal/json"
	r "thedekk/webapp/internal/redis"
	loger "thedekk/webapp/pkg/loger"

)


func NewRecordingWallComent(key string, data []jsonstr.CommentRequest) (error) {
	jsonData, err := json.Marshal(data)
	if len(jsonData) < 0 {
		loger.Zap.Info("Data < 0")
	}
	if err != nil {
		return err
	}

	if err := r.Rdb.HSet(r.Ctx, key, "wall_comment", jsonData).Err(); err != nil {
		return err
	}

	return nil
}
package router 

import (
	"github.com/gin-gonic/gin"
	 wall "thedekk/webapp/internal/router/wall"
)

func InitRouter(rout_gin *gin.Engine) {
	wall_grup := rout_gin.Group("/wall")
	{
		wall.EditComment(wall_grup)
		wall.NewComment(wall_grup)
		wall.WallViewer(wall_grup)
		wall.EditWall(wall_grup)
	}
	Login(rout_gin)
	EditComment(rout_gin)

}

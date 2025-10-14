package router 

import (
	"github.com/gin-gonic/gin"
	wall "thedekk/webapp/internal/router/wall"
	login "thedekk/webapp/internal/router/login"
	comment "thedekk/webapp/internal/router/comment"

)

func InitRouter(rout_gin *gin.Engine) {
	//Тут я надеюсь ненадо ничего обьяснять 
	wall_grup := rout_gin.Group("/wall")
	{
		wall.EditComment(wall_grup)
		wall.NewComment(wall_grup)
		wall.WallViewer(wall_grup)
		wall.EditWall(wall_grup)
	}


	login.Login(rout_gin)


	comment_grup := rout_gin.Group("/comment")
	{
		comment.EditComment(comment_grup)
	}

}

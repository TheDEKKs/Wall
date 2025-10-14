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
		wall.WallViewer(wall_grup)
		wall.EditWall(wall_grup)
	}

	//Группа для создания акаунта или входа в него
	login_grup := rout_gin.Group("/au")
	{
		login.Registration(login_grup)
		login.LoginAcaunt(login_grup)
	}
	

	//Група для работы с коментариями
	comment_grup := rout_gin.Group("/comment")
	{
		comment.SearchUserComment(comment_grup)
		comment.NewComment(comment_grup)
		comment.EditComment(comment_grup)
	}

}

package wall

import (
	jsonstr "thedekk/webapp/internal/json"
	pkg "thedekk/webapp/pkg"

	"github.com/gin-gonic/gin"
)

func NewComment(r *gin.RouterGroup) {
	r.POST("/newcomment", createNewComment)
}

func createNewComment(c *gin.Context) {
	//Получаем JSON данные для создание комента
	jsonComment := jsonstr.NewCommentRequest{}
	if err := c.ShouldBindJSON(&jsonComment); err != nil {
		c.JSON(500, gin.H{"Error": err.Error()})
		return
	}

	token, err := c.Cookie("TOKEN_JWT")
	if err != nil {
		c.JSON(500, gin.H{"Error": err.Error})
	}

	jsonComment.Token = token
	//Создание коментариия от аунтафицированого пользователя
	id, err := pkg.NewCommentCreate(jsonComment)
	if err != nil {
		c.JSON(500, gin.H{"Error": err})
		return
	}

	//Возращаем ID коментария если все прошло хорошо
	c.JSON(200, gin.H{"comment_id": id})

}


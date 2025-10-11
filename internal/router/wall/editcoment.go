package wall

import (
		"github.com/gin-gonic/gin"
		jsonstr "thedekk/webapp/internal/json"
		pkg "thedekk/webapp/pkg"

)

func EditComment(r *gin.RouterGroup) {
	r.PUT("/editcomment", newDataComment)
}

func newDataComment(c *gin.Context) {
	//Получаем JSON 
	UpdateComment := jsonstr.EditComment{}

	if err := c.ShouldBindJSON(&UpdateComment); err != nil {
		c.JSON(500, gin.H{"Error JSON": err})
		return
	}


	//И про токен тоже не забываем
	Tokens, err := c.Cookie("TOKEN_JWT")

	if err != nil { 
		c.JSON(500, gin.H{"Error COOKIE": err})
		return

	}


	//Тут опять добавляем 
	UpdateComment.Token = Tokens


	//И обновляем
	good, err := pkg.UpdateComment(UpdateComment)

	if err != nil {
		c.JSON(500, gin.H{"Error Update": err})
		return
	}

	c.JSON(200, gin.H{"status": "Good create", "return": good})

}
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
	UpdateComment := jsonstr.EditComment{}

	if err := c.ShouldBindJSON(&UpdateComment); err != nil {
		c.JSON(500, gin.H{"Error JSON": err})
		return
	}

	Tokens, err := c.Cookie("TOKEN_JWT")

	if err != nil { 
		c.JSON(500, gin.H{"Error COOKIE": err})
		return

	}

	UpdateComment.Token = Tokens

	good, err := pkg.UpdateComment(UpdateComment)

	if err != nil {
		c.JSON(500, gin.H{"Error Update": err})
		return
	}

	c.JSON(200, gin.H{"status": "Good create", "return": good})

}
package wall

import (
		"github.com/gin-gonic/gin"
		jsonstr "thedekk/webapp/internal/json"

)

func EditComment(r *gin.RouterGroup) {
	r.PUT("/editcomment", newDataComment)
}

func newDataComment(c *gin.Context) {
	UpdateComment := jsonstr.EditComment{}

	if err := c.ShouldBindJSON(&UpdateComment); err != nil {
		c.JSON(500, gin.H{"Error": err})
	}

}
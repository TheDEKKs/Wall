package wall


import (
	"github.com/gin-gonic/gin"
	jsonstr "thedekk/webapp/internal/json"

)

func NewComment(r *gin.RouterGroup) {
	r.POST("/newcomment", createNewComment)
}

func createNewComment(c *gin.Context) {
	jsonComment := jsonstr.NewCommentRequest{}
	if err := c.ShouldBindJSON(&jsonComment); err != nil {
		c.JSON(500, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(200, jsonComment)

}

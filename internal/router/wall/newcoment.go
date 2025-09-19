package wall


import (
	"github.com/gin-gonic/gin"
)

func NewComment(r *gin.RouterGroup) {
	r.POST("/newcomment", createNewComment)
}

func createNewComment(c *gin.Context) {}

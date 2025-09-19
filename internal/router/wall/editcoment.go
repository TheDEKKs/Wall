package wall

import (
		"github.com/gin-gonic/gin"
)

func EditComment(r *gin.RouterGroup) {
	r.PUT("/editcomment", newDataComment)
}

func newDataComment(c *gin.Context) {}
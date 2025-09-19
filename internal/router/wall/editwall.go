package wall

import (
		"github.com/gin-gonic/gin"
)

func EditWall(r *gin.RouterGroup) {
	r.PUT("/editwall", newDataWall)
}

func newDataWall(c *gin.Context) {}
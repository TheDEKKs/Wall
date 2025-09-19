package wall

import (
	"github.com/gin-gonic/gin"
)

func WallViewer(r *gin.RouterGroup) {
	r.GET("/:id", dataWall)
}

func dataWall(c *gin.Context) {}

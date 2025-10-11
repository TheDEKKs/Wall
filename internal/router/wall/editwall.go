package wall

import (
	"strconv"
	pkg "thedekk/webapp/pkg"

	"github.com/gin-gonic/gin"
)

func EditWall(r *gin.RouterGroup) {
	r.PUT("/editwall", newDataWall)
}

func newDataWall(c *gin.Context) {
	mat := c.Query("mat")
	id := c.Query("Wall")

	token, err := c.Cookie("TOKEN_JWT")
	if err != nil {
		c.JSON(500, gin.H{"Error": err.Error})
	}
	
	id_wall, _ := strconv.Atoi(id)

	switch mat {
	case "true":
		if err := pkg.ExaminationAfftion(token, id_wall, true); err != nil {
			c.JSON(500, gin.H{"Error": err.Error})
			return
		}
	case "false":
		if err := pkg.ExaminationAfftion(token, id_wall, true); err != nil {
			c.JSON(500, gin.H{"Error": err.Error})

			return
		}
	default:
		c.JSON(500, gin.H{"Error": err.Error})
		return
	}

	c.JSON(200, gin.H{"answer": "Good update"})

	
}
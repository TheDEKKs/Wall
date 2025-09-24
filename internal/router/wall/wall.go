package wall

import (
	"fmt"
	pkg "thedekk/webapp/pkg"

	"github.com/gin-gonic/gin"
)

func WallViewer(r *gin.RouterGroup) {
	r.GET("/:id", dataWall)
}

func dataWall(c *gin.Context) {
	 cookiesJS, err := c.Cookie("TOKEN_JWT")
	 if err != nil {
		c.String(500, "Error %s", err)
		return
	}

	user, pass, err := pkg.ValidateToken(cookiesJS)
	fmt.Println(cookiesJS)
	if err != nil {
		c.String(500, err.Error())
	}
	c.JSON(200, gin.H{"user":user, "pass":pass,})
	
} 

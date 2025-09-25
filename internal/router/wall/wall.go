package wall

import (
	//"fmt"
	//pkg "thedekk/webapp/pkg"
	//"fmt"
	database "thedekk/webapp/internal/database"

	"github.com/gin-gonic/gin"
)

func WallViewer(r *gin.RouterGroup) {
	r.GET("/:id", dataWall)
}

func dataWall(c *gin.Context) {
	json_comment, err := database.SearchComment(c.Param("id"))

	if err != nil {
		c.JSON(500, gin.H{"status": "Error Search Comment", "Error": err.Error()})
		return
	}

	c.JSON(200, json_comment)
	 /*
	cookiesJS, err := c.Cookie("TOKEN_JWT")
	 if err != nil {
		c.String(500, "Error %s %s", err, cookiesJS)
		return
	}
	
	
	user, err := pkg.ValidateToken(cookiesJS)
	fmt.Println(cookiesJS)
	if err != nil {
		c.String(500, err.Error())
	}
	c.JSON(200, gin.H{"user":user.Name, "pass":user.Password,})
	*/
} 

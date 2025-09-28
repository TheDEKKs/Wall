package router

import (
	"strconv"
	database "thedekk/webapp/internal/database"

	"github.com/gin-gonic/gin"
)


func EditComment(r *gin.Engine) {
	r.GET("/searchallcomment", SearchAllComment)

}

func SearchAllComment(c *gin.Context){
	//Если есть такие поля записываем
	id := c.Query("id")
	idInt, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(500, gin.H{"error id in int": err})
		return
	}
	answer, err := database.SearchAllComment(idInt)

	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}

	c.JSON(200, answer)
}
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
	hach := c.Query("hach")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(500, gin.H{"error id in int": err})
		return
	}

	hachs, err := strconv.Atoi(hach)

	answer, err := database.SearchAllComment(idInt, hachs)

	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}

	c.JSON(200, answer)
}
package comment

import (
	"strconv"
	database "thedekk/webapp/internal/database"

	"github.com/gin-gonic/gin"
)


func EditComment(r *gin.RouterGroup) {
	r.GET("/searchallcomment", SearchAllComment)

}

func SearchAllComment(c *gin.Context){
	//Если есть такие поля записываем
	id := c.Query("id")
	//Если хеш не 1 то мы ищем в хеше
	//Еслти 1 то пропускаем хеш
	hach := c.Query("hach")

	idInt, _ := strconv.Atoi(id)

	hachs, _ := strconv.Atoi(hach)


	//Вызываем функцию поиска всех моентариев
	answer, err := database.SearchAllComment(idInt, hachs)

	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}

	//Возращаем коментарии
	c.JSON(200, answer)
}
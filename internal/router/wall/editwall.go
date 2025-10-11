package wall

import (
	"net/http"
	pkg "thedekk/webapp/pkg"

	"github.com/gin-gonic/gin"
)

func EditWall(r *gin.RouterGroup) {
	r.PUT("/editwall", newDataWall)
}

func newDataWall(c *gin.Context) {
	//Получаем нужные нам данные из запроса
	mat := c.Query("mat")
	token, err := c.Cookie("TOKEN_JWT")

	if err != nil {
		if err == http.ErrNoCookie{
			c.JSON(500, gin.H{"Error": "Not Cookie, Not Login"})
			return
		}
		c.JSON(500, gin.H{"Error": err.Error})
		return 
	}
	
	switch mat {
	//Тут мы отправляем запрос на обновление стены нашего пользователя
	//Да я сюда впихнул switch конструкция ничего лучше я не придумал
	case "true":
		if err := pkg.ExaminationAfftion(token, true); err != nil {
			c.JSON(500, gin.H{"Error": err.Error})
			return
		}
	case "false":
		if err := pkg.ExaminationAfftion(token, false); err != nil {
			c.JSON(500, gin.H{"Error": err.Error})

			return
		}
	default:
		c.JSON(500, gin.H{"Error": err.Error})
		return
	}

	c.JSON(200, gin.H{"answer": "Good update"})

	
}
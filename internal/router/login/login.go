package login

import (
	"github.com/gin-gonic/gin"
)


func LoginAcaunt(c *gin.RouterGroup) {
	c.POST("/login", login)
}

func login(c *gin.Context) {

}
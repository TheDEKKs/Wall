package router

import (
	"net/http"
	database "thedekk/webapp/internal/database"
	jsonstr "thedekk/webapp/internal/json"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gin-gonic/gin"
	"os"
)

func Login(r *gin.Engine) {
	r.POST("/login", login_post)
}

func login_post (c *gin.Context){
	//database.AddUser(1, "Tets", "pass")
	
	var json_login jsonstr.POST_Login
	if err := c.ShouldBindJSON(&json_login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.String(500, "Error JSON Request")
		return
	}

	if err := database.AddUser(json_login.ID_Telegram, json_login.Name, json_login.Password); err == "Good" {
		c.String(200, "Good create user")
		return
	} else {
		c.String(500, "Error Add User %s", err)
		return
	}




}


/*func jwtCreate() {
	os.Getenv("secretKey")
}*/
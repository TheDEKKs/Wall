package router

import (
	//"fmt"

	"fmt"
	"net/http"
	database "thedekk/webapp/internal/database"
	jsonstr "thedekk/webapp/internal/json"
	pkg "thedekk/webapp/pkg"

	//"time"

	"github.com/gin-gonic/gin"
)


func Login(r *gin.Engine) {
	r.POST("/login", login_post)
}

func login_post (c *gin.Context){
	//database.AddUser(1, "Tets", "pass")
	//SetCookie("TOKEN_JWT", "test", 1000, "/", "", false, false)


	var json_login jsonstr.POST_Login

	if err := c.ShouldBindJSON(&json_login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Error JSON request", "error": err.Error()})
		return
	}



	if err := database.AddUser(json_login.ID_Telegram, json_login.Name, json_login.Password); err == nil {

		token, err := pkg.JwtCreate(json_login.Name, json_login.Password)

		if err != nil {
			c.JSON(500, gin.H{"status": "Error create token", "Error": err,})
			return
		}

		maxAge := 14 * 86400
		
		fmt.Println(token)
		c.SetCookie("TOKEN_JWT", string(token), maxAge, "/", "", false, false)
		c.JSON(http.StatusOK, gin.H{"status": "Create User and Cookie"})
		
	
	} else {
		c.JSON(500, gin.H{"status": "Error create user", "Error": err,})
		return
	}

/*
	cookie := &http.Cookie{
		Name: "JWT_TOKEN",
		Value: token,
		Path: "/",
		Domain: "localhost",
		MaxAge: maxAge,
		HttpOnly: false,
		Secure: false,
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(c.Writer, cookie)

	*/

//	c.String(200, "%s, %s", cookies, err.Error())

	return
}


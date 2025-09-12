package main

import (
	database "thedekk/webapp/bootstrap"
//	"net/http"
	"github.com/gin-gonic/gin"
)
func main() {
        r := gin.Default()
	
	database.InitDB()
	database.Add()
	r.GET("/", func(c *gin.Context){
		c.String(200, "Тут будет главная траница")
	})
	

	// Стена
	wall := r.Group("/wall") 
	{

		wall.GET("/:username", func(c *gin.Context) {
			c.String(200, "Тут будет стена человека")
		})

		wall.POST("/newcomment", func(c *gin.Context) {
			c.String(200, "Запрос на создание комента")
		})

		wall.PUT("/edit", func(c *gin.Context) {
			c.String(200, "Запрос на изменение коментария") 
		})

		wall.POST("/createwall", func(c *gin.Context) {
			c.String(200, "Запрос на создание стены")
		})

	}


	// Панель управления 
	adminpage := r.Group("adminpage")
	{
		adminpage.GET("/", func(c *gin.Context) {
			c.String(200, "Страница для входа")
		})

		adminpage.POST("/login", func(c *gin.Context) {
			c.String(200, "Логинимся в панеле")
		})

		adminpage.GET("/controlpage", func(c *gin.Context) {
			c.String(200, "Страница для управлением сайта")
		})

	}


	// Настройка акаунта
	user_wall_panel := r.Group("acaunt")
	{
		user_wall_panel.POST("mywall/settings", func(c *gin.Context) {
			c.String(200, "Настройка твоей стены")
		})

		user_wall_panel.POST("mywall/comment/delete", func(c *gin.Context) {
			c.String(200, "Запрос на удаление определеного комента на стене")
		})	
	}


        r.Run()

    }


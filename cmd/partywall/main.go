package main

import (
	//"net/http"
	database "thedekk/webapp/internal/database"
	router "thedekk/webapp/internal/router"

	//	"net/http"
	"github.com/gin-gonic/gin"
	//"golang.org/x/text/message"
)

type comment_user_post struct {
	text_comment string `json:"name"`
}

func main() {
	r := gin.Default()

	router.InitRouter(r)

	database.InitDB()

	/*	database.AddUser(123456789, "EgorTest")
		r.GET("/", func(c *gin.Context){
			c.String(200, "Тут будет главная траница")
		})


		// Стена
		wall := r.Group("/wall")
		{


		  	//GET ЗАПРОСЫ
			wall.GET("/:username", func(c *gin.Context) {
				c.String(200, "Тут будет стена человека")
			})

			wall.GET("win/newcomment", func(c *gin.Context) {
				c.String(200, "Окно с созданием комента")
			})


			//POST, PUT ЗАПРОСЫ ДЛЯ РАБОТЫ СО СТЕНОЙ И КОМЕНТАРИЯМИ
			wall.POST("/newcomment", func(c *gin.Context) {
				var comment comment_user_post
				c.String(200, "Запрос на создание комента")

				if err := c.ShouldBindJSON(&comment); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
				}

				c.JSON(http.StatusOK, gin.H{
					"message": "Good",
					"comment": comment.text_comment,
				})
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

	*/


	r.Run()

}

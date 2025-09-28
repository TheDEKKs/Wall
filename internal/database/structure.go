package database

import ("time")

//Структура базы данных юзера
type User struct {
	// Данные акаунта 
	Id_User int `gorm:"unique;primaryKey"`
	Name_User string `gorm:"unique;not null"`
	Id_Telegram int  `gorm:"unique;not null"`
	Id_Wall int `gorm:"unique"`
	password string 
}

// Стурктура базы данных стены
type Wall struct {
	// Настройка стены
	Id_Wall int `gorm:"unique;primaryKey"`
	Id_Creator int  `gorm:"unique;not null"`

	// Настройки для коментариев на стене
	Mat bool 
	Anonymously bool 

	
}

// Структура базы данных коментариев
type Comment struct {
	// Данные 
	Id_Comment int `gorm:"unique;primaryKey" json"id"`
	Id_Wall int `gorm:"not null" json:"id_wall"`
	Id_Commentator int `gorm:"not null" json"id_commentor"`

	Text_Comment string `gorm:"size:128; not null" json"text"`
	CreatedAt time.Time
	UpdatedAt time.Time
}





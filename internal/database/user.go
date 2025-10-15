package database

import (
	loger "thedekk/webapp/pkg/loger"

)

//Ищем ID пользователя по его нику
func ID_User(user_name string) (int, error) {
	//Выбираем в таблице User пользователя с именем user_name и возращаем его ID
	userSearchId := User{}

	if err := db.Find(&userSearchId, "Name_User = ?", user_name).Error; err != nil {
		return 0, err
	}

	return userSearchId.Id_User, nil
}

// Функция для добавления в таблицу нового пользователя
func AddUser(Telegram_ID int, Username, password_user string) error {

	//Создание записи
	user := User{Name_User: Username, Id_Telegram: Telegram_ID, Id_Wall: 0000, Password: password_user}
	//Создаем заппись
	if err := db.Create(&user).Error; err != nil {
		loger.Zap.Warn(err.Error())
		return err
	}

	/*
		var us User
		result := db.First(&us, "Id_Telegram = ?", Telegram_ID)
		if result.Error != nil {
			fmt.Println("error: %v", result.Error)
		} else {*/
	//Создание стены
	wall, err := CerateWall(user.Id_User)

	//Если успешно создалась стена
	if err == nil {
		//В юзера добавляем ID стены
		if err := db.Model(&User{}).Where("Id_Telegram = ?", Telegram_ID).Updates(User{Id_Wall: wall}).Error; err != nil {
			loger.Zap.Warn(err.Error())
			return err
		}
	} else {
		//Если есть ошибка то удаляем юзера из базы
		db.Delete(&User{}, user.Id_User)
		loger.Zap.Warn(err.Error())
		return err
	}
		return nil

}
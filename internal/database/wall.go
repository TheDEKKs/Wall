package database

// Создание стены
func CerateWall(ID_Creator int) (int, error) {
	//Создание строки с ID создателя который создал стену
	wall := Wall{Id_Creator: ID_Creator}

	if err := db.Create(&wall).Error; err != nil {
		return 0, err
	}
	//Получаем ID стены и возвращаем его
	return wall.Id_Wall, nil
	/*var wall_inf Wall
	result := db.First(&wall_inf, "Id_Creator = ?", ID_Creator)	//Выбираем строку с ID создателя
	if result.Error != nil {
		fmt.Println("Error Search Wall - %v", result.Error)
	} else {
		return wall_inf.Id_Wall
	}

	return 0000 */
}
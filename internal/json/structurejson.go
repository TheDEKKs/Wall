package json

//Структура для анутфикации
type POST_Login struct{
	Password string `json:"password"`	//Пароль нужен
	Name string `json:"User"`	//Имя пользователя
	ID_Telegram int `json:"ID_Telegram"`	//ID телеграма
}

//Стуруктура для нового коментария
type NewCommentRequest struct{
	Token string `json:"token"`		//Токен создателя
	Text_coment string `json:"comment"` 	//Текст кментария
	ID_Wall int `json:"id_wall"`	//ID стены на котором записан
}

//Стурктура для возразение коментариев 
type CommentRequest struct {
	Id_Commentor int `json:"id"`	//ID создателя
	Comment string `json:"comment`	//Текст коментария
}

type EditComment struct {
	Token string `json:"token"`		//Токен
	Id_Comment int `json:"id_comment"`	//ID коментария который будет изменен
	New_Comment string `json:"new_comment"`		//Новый коментарий
}


type ReturnAllComment struct {
	Id_Wall int `json:"Id_Wall"`	//Стена коментария
	Id_Comment int `json:"Id_Comment"`		//ID коментария
	Text_Comment string `json:"text"`	//Текст коментария
}
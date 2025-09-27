package json

//Структура для анутфикации
type POST_Login struct{
	Password string `json:"password"`
	Name string `json:"User"`
	ID_Telegram int `json:"ID_Telegram"`
}

//Стуруктура для нового коментария
type NewCommentRequest struct{
	Token string `json:"token"`
	Text_coment string `json:"comment"` 
	ID_Wall int `json:"id_wall"`
}

//Стурктура для возразение коментариев 
type CommentRequest struct {
	Id_Commentor int `json:"id"`
	Comment string `json:"comment`
}
package json

type POST_Login struct{
	Password string `json:"password"`
	Name string `json:"User"`
	ID_Telegram int `json:"ID_Telegram"`
}
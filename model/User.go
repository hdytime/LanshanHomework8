package model

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Like struct {
	ArticleId int `json:"article_id"`
	UserId    int `json:"user_id"`
}

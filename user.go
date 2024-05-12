package todo

type User struct {
	Id       int    `json:"-"`
	Name     string `json:"name" binding:"required"`
	Username string `json:"Username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

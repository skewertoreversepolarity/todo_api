package todo

type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

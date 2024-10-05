package models

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
	AddTime  int64  `json:"add_time"`
}

func (u User) TableName() string {
	return "user"
}

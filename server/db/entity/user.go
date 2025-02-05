package entity

type User struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	CreatedAt int64  `json:"created_at"`
}

package entity

type User struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	CreatedAt  int64  `json:"created_at"`
	UUID       string `gorm:"unique" json:"uuid"`
	JoinedRoom string `gorm:"-" json:"joined_room"`
}

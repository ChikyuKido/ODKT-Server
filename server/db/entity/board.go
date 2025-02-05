package entity

type Board struct {
	ID        uint        `gorm:"primaryKey" json:"id"`
	Name      string      `json:"name"`
	Payload   []byte      `json:"payload"`
	TypeClass interface{} `gorm:"-" json:"type_class"`
}

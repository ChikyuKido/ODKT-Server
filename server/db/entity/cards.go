package entity

type CardStreet struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	City       string `json:"city"`
	Name       string `gorm:"unique" json:"name"`
	Rent       string `json:"rent"`
	Price      int32  `json:"price"`
	HousePrice int32  `json:"house_price"`
	HotelPrice int32  `json:"hotel_price"`
}

type CardChance struct {
	ID        uint        `gorm:"primaryKey" json:"id"`
	Type      string      `json:"type"`
	Text      string      `json:"text"`
	Payload   []byte      `json:"payload"`
	TypeClass interface{} `gorm:"-" json:"type_class"`
}

type CardBank struct {
	ID        uint        `gorm:"primaryKey" json:"id"`
	Type      string      `json:"type"`
	Text      string      `json:"text"`
	Payload   []byte      `json:"payload"`
	TypeClass interface{} `gorm:"-" json:"type_class"`
}

type CardRailroad struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Price int32  `json:"price"`
	Name  string `gorm:"unique" json:"name"`
}

type CardSpecial struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Name      string `gorm:"unique" json:"name"`
	Price     int32  `json:"price"`
	PriceName string `json:"price_name"`
}

type CardOther struct {
	ID        uint        `gorm:"primaryKey" json:"id"`
	Name      string      `gorm:"unique" json:"name"`
	Payload   []byte      `json:"payload"`
	TypeClass interface{} `gorm:"-" json:"type_class"`
}

type CardSettings struct {
	ID      uint   `gorm:"primaryKey" json:"id"`
	Name    string `gorm:"unique" json:"name"`
	Payload []byte `json:"payload"`
}

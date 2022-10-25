package users

type User struct {
	Id      string  `gorm:"primaryKey; not null" json:"id"`
	Balance float64 `json:"balance"`
}

type Users struct {
	Users []User `json:"users"`
}

package users

type User struct {
	Id      string `gorm:"primaryKey; not null" json:"id"`
	Balance int    `json:"balance"`
	Version int    `json:"version"`
}

type Users struct {
	Users []User `json:"users"`
}

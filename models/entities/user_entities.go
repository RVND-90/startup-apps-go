package entities

type User struct {
	Id int64 `gorm:"id,primaryKey"`
	Username string `json:"username" gorm:"username"`
	Password string `json:"password" gorm:"password"`
	Salt string `gorm:"salt"`
}
package entity

type User struct {
	Base
	Username  string
	FirstName string
	LastName  string
	Email     string `gorm:"unique"`
	Password  string
	Active    bool
}

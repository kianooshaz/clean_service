package entity

type User struct {
	Base
	Username  string
	FirstName string
	LastName  string
	Email     string
	Password  string
	Active    bool
}

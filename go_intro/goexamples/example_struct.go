package goexamples

type User struct {
	id    string
	email string
	Name  string
	Phone string
}

//NewUser creates new USer
func NewUser() User {
	return User{}
}

func (u *User) GetName() string {
	return u.GetName()
}

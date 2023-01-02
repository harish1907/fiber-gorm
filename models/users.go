package models

import "golang.org/x/crypto/bcrypt"

type MyUser struct {
	Id        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" gorm:"unique"`
	Password  []byte `json:"-"`
}


func (user *MyUser) SetPassword(password string){
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	user.Password = hash
}

func (user *MyUser) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword(user.Password, []byte(password))
}
package internal

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"unicode"
)

type UserService struct {
	repo *Repository
}

func NewUserService(repository *Repository) *UserService{
	return &UserService{repo: repository}
}

func (us *UserService)SignUp(user User) (err error){
	if !isNotEmpty(user){
		return fmt.Errorf("enter all required data")
	}
	if user.Psw!=user.ConfPsw{
		return fmt.Errorf("passwords does not match")
	}
	if err=verifyPassword(user.Psw);err!=nil{
		return err
	}
	if _,err=us.repo.Get(user.Name);err==nil{
		return fmt.Errorf("user with this name is already registered")
	}
	if user.PswHash,err=bcrypt.GenerateFromPassword([]byte(user.Psw),bcrypt.DefaultCost);err!=nil{
		return fmt.Errorf("cannot hash password")
	}
	if err := us.repo.Create(user);err!=nil{
		return err
	}
	return nil
}

func (us *UserService)LogIn(user User) (u *User,err error){
	storedUser,err:=us.repo.Get(user.Name)
	if err!=nil{
		return nil,fmt.Errorf("there is no user with this name")
	}
	if err=bcrypt.CompareHashAndPassword(storedUser.PswHash,[]byte(user.Psw));err!=nil{
		return nil,fmt.Errorf("wrong password")
	}
	return storedUser,nil
}

func verifyPassword(s string) error {
	letters := 0
	number,upper,special,sevenOrMore:=false,false,false,false
	for _, c := range s {
		switch {
		case unicode.IsNumber(c):
			number = true
		case unicode.IsUpper(c):
			upper = true
			letters++
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			special = true
		case unicode.IsLetter(c):
			letters++
		case c==' ':
			return fmt.Errorf("password cannot contain spaces")
		}
	}
	sevenOrMore = letters >= 7
	if number==false{
		return fmt.Errorf("password must contain at least 1 number")
	}
	if upper==false{
		return fmt.Errorf("password must contain at least 1 uppercase letter")
	}
	if special==false{
		return fmt.Errorf("password must contain at least special character")
	}
	if sevenOrMore==false{
		return fmt.Errorf("password must contain 7 or more symbols")
	}
	return nil
}

func isNotEmpty(u User)bool{
	return u.Name!=""||u.Psw!=""||u.ConfPsw!=""||u.Number!=""||u.Color!=""
}

package handlers

import (
	"L567/internal"
	"html/template"
	"log"
	"net/http"
)

var Errors = make(map[string]string)

func MakeLoginHandler(service *internal.UserService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl, _ := template.ParseFiles("./static/index.html")
		tpl.Execute(w, Errors)
		Errors["Login"] = ""
	}
}

func MakeCheckSignUpHandler(service *internal.UserService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := service.SignUp(internal.User{
			Name:    r.FormValue("Name"),
			Psw:     r.FormValue("Password"),
			ConfPsw: r.FormValue("ConfirmPassword"),
			Number:  r.FormValue("PhoneNumber"),
			Color:   r.FormValue("Color"),
		})
		if err != nil {
			Errors["SignUp"] = err.Error()
			log.Print(Errors)
			http.Redirect(w, r, "/signup", http.StatusSeeOther)
		} else {
			Errors["SignUp"] = ""
		}
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}

func MakeSignUpHandler(service *internal.UserService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl, _ := template.ParseFiles("./static/signup.html")
		tpl.Execute(w, Errors)
		Errors["SignUp"] = ""
	}
}

func MakePersonalCabinetHandler(service *internal.UserService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := service.LogIn(internal.User{
			Name: r.FormValue("Name"),
			Psw:  r.FormValue("Password"),
		})
		if err != nil {
			Errors["Login"] = err.Error()
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		} else {
			tpl, _ := template.ParseFiles("./static/personal_cabinet.html")
			tpl.Execute(w, user)
		}
	}
}
